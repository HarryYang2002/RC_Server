package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	u := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("cannot upgrade: %v\n", err)
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("cannot close connect: %v\n", err)
		}
	}(conn)

	done := make(chan struct{})
	go func() {
		for {
			m := make(map[string]interface{})
			err := conn.ReadJSON(&m)
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseNoStatusReceived, websocket.CloseNormalClosure) {
					fmt.Printf("unexpect read error； %v\n", err)
				}
				done <- struct{}{}
				break
			}
			fmt.Printf("message received: %v\n", m)
		}
	}()

	i := 0
	for {
		select {
		case <-time.After(time.Millisecond * 200):
		case <-done:
			return
		}

		i++
		err := conn.WriteJSON(map[string]interface{}{
			"Hello": "web socket",
			"msg":   strconv.Itoa(i),
		})
		if err != nil {
			fmt.Printf("cannot write JSON: %v\n", err)
		}
	}
}

//func handleWebSocket(w http.ResponseWriter, r *http.Request) {
//	_, err := fmt.Fprintf(w, "Hello Web Socket")
//	if err != nil {
//		return
//	}
//}
