package messagingdriver

import (
	"context"
	"log"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gorilla/mux"
	gorillaWS "github.com/gorilla/websocket"

	"github.com/prasetyowira/message/internal/app/chat/messaging/messagingdriver/websocket"
)

//nolint
var (
	upgrader = gorillaWS.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func RegisterWebSocketHandlers(router *mux.Router, subscriber message.Subscriber) {
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWS(w, r, subscriber)
	})
}

func ServeWS(w http.ResponseWriter, r *http.Request, subscriber message.Subscriber) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(gorillaWS.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	messages, err := subscriber.Subscribe(context.Background(), websocket.MessagingTopic)
	if err != nil {
		panic(err)
	}
	go websocket.Writer(ws, messages)
	go websocket.WriterPing(ws)
	websocket.Reader(ws)
}
