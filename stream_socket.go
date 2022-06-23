package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	fmt.Println("Connection Accepted")
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(conn)

	for scanner.Scan() {
		n, err := writer.WriteString(scanner.Text() + "\n")

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
