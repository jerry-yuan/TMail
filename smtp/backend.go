package smtp

import (
	"github.com/emersion/go-smtp"
	"log"
)

// The Server implements SMTP server methods.
type Server struct{}

func (bkd *Server) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	log.Println("Login:", state, username, password)
	return NewSession(), nil
}

func (bkd *Server) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	log.Println("AnonymousLogin:", state)
	return NewSession(), nil
}
