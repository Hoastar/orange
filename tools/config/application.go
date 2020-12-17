package config

import "github.com/spf13/viper"

type Application struct {
	Readtimeout   int
	Writertimeout int
	Name          string
	Domain        string
	Host          string
	IsHttps       bool
	Mode          string
	Port          string
	JwtSecret     string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Readtimeout: cfg.GetInt("readtimeout"),
		Writertimeout: cfg.GetInt("writertimeout"),
		Domain: cfg.GetString("domain"),
		Host: cfg.GetString("host"),
		IsHttps: cfg.GetBool("ishttps"),
		JwtSecret: cfg.GetString("jwtSecret"),
		Mode: cfg.GetString("mode"),
		Name: cfg.GetString("name"),
		Port: setPortDefault(cfg),

	}
}

var ApplicationConfig = new(Application)

func setPortDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8080"
	} else {
		return cfg.GetString("port")
	}
}

func isHttpsDefault(cfg *viper.Viper) bool {
	if cfg.GetString("ishttps") == "" || cfg.GetBool("ishttps") == false {
		return false
	} else {
		return true
	}
}
