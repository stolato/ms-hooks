package main

import (
	"ms-hooks/internal/handlers"
	"ms-hooks/pkg/socket"
)

func main() {
	io := socket.InitSocket()
	go socket.SocketI(io)
	handlers.IniHandler(io)
}
