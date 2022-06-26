package domain

type Mail struct {
	modelBase
	Sender         string            `gorm:"column:sender;not null;type:varchar(64);index:idx_sender;comment:mail sender name"`
	SenderDomain   string            `gorm:"column:sender_domain;not null;type:varchar(255);index:idx_sender;comment:mail sender's domain"`
	Receiver       string            `gorm:"column:receiver;not null;type:varchar(64);index:idx_receiver;comment:mail receiver's name"`
	ReceiverDomain string            `gorm:"column:receiver_domain;not null;type:varchar(255);index:idx_receiver;comment:mail receiver's domain"`
	Subject        string            `gorm:"column:subject;not null;type:varchar(1024);default: ;comment:the subject of mail"`
	Headers        map[string]string `gorm:"column:headers;not null;type:json;comment:the headers of mail;serializer:json"`
	Body           string            `gorm:"column:body;not null;type:text;comment:the body of mail"`
}

func (m Mail) TableName() string {
	return "mail"
}
