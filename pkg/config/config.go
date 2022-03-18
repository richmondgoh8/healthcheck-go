package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APP
}

type APP struct {
	LOGGING_LEVEL string
}

func InitReader() Config {
	var conf Config
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal("error reading config")
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
