package smtp

import (
	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-smtp"
)

type Mail struct {
	sender    string
	receivers []string
	bodyType  smtp.BodyType
	message   parsemail.Email
}

func NewMail() *Mail {
	return &Mail{}
}

func (m Mail) GetSender() string {
	return m.sender
}
func (m Mail) SetSender(sender string) {
	m.sender = sender
}

func (m *Mail) AppendReceiver(recv string) {
	m.receivers = append(m.receivers, recv)
}

func (m *Mail) GetReceivers() []string {
	return m.receivers
}

func (m *Mail) SetReceivers(receivers []string) {
	m.receivers = receivers
}

func (m *Mail) GetBodyType() smtp.BodyType {
	return m.bodyType
}

func (m Mail) SetBodyType(bodyType smtp.BodyType) {
	m.bodyType = bodyType
}

func (m *Mail) GetMessage() parsemail.Email {
	return m.message
}

func (m *Mail) SetMessage(message parsemail.Email) {
	m.message = message
}
