package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	Database struct {
		Dialect   string `default:"mysql"`
		Host      string `default:"127.0.0.1"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"true"`
	}
	Jwt struct {
		Secret     string
		ExpireTime int16
	}
}

const (
	DEV = "develop"
	PRD = "production"
	DOC = "docker"
)

func Load() (*Config, string) {
	var env *string
	if value := os.Getenv("GO_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse()
	}

	config := &Config{}
	if err := configor.Load(config, "application."+*env+".yml"); err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(2)
	}
	os.Setenv("JwtSecret", config.Jwt.Secret)
	return config, *env
}
