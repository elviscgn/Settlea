package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func (c *Client) handleNewMessage(m string, hub *Hub) {
	var messageData Message
	err := json.Unmarshal([]byte(m), &messageData)
	if err != nil {
		log.Println("Error unmarshaling message:", err)
		return
	}

	msg := &Message{
		Action:   messageData.Action,
		Content:  messageData.Content,
		RoomID:   c.RoomID,
		Username: c.Username,
	}

	switch messageData.Action {
	case SendMessageAction:
		hub.Broadcast <- msg

	case PingAction:
		msg.Action = "pong"
		hub.Broadcast <- msg
	default:
		log.Println("unimplemented action")

	}

}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// m = bytes.TrimSpace(bytes.Replace(m, newline, space, -1))

		c.handleNewMessage(string(m), hub)

	}
}
