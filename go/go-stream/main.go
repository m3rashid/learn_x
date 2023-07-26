package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(connection)
	}
}

func (fs *FileServer) readLoop(connection net.Conn) {
	buffer := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(connection, binary.BigEndian, &size)

		n, err := io.CopyN(buffer, connection, size)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(buffer.Bytes())
		fmt.Printf("Received %d bytes \n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.BigEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader((file)), int64(size))
	if err != nil {
		return err
	}

	fmt.Printf("Written %d bytes over the network \n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(4000000)
	}()
	server := &FileServer{}
	server.start()
}
