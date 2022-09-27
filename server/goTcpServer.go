package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func RunTcpClientConction() {
	c, err := net.Dial("tcp", "1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> Digite algo")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client desligando...")
			return
		}
	}
}

func RunTcpServer() {
	PORT := ":1234"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
