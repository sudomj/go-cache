package main

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/x1bdev/go-cache/pkg/server"
)

func main() {

	port := 5000
	server := server.NewServer(port)
	go server.Listen()

	for i := range 10 {

		slog.Info(fmt.Sprintf("client %d", i))
		client := NewClient(port)

		go func() {
			time.Sleep(time.Second)

			err := client.WriteToServer([]byte("set key value"))

			if err != nil {
				panic(err)
			}
		}()
	}

	select {}
}

type Client struct {
	port int
}

func NewClient(port int) *Client {
	{
		return &Client{
			port: port,
		}
	}
}

func (c *Client) WriteToServer(message []byte) error {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", c.port))

	if err != nil {
		slog.Error("could not connect to the server", "error", err.Error())
		return err
	}

	n, err := conn.Write(message)

	if err != nil {
		slog.Error("could not write to the server", "error", err.Error())
		return err
	}

	slog.Info(fmt.Sprintf("written %d bytes into the server", n))
	return nil
}
