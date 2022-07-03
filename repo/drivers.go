package repo

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
 * mysql driver
 */

func factoryMySQLConn() gorm.Dialector {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.database", "tmail")
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"), viper.GetString("db.password"),
		viper.GetString("db.host"), viper.GetInt("db.port"), viper.GetString("db.database"))
	return mysql.Open(dsn)
}

/*
 * sqlite driver
 */
func factorySQLiteConn() gorm.Dialector {
	viper.SetDefault("db.file", "./tmail.db")

	return sqlite.Open(viper.GetString("db.file"))
}
