package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

const qName = "chat-messages"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createRabbitMQConn() (*amqp.Connection, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "localhost"
	}

	port := os.Getenv("RABBITMQ_PORT")
	if port == "" {
		port = "5672"
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s", url, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn, err
}

func createRabbitMQConsumer(ch *amqp.Channel, who string) {
	q, err := ch.QueueDeclare(
		fmt.Sprintf("%s-%s", qName, who), // name
		false,                            // durable
		false,                            // delete when unused
		false,                            // exclusive
		false,                            // no-wait
		nil,                              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func(who string) {
		for d := range msgs {
			log.Printf("%s Received a message: %s", who, d.Body)
		}
	}(who)

	log.Printf(" %s [*] Waiting for messages. To exit press CTRL+C", who)
	<-forever
}

func createRabbitMQChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch, err
}

func sendMessage(ch *amqp.Channel, to, msg string) {
	err := ch.Publish(
		"",                              // exchange
		fmt.Sprintf("%s-%s", qName, to), // routing key
		false,                           // mandatory
		false,                           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")
}

func handleConn(conn net.Conn, rabbitConn *amqp.Connection) {
	who := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	log.Print("ChatService: starting rabbitMQChannel per Thread...")
	ch, err := createRabbitMQChannel(rabbitConn)
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	log.Print("ChatService: Creating a message queue per Client...")
	go createRabbitMQConsumer(ch, who)

	for input.Scan() {
		if input.Err() != nil {
			log.Fatal(input.Err().Error())
			break
		}

		sendMessage(ch, who, input.Text())
	}

	log.Print("ChatService: Clossing rabbitMQ connection and channel...")
	conn.Close()
}

func chatListen(rabbitConn *amqp.Connection) {
	port := os.Getenv("CHAT_TCP_PORT")
	if port == "" {
		port = "32080"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, rabbitConn)
	}
}

func ChatService(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("ChatService: starting rabbitMQ connection...")
	rabbitConn, err := createRabbitMQConn()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConn.Close()

	go chatListen(rabbitConn)

	log.Print("ChatService: listening for incoming connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		var bArr []byte
		_, err = conn.Read(bArr)
		if err != nil {
			log.Print(err)
			continue
		}

		probe, err := isProbe(conn)
		if probe {
			if err != nil {
				log.Print(err)
			}
			continue
		}

		//go handleConn(conn, rabbitConn)
	}
}

func isProbe(conn net.Conn) (bool, error) {
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			return true, err
		}

		if strings.Contains(string(tmp[:n]), "R-Type: Probe") {
			conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 3\r\n\r\nfoo"))
			conn.Close()
			return true, nil
		}

		return false, nil
	}
}

func main() {
	log.Print("ChatService: starting server...")
	port := os.Getenv("CHAT_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("ChatService: listening on port %s", port)
	ChatService(port)
}
