package ws

import (
	"github.com/sirupsen/logrus"
	"github/trad3r/go_temp.git/internal/models"
	"os"
)

const (
	OutputTypeNewUser        = "newUser"
	OutputTypeMessage        = "newMessage"
	OutputTypeUnregisterUser = "unregisterUser"
)

type Hub struct {
	Logger     *logrus.Logger
	Clients    map[*Client]bool
	Broadcast  chan InputMessage
	Register   chan *Client
	Unregister chan *Client
}

type InputMessage struct {
	Client *Client
	Text   []byte
}

type OutputMessage struct {
	Type        string
	currentUser *models.User
	Template    string
	Text        []byte
}

func newOutputMessage(typeMessage string, currentUser *models.User, currentClient bool, text []byte) OutputMessage {
	tmpl := myMessage
	if !currentClient {
		tmpl = otherMessage
	}

	return OutputMessage{
		Type:        typeMessage,
		currentUser: currentUser,
		Template:    tmpl,
		Text:        text,
	}
}

func NewHub(f *os.File) *Hub {
	return &Hub{
		Logger:     newLogger(f),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan InputMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func newLogger(f *os.File) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(f)

	return logger
}

func (h Hub) Run() {
	h.Logger.Info("Start hub")
	for {
		select {
		case c := <-h.Register:
			h.Clients[c] = true
			for client := range h.Clients {
				client.Send <- newOutputMessage(OutputTypeNewUser, c.User, false, nil)
				if c.User != client.User {
					c.Send <- newOutputMessage(OutputTypeNewUser, client.User, false, nil)
				}
			}
			h.Logger.Info("Register new client:", c.User)
		case client := <-h.Unregister:
			h.unregisterClient(client)
			for client := range h.Clients {
				client.Send <- newOutputMessage(OutputTypeUnregisterUser, client.User, false, nil)
			}
		case inputMessage := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- newOutputMessage(OutputTypeMessage, inputMessage.Client.User, inputMessage.Client == client, inputMessage.Text):
				default:
					h.unregisterClient(client)
				}
			}
		}
	}
}

func (h Hub) unregisterClient(client *Client) {
	if _, ok := h.Clients[client]; ok {
		delete(h.Clients, client)
		close(client.Send)
		h.Logger.Info("Unregister client:", client.User)
	}
}
