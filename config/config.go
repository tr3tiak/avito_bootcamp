package config

import (
	"os"
)

type ConfigDB struct {
	NameDB     string
	UserDB     string
	PasswordDB string
	Port       string
}

func InitConfigDB() *ConfigDB {
	cfg := ConfigDB{
		UserDB:     os.Getenv("user_db"),
		PasswordDB: os.Getenv("password_db"),
		NameDB:     os.Getenv("name_db"),
		Port:       os.Getenv("port"),
	}
	return &cfg
}
