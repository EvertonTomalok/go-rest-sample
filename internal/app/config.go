package app

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const LocalHost string = "0.0.0.0"
const DefaultPort string = "8080"

type Config struct {
	App struct {
		Port string
		Host string
	}
}

func Configure(ctx context.Context) Config {
	_ = godotenv.Load()

	viper.SetDefault("App.Host", LocalHost)
	viper.SetDefault("App.Port", DefaultPort)
	viper.AutomaticEnv()

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("unmarshaling config: %+v", err)
	}

	log.Print("configuration loaded")

	return cfg
}
