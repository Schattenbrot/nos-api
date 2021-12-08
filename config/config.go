package config

import (
	"log"

	"github.com/Schattenbrot/nos-api/models"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}

type Application struct {
	Version string
	Logger  *log.Logger
	Models  models.Models
}

var Cfg Config
var App *Application
