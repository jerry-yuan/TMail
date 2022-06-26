package domain

type Attachment struct {
	modelBase
	MailId   int64  `gorm:"column:mail_id;type:int(8);not null;index:idx_mail_id;comment:the mail this attachment belongs to"`
	Filename string `gorm:"column:name;type:varchar(512);not null;comment:the filename of this attachment"`
	MimeType string `gorm:"column:mime_type;type:varchar(32);not null;comment:the mime type of this attachment"`
	Size     int64  `gorm:"column:size;type:int(8);not null;comment:the size of attachment"`
	FileKey  string `gorm:"column:key;type:varchar(128);not null;comment:the file key in oss"`
}

func (a Attachment) TableName() string {
	return "attachment"
}
