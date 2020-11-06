package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github/trad3r/go_temp.git/internal/ws"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1 << 20,
	WriteBufferSize: 1 << 20,
}

func (h *Hadler) Main(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Hadler) WsHandle(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	client := ws.NewClient(h.Hub, conn)
	client.Hub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}
