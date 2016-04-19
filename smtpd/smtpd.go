package smtpd

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	s.handleClient(newClient(conn))
}

func (s *Server) handleClient(c *client) {
	l, _ := c.readLine()
	fmt.Println(l)
	c.conn.Close()
}

func (s *Server) Listen(address string) (err error) {
	s.listener, err = net.Listen("tcp", address)
	if err != nil {
		return err
	}

	go s.listen()
	return nil
}
