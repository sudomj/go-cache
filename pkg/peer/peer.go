package peer

import (
	"fmt"
	"log/slog"
	"net"
)

type Peer struct {
	conn net.Conn
}

func New(conn net.Conn) *Peer {

	return &Peer{conn: conn}
}

func (p *Peer) Read() {

	slog.Info(fmt.Sprintf("new incomming connection from %s", p.conn.RemoteAddr().String()))
}
