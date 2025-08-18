package app

import (
	"context"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	LocalHost   string = "0.0.0.0"
	DefaultPort string = "8080"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	App struct {
		Port string
		Host string
	}
}

func Configure(ctx context.Context) Config {
	once.Do(func() {
		_ = godotenv.Load()

		viper.SetDefault("App.Host", LocalHost)
		viper.SetDefault("App.Port", DefaultPort)
		viper.AutomaticEnv()

		if err := viper.Unmarshal(&config); err != nil {
			log.Panicf("unmarshaling config: %+v", err)
		}

		log.Print("configuration loaded")
	})

	return config
}
