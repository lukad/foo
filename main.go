package main

import (
	"os"

	"github.com/lukad/helix/mail"
	"github.com/lukad/helix/smtpd"
	"github.com/lukad/helix/store"
	"github.com/lukad/helix/web"

	log "github.com/Sirupsen/logrus"
)

var (
	host string = ""
	port string = "8080"
)

func init() {
	if envAddress := os.Getenv("HOST"); envAddress != "" {
		host = envAddress
	}
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
}

func main() {
	envelopeChannel := make(chan smtpd.Envelope)

	go func() {
		smtpServer := smtpd.NewServer(envelopeChannel)
		if err := smtpServer.ListenAndServe(":1025"); err != nil {
			log.Fatal(err)
		}
	}()

	s := store.New()

	go func() {
		for {
			envelope := <-envelopeChannel
			m, _ := mail.ParseEnvelope(envelope)
			s.Insert(m)
		}
	}()

	server := web.NewServer(s)
	if err := server.ListenAndServe(host + ":" + port); err != nil {
		log.Fatal(err)
	}
}
