package domain

type modelBase struct {
	ID int64 `gorm:"column:id;not null;type:int(8) primary key auto_increment;comment:id"`
}
