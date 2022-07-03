package repo

import "TMail/domain"

func CreateMail(mail domain.Mail) {
	conn.Create(mail)
}
