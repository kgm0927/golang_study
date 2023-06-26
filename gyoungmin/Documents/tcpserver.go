package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	dstream, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()

	for {
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
	conn.Close()
}
