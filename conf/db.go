package conf

import (
	"fmt"
	"os"
	"strconv"
)

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSL      string
	Timezone string
}

func NewDBConfig() *DBConfig {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &DBConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSL:      os.Getenv("DB_SSL_MODE"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}
}

func (c DBConfig) ConnectionUri() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=%s", c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSL, c.Timezone)
}
