package websocket

import (
	"fmt"
	"time"

	watermillMessage "github.com/ThreeDotsLabs/watermill/message"
	gorillaWS "github.com/gorilla/websocket"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 5 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	messagePeriod = 3 * time.Second

	MessagingTopic = "messaging"
)

var (
	Upgrader  = gorillaWS.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)


func Reader(ws *gorillaWS.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	_ = ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func WriterPing(ws *gorillaWS.Conn) {
	pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		pingTicker.Stop()
		_ = ws.Close()

	}()

	for {
		select {
		case <-pingTicker.C:
			_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(gorillaWS.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func Writer(ws *gorillaWS.Conn, messages <-chan *watermillMessage.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		// message := string(msg.Payload)

		_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(gorillaWS.TextMessage, msg.Payload); err != nil {
			return
		}

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

