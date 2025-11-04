package config

import (
	"github.com/javadifa/socialmedia/repository/postgresql"
	"github.com/javadifa/socialmedia/service/authservice"
)

type HTTPServer struct {
	Port int
}
type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	Postgresql postgresql.Config
}
