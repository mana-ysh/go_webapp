package main

import (
	"github.com/gorilla/websocket"
)

// チャットしている一人のユーザ
type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

// クライアントがwebsocketからReadMessageでデータ読み込み
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// sendチャネルからメッセージを受け取って、WriteMessageで書き出し
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
