package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	timerFile  = "./logs/timer"
	timerSleep = time.Second * 15
)

type Server struct {
	httpServer http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	go timer()

	return s.httpServer.ListenAndServe()
}

// timer - запись текущего времени каждые 15 секунд, чтобы крон перезапускал сервер
func timer() {
	timerfile, err := os.OpenFile(timerFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		logrus.Panic(err)
	}
	defer timerfile.Close()

	for {
		timerfile.Truncate(0)
		timerfile.Seek(0, 0)
		timerfile.WriteString(time.Now().Format("2006-01-02T15:04:05+07:00"))
		time.Sleep(timerSleep)
	}
}
