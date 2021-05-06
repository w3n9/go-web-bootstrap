package config

import (
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/util/config"
)

type application struct {
	Server struct {
		Host string
		Port string
	}
}

var Application application

func setDefault() {
	Application.Server.Host = "0.0.0.0"
	Application.Server.Port = "8888"
}

func init() {
	baseConfigPath := config.ReadConfigDirFromProgramArgument()
	setDefault()
	err := config.Load(baseConfigPath+"/application.yml", &Application)
	if err != nil {
		panic(err)
	}
}
