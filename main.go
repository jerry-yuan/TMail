package main

import (
	smtp2 "TMail/smtp"
	"github.com/emersion/go-smtp"
	"log"
	"time"

	_ "TMail/domain"
)

func main() {
	be := &smtp2.Server{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "t.jerryzone.cn"
	s.ReadTimeout = 3600 * time.Second
	s.WriteTimeout = 3600 * time.Second
	//s.MaxMessageBytes = 1024 * 1024
	//s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
