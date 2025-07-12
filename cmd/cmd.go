package cmd

import (
	"flag"
	"fmt"
	"net"

	client "github.com/AnhBigBrother/tcp-chat/internal"
)

var (
	HOST string
	PORT int
)

func init() {
	flag.StringVar(&HOST, "host", "localhost", "Specify server host")
	flag.IntVar(&PORT, "port", 8080, "Specify port number")
	flag.Parse()
}

func Execute() {
	conn, err := net.Dial("tcp", net.JoinHostPort(HOST, fmt.Sprintf("%d", PORT)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := client.NewClient(conn)
	fmt.Printf("Connected to %s, port %d.\n", HOST, PORT)

	go c.HandleServerResponse()
	c.HandleInput()
}
