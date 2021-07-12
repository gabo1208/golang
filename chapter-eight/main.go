package main

import (
	"flag"
	"fmt"

	chat "example.com/chapter-eight/src/chat"
)

func main() {
	t := "Server"
	client := flag.Bool("Client", false, "Set if the program should be run as server or client.")
	flag.Parse()

	if *client {
		t = "Client"
		defer chat.Client()
	} else {
		defer chat.Chat()
	}

	fmt.Printf("Started chat as " + t + "\n")
}
