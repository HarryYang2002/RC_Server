package ws

import (
	"context"
	"net/http"
	"server/car/mq"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Handler creates a websocket http handler.
func Handler(u *websocket.Upgrader, sub mq.Subscriber, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			logger.Warn("cannot upgrade", zap.Error(err))
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				logger.Warn("cannot close connect", zap.Error(err))
				return
			}
		}(conn)

		msgs, cleanUP, err := sub.Subscribe(context.Background())
		defer cleanUP()
		if err != nil {
			logger.Error("cannot subscribe", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		done := make(chan struct{})
		go func() {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err,
						websocket.CloseGoingAway,
						websocket.CloseNormalClosure) {
						logger.Warn("unexpected read error", zap.Error(err))
					}
					done <- struct{}{}
					break
				}
			}
		}()

		for {
			select {
			case msg := <-msgs:
				err := conn.WriteJSON(msg)
				if err != nil {
					logger.Warn("cannot write JSON", zap.Error(err))
				}
			case <-done:
				return
			}
		}
	}
}
