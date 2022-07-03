package oss

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
)

var driver *Driver

type Driver interface {
	GetFile(string) (io.Reader, error)
	SaveFile(reader io.Reader) (string, error)
}

func init() {
	var err error
	// default configurations
	viper.SetDefault("oss.driver", "filesystem")
	// factory
	switch viper.GetString("oss.driver") {
	case "filesystem":
		driver, err = factoryFileBasedDriver()
	default:
		driver, err = nil, errors.New(fmt.Sprintf("unsupported oss driver:%s", viper.GetString("oss.driver")))
	}
	if err != nil {
		panic(err)
	}
}
