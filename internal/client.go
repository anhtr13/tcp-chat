package client

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) HandleServerResponse() {
	for {
		dec := gob.NewDecoder(c.conn)
		msg := message{}
		err := dec.Decode(&msg)

		fmt.Printf("\033[2K\r")
		if err == io.EOF {
			fmt.Println("The server is down.")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		event, data := msg.Event, msg.Data
		switch event {
		case ERROR:
			fmt.Println("Error:", data)
		case MESSAGE:
			fmt.Println(data)
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

		args := strings.Split(input, " ")

		if len(args) == 0 ||
			(len(args) == 1 && EVENT(args[0]) != GET_ROOMS) {
			fmt.Printf("\r")
			fmt.Println("Error: Wrong format.")
			continue
		}

		msg := message{
			Event: EVENT(args[0]),
			Data:  strings.Join(args[1:], " "),
		}

		enc := gob.NewEncoder(c.conn)
		err := enc.Encode(msg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
