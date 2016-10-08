package smtpd

import (
	"bitbucket.org/chrj/smtpd"
)

type Envelope smtpd.Envelope

type Server struct {
	*smtpd.Server
	envelopeChannel chan Envelope
}

func NewServer(envelopeChannel chan Envelope) *Server {
	server := &Server{
		Server:          &smtpd.Server{},
		envelopeChannel: envelopeChannel,
	}
	server.Handler = server.handler
	return server
}

func (s *Server) handler(peer smtpd.Peer, envelope smtpd.Envelope) error {
	s.envelopeChannel <- Envelope(envelope)
	return nil
}
