package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) HandleServerMessages() {
	reader := bufio.NewReader(c.conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("The server is down.")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Error when reading message: ", err.Error())
			os.Exit(1)
		}

		fmt.Printf("\033[2K\r")

		msg := message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println("Cannot unmarshal payload: ", err.Error())
			return
		}

		event, payload := msg.Event, msg.Payload

		switch event {
		case ERROR:
			fmt.Println("Error:", payload)
		case MESSAGE:
			fmt.Println(payload)
		}

		fmt.Print("> ")
	}
}

func (c *Client) HandleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")

		if len(args) == 0 ||
			(len(args) == 1 && EVENT(args[0]) != GET_ROOMS) {
			fmt.Printf("\r")
			fmt.Println("Error: Wrong format.")
			continue
		}

		c.conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
		payload, _ := json.Marshal(
			message{
				Event:   EVENT(args[0]),
				Payload: strings.Join(args[1:], " "),
			},
		)
		payload = append(payload, '\n')
		_, err := c.conn.Write(payload)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
