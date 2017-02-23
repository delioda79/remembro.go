package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	protocol := Protocol{Commands: []string{}, length: 0, lengthString: ""}
	// Make a buffer to hold incoming data.
	for {
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		text := strings.TrimSpace(string(buf[:reqLen]))
		err = protocol.check(text)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(protocol.Commands)
		// Send a response back to person contacting us.
		conn.Write([]byte("Message received."))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
