package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWritter(conn, ch)
	who := conn.RemoteAddr().String()
	messages <- who + " has arrived"
	entering <- enter{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Err() != nil {
			log.Fatal(input.Err().Error())
			break
		}
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWritter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
