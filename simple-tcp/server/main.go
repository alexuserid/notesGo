package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	// открываем порт для коннекта
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message Received:", string(message))
		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage + "\n"))
	}
}
