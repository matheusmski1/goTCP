package process

import (
	"a3-prototipo/model"
	"bytes"
	"encoding/gob"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func ServerTCP(wg *sync.WaitGroup, port string, parent fyne.Window) {
	defer wg.Done()

	server, _ := net.Listen("tcp", ":"+port)
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}
		go handle(conn, parent)
	}
}

func logerr(err error) bool {
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout:", err)
		} else if err == io.EOF {
		} else {
			log.Println("read error:", err)
		}
		return true
	}
	return false
}

func read(conn net.Conn, parent fyne.Window) {
	// create a temp buffer
	tmp := make([]byte, 500)

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {
		_, err := conn.Read(tmp)
		if logerr(err) {
			break
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(tmp)
		student := new(model.Student)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)
		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(student)

		// lets print out!
		fmt.Println(student)

		var total float64
		total = 0
		for i := range student.Notas {
			total += student.Notas[i]
		}
		media := total / float64(len(student.Notas))

		text := fmt.Sprintf("Aluno: %s", student.Nome)
		mediaStr := fmt.Sprintf("Média: %.2f", media)

		dialog.ShowCustom("Novo resultado", "visto", container.NewVBox(&canvas.Text{Text: text}, &canvas.Text{Text: mediaStr}), parent)

		return
	}
}

func resp(conn net.Conn) {
	//msg := Message{ID: "Yo", Data: "Hello back"}
	msg := "exito na ação"
	bin_buf := new(bytes.Buffer)

	// create a encoder object
	gobobje := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobje.Encode(msg)

	conn.Write(bin_buf.Bytes())
	conn.Close()
}

func handle(conn net.Conn, parent fyne.Window) {
	timeoutDuration := 2 * time.Second
	fmt.Println("Launching server...")
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	read(conn, parent)
	resp(conn)
}
