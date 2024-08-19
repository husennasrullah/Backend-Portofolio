package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetServerProtocol() string
	GetServerHost() string
	GetServerPort() int
	GetPostgreSQLAddress() string
	GetPostgreSQLDefaultSchema() string
	///
	GetRedisHost() string
	GetRedisPort() int
	GetRedisDB() int
	GetRedisPassword() string
	GetRedisTimeout() int
	GetRedisRequestVolumeThreshold() int
	GetRedisSleepWindow() int
	GetRedisErrorPercentThreshold() int
	GetRedisMaxConcurrentRequests() int
}

func GenerateConfiguration(arguments string) {
	var err error
	enviName := os.Getenv("projectconfig")
	if arguments == "development" {
		temp := DevelopmentConfig{}
		err = gonfig.GetConf(enviName+"/config.json", &temp)
		if err != nil {
			log.Print("Error get config -> ", err)
			os.Exit(2)
		}
		err = envconfig.Process(enviName+"/config.json", &temp)
		if err != nil {
			log.Print("Error process config -> ", err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	} else {
		//setup if there is another config or for production config
	}
}
