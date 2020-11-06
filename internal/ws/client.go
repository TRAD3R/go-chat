package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github/trad3r/go_temp.git/internal/models"
	"html/template"
	"log"
	"time"
)

const (
	maxMessageSize = 512
	pongTime       = time.Second * 60
	pingTime       = time.Second * 50
	writeWait      = time.Second * 10

	myMessage    = "myMessage.html"
	otherMessage = "otherMessage.html"
)

type Client struct {
	User  *models.User
	Image string
	Hub   *Hub
	Conn  *websocket.Conn
	Send  chan OutputMessage
}

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		User: models.NewUser(),
		Hub:  hub,
		Conn: conn,
		Send: make(chan OutputMessage),
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongTime))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongTime)); return nil })

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message := bytes.TrimSpace(bytes.Replace(msg, newLine, space, -1))

		c.Hub.Broadcast <- InputMessage{Client: c, Text: message}
	}
}

func (c Client) WritePump() {
	ticker := time.NewTicker(pingTime)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case outputMessage, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			message, err := prepareMessage(outputMessage, c)
			if err != nil {
				logrus.Info(fmt.Sprintf("client.go 95: %v", err.Error()))
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			//n := len(c.Send)
			//for i := 0; i < n; i++ {
			//	w.Write(newLine)
			//	w.Write(<-c.Send)
			//}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

type MessageResponse struct {
	Text  string
	Date  string
	Image string
}

type NewUser struct {
	Id    string
	Name  string
	Image string
}

type Response struct {
	Type   string      `json:"type"`
	Text   string      `json:"text"`
	Append interface{} `json:"append"`
}

func prepareMessage(outputMessage OutputMessage, client Client) ([]byte, error) {
	var msg string
	var err error
	switch outputMessage.Type {
	case OutputTypeMessage:
		msg, err = newMessage(outputMessage)
	case OutputTypeNewUser:
		msg, err = newUser(client, outputMessage.currentUser)
	case OutputTypeUnregisterUser:
		msg, err = unregisterUser(outputMessage.currentUser)
	}

	if err != nil {
		return nil, err
	}

	response := Response{
		Type: outputMessage.Type,
		Text: msg,
	}

	response.Append = client.User

	return json.Marshal(&response)
}

func unregisterUser(user *models.User) (string, error) {
	return string(user.Id), nil
}

func newUser(client Client, user *models.User) (string, error) {
	tmpl, err := template.ParseFiles("./internal/templates/userOnline.html")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, NewUser{
		Id:    string(user.Id),
		Name:  user.Name,
		Image: user.Image,
	}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func newMessage(outputMessage OutputMessage) (string, error) {
	tmpl, err := template.ParseFiles("./internal/templates/" + outputMessage.Template)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, MessageResponse{
		Text:  string(outputMessage.Text),
		Date:  time.Now().Format("15:04 02.01.2006"),
		Image: outputMessage.currentUser.Image,
	}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
