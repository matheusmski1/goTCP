package process

import (
	"a3-prototipo/model"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"sync"
)

var connection net.Conn

func ClientTCP(wg *sync.WaitGroup, port string, student model.Student) {
	//defer wg.Done()
	connection, _ = net.Dial("tcp", ":"+port)
	for {
		Send(student)
		recv()
		return
	}

}

//func ClientTCP(wg *sync.WaitGroup, port string) {
//	defer wg.Done()
//
//	c, err = net.Dial("tcp", ":"+port)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for {
//		reader := bufio.NewReader(os.Stdin)
//		fmt.Print(">> ")
//		text, _ := reader.ReadString('\n')
//		fmt.Fprintf(c, text+"\n")
//
//		message, _ := bufio.NewReader(c).ReadString('\n')
//		fmt.Print("->: " + message)
//		if strings.TrimSpace(string(text)) == "STOP" {
//			fmt.Println("TCP client exiting...")
//			return
//		}
//	}
//}
//
//var encoder *gob.Encoder
//
//func SendStudentTCP(student model.Student) {
//	fmt.Println(student)
//	if encoder == nil {
//		encoder = gob.NewEncoder(c)
//	}
//	err := encoder.Encode(student)
//	if err != nil {
//		fmt.Println(err)
//		print("erro no encode")
//	}
//
//}

func Send(student model.Student) {
	buffer := new(bytes.Buffer)

	encoder := gob.NewEncoder(buffer)
	encoder.Encode(student)

	connection.Write(buffer.Bytes())
}

func recv() {
	tmp := make([]byte, 500)
	connection.Read(tmp)

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	tmpbuff := bytes.NewBuffer(tmp)
	tmpstruct := new(model.Student)

	// creates a decoder object
	gobobjdec := gob.NewDecoder(tmpbuff)
	// decodes buffer and unmarshals it into a Message struct
	gobobjdec.Decode(tmpstruct)

	fmt.Println(tmpstruct)
}
