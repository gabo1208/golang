package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	who := conn.RemoteAddr().String()
	fmt.Printf("You are %s.\n", who)
	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatal(err)
		}

		exitClient(done)
	}()
	inputProcessing(conn, done)
	conn.Close()
	<-done
}

func inputProcessing(dst io.Writer, done chan struct{}) {
	input := bufio.NewReader(os.Stdin)
Loop:
	for {
		text, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(text)
			break
		}

		switch text {
		case ":q\n", ":Q\n":
			go exitClient(done)
			break Loop
		default:
			if _, err := io.Copy(dst, strings.NewReader(text)); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func exitClient(done chan struct{}) {
	log.Println("done")
	done <- struct{}{}
}
