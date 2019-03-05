package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	//HTTPS bool `default:"false" env:"HTTPS"`
	//Port  uint `default:"7000" env:"PORT"`
	DB struct {
		Name     string `env:"DBName" default:"mteg_cms"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBID"  default:"root"`
		Password string `env:"DBPW"  default:"_media_"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml"); err != nil {
		panic(err)
	}
}
