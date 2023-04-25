package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App    app
	Server server
	Db     db
	Mongo  mongo
}

type app struct {
	Name    string
	Author  string
	Email   string
	Version string
}

type server struct {
	Port uint64
}

type db struct {
	Host     string
	Port     uint64
	Name     string
	UserName string
	Password string
}

type mongo struct {
	Db         string
	Collection string
	Host       string
	Port       uint64
	Username   string
	Password   string
}

const cfgFilePath string = "./config.toml"

var Cfg Config

func init() {
	_, _ = toml.DecodeFile(cfgFilePath, &Cfg)
	log.Println(Cfg)
	log.Println("Config Init Over!")
}
