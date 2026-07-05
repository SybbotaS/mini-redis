package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to Mini Redis")
	fmt.Println("Type 'exit' to quit")

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		command, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)

		if strings.EqualFold(command, "exit") {
			fmt.Println("Bye!")
			return
		}

		_, err = fmt.Fprintln(conn, command)
		if err != nil {
			log.Fatal(err)
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strings.TrimSpace(response))
	}
}
