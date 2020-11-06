package handler

import (
	"github.com/gin-gonic/gin"
	"github/trad3r/go_temp.git/internal/ws"
)

type Hadler struct {
	Hub *ws.Hub
}

func NewHandler(hub *ws.Hub) *Hadler {
	return &Hadler{
		Hub: hub,
	}
}

func (h *Hadler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./public/assets")
	r.LoadHTMLFiles("public/index.html")

	r.GET("/", h.Main)
	r.GET("/ws", h.WsHandle)

	return r
}
