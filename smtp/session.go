package smtp

import (
	"TMail/domain"
	"errors"
	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-smtp"
	"io"
	"log"
)

type Session struct {
	current *domain.Mail
	mails   []*domain.Mail
}

func (s *Session) AuthPlain(username, password string) error {
	return nil
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail!", from)
	s.current = domain.NewMail()
	s.current.SetSender(from)
	s.current.SetBodyType(opts.Body)
	s.mails = append(s.mails, s.current)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt!")
	if s.current == nil {
		return errors.New("sender is not defined")
	}
	s.current.AppendReceiver(to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	log.Println("Data!")
	if s.current == nil {
		return errors.New("sender is not defined")
	}
	if message, err := parsemail.Parse(r); err != nil {
		return err
	} else {
		s.current.SetMessage(message)
	}
	return nil
}

func (s *Session) Reset() {
	log.Println("Reset!")
}

func (s *Session) Logout() error {
	log.Println("Logout!", s)
	s.Trace()
	return nil
}

func (s *Session) Trace() {
	for i := 0; i < len(s.mails); i++ {
		m := s.mails[i]
		log.Println("======================================")
		log.Println("from:", m.GetSender())
		log.Println("to:", m.GetReceivers())
		log.Println("email:")
		message := m.GetMessage()
		log.Println("\theaders:", message.Header)
		log.Println("\tSubject:", message.Subject)
		log.Println("\tFrom:", message.From)
		log.Println("\tReplyTo:", message.ReplyTo)
		log.Println("\tTo:", message.To)
		log.Println("\tCc:", message.Cc)
		log.Println("\tBcc:", message.Bcc)
		log.Println("\tDate:", message.Date)
		log.Println("\tMessageID:", message.MessageID)
		log.Println("\tInReplyTo:", message.InReplyTo)
		log.Println("\tReferences:", message.References)
		log.Println("\tResent:")
		log.Println("\t\tFrom:", message.ResentFrom)
		log.Println("\t\tSender:", message.ResentSender)
		log.Println("\t\tTo:", message.ResentTo)
		log.Println("\t\tDate:", message.ResentDate)
		log.Println("\t\tCc:", message.ResentCc)
		log.Println("\t\tBcc:", message.ResentBcc)
		log.Println("\t\tMessageId:", message.ResentMessageID)
		log.Println("\tContentType:", message.ContentType)
		log.Println("\tHTMLBody:", message.HTMLBody)
		log.Println("\tTextBody:", message.TextBody)
		log.Println("\tAttachments:", message.Attachments)
		log.Println("\tEmbededFiles:", message.EmbeddedFiles)
	}
}

func NewSession() *Session {
	return &Session{
		current: nil,
		mails:   []*domain.Mail{},
	}
}
