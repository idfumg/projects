package config

import (
	"errors"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	dbHost     string
	dbPort     string
	dbUsername string
	dbTable    string
	dbPassword string
	dbSslMode  string
	dbUrl      string
	webPort    string
}

func New() (*Config, error) {
	gotenv.Load()

	config := &Config{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbUsername: os.Getenv("DB_USERNAME"),
		dbTable:    os.Getenv("DB_TABLE"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbSslMode:  os.Getenv("SSL_MODE"),
		dbUrl:      os.Getenv("DB_URL"),
		webPort:    os.Getenv("WEB_PORT"),
	}

	if len(config.GetDBHost()) == 0 && len(config.GetDBUrl()) == 0 {
		return nil, errors.New("we should provide the full db url or its parts")
	}

	if config.webPort == "" {
		config.webPort = "50051"
	}

	return config, nil
}

func (c *Config) GetDBHost() string     { return c.dbHost }
func (c *Config) GetDBPort() string     { return c.dbPort }
func (c *Config) GetDBUsername() string { return c.dbUsername }
func (c *Config) GetDBTable() string    { return c.dbTable }
func (c *Config) GetDBPassword() string { return c.dbPassword }
func (c *Config) GetDBSslMode() string  { return c.dbSslMode }
func (c *Config) GetDBUrl() string      { return c.dbUrl }
func (c *Config) GetWebPort() string    { return c.webPort }
