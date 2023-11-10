package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	//listen on a port
	ln, err := net.Listen("tcp", "127.0.0.1:9009")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Listerning... ")

	for {
		// accept connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received ", msg)
	}
	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9009")
    if err != nil {
		fmt.Println(err)
	}
	// send message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}