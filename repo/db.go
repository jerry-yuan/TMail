package repo

import (
	"TMail/domain"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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
	initDriverDefaults()

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

func initDriverDefaults() {
	driver := strings.ToLower(viper.GetString("db.driver"))
	switch driver {
	case MySQL:
		initMysqlDefaultConfigurations()
	case SQLite:
		initSQLiteDefaultConfigurations()
	}
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

/*
 * mysql driver
 */
func initMysqlDefaultConfigurations() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.database", "tmail")
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "")
}

func factoryMySQLConn() gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetInt("db.port"), viper.GetString("db.database"))
	return mysql.Open(dsn)
}

/*
 * sqlite driver
 */
func initSQLiteDefaultConfigurations() {
	viper.SetDefault("db.file", "./tmail.db")
}

func factorySQLiteConn() gorm.Dialector {
	return sqlite.Open(viper.GetString("db.file"))
}
