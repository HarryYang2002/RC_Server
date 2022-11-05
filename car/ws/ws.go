package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"server/car/mq"
)

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
					if !websocket.IsCloseError(err, websocket.CloseNoStatusReceived, websocket.CloseNormalClosure) {
						logger.Warn("unexpect read error", zap.Error(err))
					}
					done <- struct{}{}
					break
				}
			}
		}()

		for {
			select {
			case msg := <-msgs:
				//停顿三秒，方便观察，todo：fix
				//time.Sleep(3 * time.Second)
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
