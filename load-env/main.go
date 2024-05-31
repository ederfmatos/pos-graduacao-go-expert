package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
)

type Environment struct {
	DatabaseDriver         string `mapstructure:"DB_DRIVER"`
	DatabaseHost           string `mapstructure:"DB_HOST"`
	DatabasePort           string `mapstructure:"DB_PORT"`
	DatabaseUser           string `mapstructure:"DB_USER"`
	DatabasePassword       string `mapstructure:"DB_PASSWORD"`
	DatabaseName           string `mapstructure:"DB_NAME"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	JwtSecret              string `mapstructure:"JWT_SECRET"`
	JwtIssuer              string `mapstructure:"JWT_ISSUER"`
	JwtExpirationInSeconds int    `mapstructure:"JWT_EXPIRATION_IN_SECONDS"`
}

func LoadEnvironment(path string) *Environment {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	environment := Environment{}
	err = viper.Unmarshal(&environment)
	if err != nil {
		panic(err)
	}
	return &environment
}

func main() {
	environment := LoadEnvironment(".env")
	_ = json.NewEncoder(os.Stdout).Encode(environment)
}
