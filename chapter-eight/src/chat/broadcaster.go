package chat

import (
	"strings"
)

type client chan<- string
type enter struct {
	client
	ip string
}

var (
	entering = make(chan enter)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]string)

	for {
		select {
		case msg := <-messages:
			for client := range clients {
				if !strings.HasPrefix(msg, clients[client]) {
					client <- msg
				}
			}
		case cli := <-entering:
			clients[cli.client] = cli.ip
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
