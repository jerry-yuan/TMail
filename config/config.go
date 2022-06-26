package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	viper.SetConfigName("tmail")
	viper.AddConfigPath("/etc/tmail")
	viper.AddConfigPath("$HOME/.tmail")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(filepath.Dir(os.Args[0]))
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
