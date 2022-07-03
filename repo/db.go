package repo

import (
	"TMail/domain"
	"errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

var conn *gorm.DB

const (
	MySQL  = "mysql"
	SQLite = "sqlite"
)

func init() {
	// default configurations
	viper.SetDefault("db.driver", "mysql")

	// connect to database
	var err error
	conn, err = factoryGormConn()
	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	// initialize database
	migrator := conn.Migrator()

	err = migrator.AutoMigrate(&domain.Mail{}, &domain.Attachment{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

}

func factoryGormConn() (db *gorm.DB, err error) {
	dialetor, err := factoryGormDialetor()
	if err != nil {
		return
	}
	db, err = gorm.Open(dialetor, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return
}

func factoryGormDialetor() (dialector gorm.Dialector, err error) {
	driver := strings.ToLower(viper.GetString("db.driver"))

	switch driver {
	case MySQL:
		dialector = factoryMySQLConn()
	case SQLite:
		dialector = factorySQLiteConn()
	default:
		err = errors.New("unsupported database driver:" + driver)
	}
	return
}

func Transactional(transactionOps func(tx *gorm.DB) error) error {
	return conn.Transaction(transactionOps)
}
