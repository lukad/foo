package mail

import (
	"bytes"
	"io/ioutil"
	"net/mail"

	"github.com/lukad/helix/smtpd"
)

type Mail struct {
	Sender     string      `json:"sender"`
	Recipients []string    `json:"recipients"`
	Header     mail.Header `json:"header"`
	Body       []byte      `json:"body"`
}

func ParseEnvelope(envelope smtpd.Envelope) (Mail, error) {
	var msg *mail.Message
	var err error

	r := bytes.NewReader(envelope.Data)

	if msg, err = mail.ReadMessage(r); err != nil {
		return Mail{}, err
	}

	body, _ := ioutil.ReadAll(msg.Body)

	return Mail{
		Sender:     envelope.Sender,
		Recipients: envelope.Recipients,
		Header:     msg.Header,
		Body:       body,
	}, nil
}
