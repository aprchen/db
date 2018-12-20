package db

import (
	"errors"
	"os"
)

type MysqlMessage struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

func (mm *MysqlMessage) Check() error {
	if len(mm.Host) == 0 {
		mm.Host = "127.0.0.1"
	}
	if mm.Port == 0 {
		mm.Port = 3306
	}
	if len(mm.User) == 0 {
		mm.User = "root"
	}
	if len(mm.Name) == 0 {
		return errors.New("mysql need db name")
	}
	return nil
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func MysqlMessageFromEnv() MysqlMessage {
	return MysqlMessage{
		Host:     EnvString("MYSQL_HOST", ""),
		Name:     EnvString("MYSQL_NAME", ""),
		User:     EnvString("MYSQL_USER", ""),
		Password: EnvString("MYSQL_PASSWORD", ""),
	}
}
