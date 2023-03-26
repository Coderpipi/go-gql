package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"os"
)

var Config *Cfg

type App struct {
	Name string `toml:"name"`

	ServerPort string `toml:"server_port"`
}

type Cfg struct {
	BasePath string `env:"BASE_PATH"`
	Env      string `env:"ENV"`
	App      *App   `toml:"app"`
	Mysql    *Mysql `toml:"mysql"`
}

func InitConfig() {

	if len(os.Getenv("ENV")) == 0 {
		// if no ENV specified
		envFile := ".env"
		for i := 0; i < 10; i++ {
			if _, err := os.Stat(envFile); err != nil {
				envFile = "../" + envFile
				continue
			}
			break
		}
		// local debugging
		err := godotenv.Load(envFile)
		if err != nil {
			panic(err)
		}
	}
	Config = &Cfg{}
	if err := env.Parse(Config); err != nil {
		panic(err)
	}
	// read config xxx.toml
	_, err := toml.DecodeFile(fmt.Sprintf("%s/config-%s.toml", Config.BasePath, Config.Env), Config)
	if err != nil {
		panic(err)
	}
	// init mysql
	initMysql()
}
