package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	//  коннектимся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")

	for {
		// os.Stdin создает и открывает новый файл, указывающий на
		// стандартный вывод
		// os.Stdin == os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")

		// создается новый гошный буферизированный ридер для os.Stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		// чтение из stdin
		text, _ := reader.ReadString('\n')
		// запись во врайтер conn, таким образом байты передаются в сокет
		fmt.Fprint(conn, text + "\n")

		// читаем то, что пришло в качестве ответа на сокет
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}
