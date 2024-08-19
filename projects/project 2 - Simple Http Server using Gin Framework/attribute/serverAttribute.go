package attribute

import (
	"crudproduct/config"
	"crudproduct/util"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/nexsoftgit/go-redis"
	"gorm.io/gorm"
	"strconv"
)

var ServerAttribute serverAttribute

type serverAttribute struct {
	GormClient  *gorm.DB
	RedisClient *redis.Client
}

func SetServerAttribute() {
	//set gorm connection to DB
	ServerAttribute.GormClient = util.ConnectDatabase(config.ApplicationConfiguration.GetPostgreSQLAddress())

	redisHost := config.ApplicationConfiguration.GetRedisHost()
	redisDB := config.ApplicationConfiguration.GetRedisDB()
	redisPassword := config.ApplicationConfiguration.GetRedisPassword()
	redisPort := config.ApplicationConfiguration.GetRedisPort()
	redisTimeout := config.ApplicationConfiguration.GetRedisTimeout()
	redisVolumeThreshold := config.ApplicationConfiguration.GetRedisRequestVolumeThreshold()
	redisSleepWindow := config.ApplicationConfiguration.GetRedisSleepWindow()
	redisErrorPercentThreshold := config.ApplicationConfiguration.GetRedisErrorPercentThreshold()
	redisMaxConcurrentRequest := config.ApplicationConfiguration.GetRedisMaxConcurrentRequests()

	optCB := &hystrix.CommandConfig{
		Timeout:                redisTimeout,
		RequestVolumeThreshold: redisVolumeThreshold,
		SleepWindow:            redisSleepWindow,
		ErrorPercentThreshold:  redisErrorPercentThreshold,
		MaxConcurrentRequests:  redisMaxConcurrentRequest,
	}

	ServerAttribute.RedisClient = getRedisClient(redisHost, redisPort, redisDB, redisPassword, optCB)
}

func getRedisClient(host string, port int, db int, password string, optCB *hystrix.CommandConfig) *redis.Client {
	redisAddress := host + ":" + strconv.Itoa(port)
	opts := &redis.Options{
		CircuitBreaker: optCB,
		Addr:           redisAddress,
		Password:       password,
		DB:             db,
	}

	return redis.NewClient(opts)
}
