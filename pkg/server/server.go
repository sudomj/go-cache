package server

import (
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {

	return &Server{
		Port: port,
	}
}

func (s *Server) Listen() error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))

	if err != nil {
		slog.Info("could not start listen server", "error", err.Error())
		return err
	}

	s.accept(listener)
	return nil
}

func (s *Server) accept(listener net.Listener) {

	for {

		conn, err := listener.Accept()

		if err != nil {
			slog.Error("could not accept connection", "error", err.Error())
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {

	slog.Info(fmt.Sprintf("new incomming connection from %s", conn.RemoteAddr().String()))

}
