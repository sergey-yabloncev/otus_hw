package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Logger LoggerConf
	DB     DBConf
	Server ServerConf
}

type LoggerConf struct {
	Level string
	Path  string
}

type ServerConf struct {
	Port string
}

type DBConf struct {
	User     string
	Password string
	Host     string
	Port     uint64
	Name     string
}

func NewConfig(path string) (Config, error) {
	config := Config{}

	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Println(err)
		return config, err
	}

	return config, nil
}
