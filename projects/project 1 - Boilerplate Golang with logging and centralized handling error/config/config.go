package config

import (
	"fmt"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	LocalPort   string `mapstructure:"LOCAL_PORT"`
	AppName     string `mapstructure:"APP_NAME"`
	AppVersion  string `mapstructure:"APP_VERSION"`
	Build       int    `mapstructure:"BUILD"`
	Mode        string `mapstructure:"MODE"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      string `mapstructure:"DB_PORT"`
	SSLMode     string `mapstructure:"SSL_MODE"`
	TZ          string `mapstructure:"TZ"`
	RedisHost   string `mapstructure:"REDISHOST"`
	RedisPort   int    `mapstructure:"REDISPORT"`
	RedisNumber int    `mapstructure:"REDISDBNUMBER"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Cannot load file .env: ", logrus.Fields{
			"error message": err,
		})
		panic(err)
	}

	// Bind environment variables to Viper
	viper.AutomaticEnv()
	bindEnv()

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	// set gcp credentials temp filename
	logrus.Println("initializing GCP Credentials")

	return &config, nil
}

func bindEnv() {
	// Explicitly bind each environment variable
	viper.BindEnv("LOCAL_PORT")
	viper.BindEnv("APP_NAME")
	viper.BindEnv("APP_VERSION")
	viper.BindEnv("BUILD")
	viper.BindEnv("MODE")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("SSL_MODE")
	viper.BindEnv("TZ")
	viper.BindEnv("REDISHOST")
	viper.BindEnv("REDISPORT")
	viper.BindEnv("REDISDBNUMBER")
}
