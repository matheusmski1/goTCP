package main

import (
	"goTcp/server"
)

func main() {
	server.RunTcpServer()
	go server.RunTcpClientConction()
}
