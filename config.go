package main

import "github.com/spf13/viper"

type AppConfig struct {
	PORT int
	ADDRESS string
	
	DBHost string
	DBName string
	DBPort int
	DBUser string
	DBPass string
}

func NewAppConfig() *AppConfig {

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return &AppConfig{

		PORT: viper.GetInt("PORT"),
		ADDRESS: viper.GetString("ADDRESS"),

		DBName: viper.GetString("DB_NAME"),
		DBPort: viper.GetInt("DB_PORT"),
		DBPass: viper.GetString("DB_PASS"),
		DBUser: viper.GetString("DB_USER"),
		DBHost: viper.GetString("DB_HOST"),
	}
}
