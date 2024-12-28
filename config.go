package server

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string `mapstructure:"app_name"`
	Port     string `mapstructure:"port"`
	RedisURL string `mapstructure:"redis_url"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("BASE_SERVER")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("app_name", "Base Server")
	viper.SetDefault("port", "8080")
	viper.SetDefault("redis_url", "localhost:6379")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using defaults: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	return &cfg
}
