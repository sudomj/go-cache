package main

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/x1bdev/go-cache/pkg/peer"
)

type Server struct {
	Port  int
	Peers chan *peer.Peer
}

func NewServer(port int) *Server {

	return &Server{
		Port:  port,
		Peers: make(chan *peer.Peer),
	}
}

func (s *Server) Listen() error {

	addr := fmt.Sprintf(":%d", s.Port)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		slog.Error("could not listen to tcp connections", "error", err.Error())
		return err
	}

	go s.readConnection()

	return s.acceptConnections(listener)
}

func (s *Server) acceptConnections(listener net.Listener) error {

	for {

		conn, err := listener.Accept()

		if err != nil {
			slog.Error("could not accept tcp connection", "error", err.Error())
			return err
		}

		s.Peers <- peer.New(conn)
	}
}

func (s *Server) readConnection() {

	for peer := range s.Peers {

		go peer.Read()
	}
}

type Client struct {
	conn net.Conn
}

func NewClient() (*Client, error) {

	conn, err := net.Dial("tcp", ":5000")

	if err != nil {
		slog.Error("could not establish connection", "error", err.Error())
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Write(data []byte) error {

	n, err := c.conn.Write(data)

	if err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("written %d bytes into the server", n))
	return nil
}

func main() {

	server := NewServer(5000)
	go server.Listen()

	go func() {

		time.Sleep(time.Second)

		client, err := NewClient()

		if err != nil {
			panic(err)
		}

		client.Write([]byte("My message"))
	}()

	select {}
}
