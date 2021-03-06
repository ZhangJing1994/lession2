package main

import (
	"bytes"
	"io"
	config "lession2/week9/situtaion/stickyPackage"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handler(conn)
	}
}

func handler(c net.Conn) {
	defer c.Close()
	buf := make([]byte, config.BufferSize)
	result := bytes.NewBuffer(nil)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("ending: " + err.Error())
				return
			} else {
				log.Println("Read error: " + err.Error())
				break
			}
		}
		result.Write(buf[0:n])
		log.Printf("recevie size[%d]: %v", n, result.String())
		result.Reset()
	}
}
