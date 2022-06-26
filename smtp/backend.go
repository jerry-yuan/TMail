package smtp

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	// default configurations
	viper.SetDefault("smtp.bind", "0.0.0.0")
	viper.SetDefault("smtp.port", 25)
	viper.SetDefault("smtp.domain", "tmail.local")
	viper.SetDefault("smtp.conn.readTimeout", 0)
	viper.SetDefault("smtp.conn.writeTimeout", 0)
	viper.SetDefault("smtp.message.maxSize", 0)
	// initialize smtp server

	be := &Server{}

	s := smtp.NewServer(be)

	s.Addr = fmt.Sprintf("%s:%d", viper.GetString("smtp.bind"), viper.GetInt("smtp.port"))
	s.Domain = viper.GetString("smtp.domain")
	s.ReadTimeout = time.Duration(viper.GetInt32("smtp.conn.readTimeout")) * time.Second
	s.WriteTimeout = time.Duration(viper.GetInt32("smtp.conn.writeTimeout")) * time.Second
	s.MaxMessageBytes = viper.GetInt("smtp.maxMessageBytes")
	s.MaxRecipients = viper.GetInt("smtp.message.maxSize")
	s.AllowInsecureAuth = true
	go func() {
		log.Printf("Starting server at %s with domain %s", s.Addr, s.Domain)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

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
