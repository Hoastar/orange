package config

import "github.com/spf13/viper"

type Database struct {
	Dbtype string
	Host string
	Port int
	Name string
	Password string
	Username string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port: cfg.GetInt("port"),
		Dbtype: cfg.GetString("dbtype"),
		Host: cfg.GetString("host"),
		Name: cfg.GetString("Name"),
		Password: cfg.GetString("password"),
		Username: cfg.GetString("username"),
	}
}

var DatabaseConfig = new(Database)