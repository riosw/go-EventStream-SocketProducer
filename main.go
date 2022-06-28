package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	messageproducer "github.com/riosw/go-EventStream-SocketProducer/messageProducer"
)

func main() {
	var PORT string = ":9990"
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server Socket is Listening on Port %s \n", PORT)

	conn, err := listener.Accept()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connection Accepted")

	generator := messageproducer.SimpleStringMessageGenerator{
		StringMessage: "car",
	}

	writer := bufio.NewWriter(conn)

	for {
		n, err := writer.WriteString(generator.EmitMessage().Content + "\n")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d bytes written to connection\n", n)

		err = writer.Flush()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
