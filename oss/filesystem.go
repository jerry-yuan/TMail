package oss

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"regexp"
)

type fsObjectStorage struct {
	basePath    string
	prefixDepth int
}

func factoryFileBasedDriver() (driver Driver, err error) {
	viper.SetDefault("oss.base", "/var/lib/tmail")
	viper.SetDefault("oss.prefixDepth", 2)

	driver = &fsObjectStorage{
		basePath: viper.GetString("oss.base"),
	}
	// initialize
	err = driver.(*fsObjectStorage).initialize()
	return
}

func (f *fsObjectStorage) GetFile(fileKey string) (io.Reader, error) {
	filePath, err := f.toFilePath(fileKey)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil
}

func (f *fsObjectStorage) SaveFile(reader io.Reader) (string, error) {
	hasher := md5.New()
	buf := make([]byte, 1024)

	for {
		bytesRead, err := reader.Read(buf)
		if bytesRead < 1 && err != nil {
			return "", err
		}
		_, err = hasher.Write(buf[bytesRead:])
		if err != nil {
			// to be continue
		}
	}
}

func (f *fsObjectStorage) toFilePath(fileKey string) (string, error) {
	match, err := regexp.Match("[0-9a-f]{32}", []byte(fileKey))
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("invalid MD5 file key")
	}

	pathParts := []string{f.basePath}

	for i := 0; i < f.prefixDepth; i++ {
		pathParts = append(pathParts, fileKey[0:2])
		fileKey = fileKey[2:]
	}

	return path.Join(pathParts...), nil
}

func (f *fsObjectStorage) initialize() error {
	_, err := os.Stat(f.basePath)
	if err != nil {
		return err
	}
	// create folders
	depth := viper.GetInt("oss.prefixDepth")
	if depth >= 15 || depth < 0 {
		return errors.New("depth must be in range (0,14]")
	}
	return makeTree(f.basePath, depth)
}

func makeTree(basePath string, depth int) (err error) {
	if depth < 0 {
		return nil
	}
	for i := 0; i < 256; i++ {
		dirId := fmt.Sprintf("%02x", i)
		dirPath := path.Join(basePath, dirId)

		err = os.MkdirAll(dirId, 0644)
		if err != nil {
			return
		}
		err = makeTree(dirPath, depth-1)
		if err != nil {
			return
		}
	}
	return nil
}
