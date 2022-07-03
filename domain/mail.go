package domain

import "time"

type Mail struct {
	modelBase
	Sender     string            `gorm:"column:sender;not null;type:varchar(64);index:idx_sender;comment:mail sender name"`
	Receivers  []string          `gorm:"column:receivers;not null;type:varchar(1024);index:idx_receiver;comment:mail receiver's name;serializer:json"`
	Subject    string            `gorm:"column:subject;not null;type:varchar(1024);default: ;comment:the subject of mail"`
	From       []string          `gorm:"column:from;not null;type:varchar(1024);comment:mail from;serializer:json"`
	ReplyTo    []string          `gorm:"column:reply_to;not null;type:varchar(1024);comment:receivers used to reply this mail;serializer:json"`
	CarbonCopy []string          `gorm:"column:carbon_copy;not null;type:varchar(1024);comment:mail copyed to;serializer:json"`
	Headers    map[string]string `gorm:"column:headers;not null;type:json;comment:the headers of mail;serializer:json"`
	Body       string            `gorm:"column:body;not null;type:text;comment:the body of mail"`
	Date       time.Time         `gorm:"column:date;not null;type:datetime;index:idx_date;comment:the mail arrival time"`
}

func (m Mail) TableName() string {
	return "mail"
}
