package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pauluswi/bigben/exception"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(fileNames ...string) Config {
	err := godotenv.Load(fileNames...)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	return &configImpl{}
}
