package socket

import (
	"fmt"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"ms-hooks/models"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var cli *socket.Client

func InitSocket() *socket.Server {
	io := socket.NewServer(nil, nil)
	http.Handle("/socket.io/", io.ServeHandler(nil))
	go func() {
		err := http.ListenAndServe(":8002", nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return io
}

func Emit(io *socket.Server, notification models.Notification, room socket.Room) {
	io.To(room).Emit("new-hook", notification)
}

func SocketI(io *socket.Server) {

	io.Of("/", nil).On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		client.On("disconnect", func(data ...any) {
			log.Println("closed", data)
		})
		client.On("join", func(data ...any) {
			channel := socket.Room(fmt.Sprintf("%v", data[0]))
			client.Join(channel)
		})
	})

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	io.Close(nil)
	os.Exit(0)
}
