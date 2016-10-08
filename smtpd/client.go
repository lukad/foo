package smtpd

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	in   *bufio.Reader
	out  *bufio.Writer
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
		in:   bufio.NewReader(conn),
		out:  bufio.NewWriter(conn),
	}
}

func (c *client) readLine() (line string, err error) {
	line, err = c.in.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(line, "\r\n"), nil
}
