package service

import (
	"crudproduct/attribute"
	"crudproduct/util"
	"fmt"
	"time"
)

func WriteToRedis(key, mdl string) {
	rds := attribute.ServerAttribute.RedisClient.Set(key, util.StructToJSON(mdl), 5*time.Hour)
	fmt.Println(rds)
}

func GetFromRedis(key string) string {
	return attribute.ServerAttribute.RedisClient.Get(key).Val()
}

func DeleteRedis() {
	for {
		keys, _, err := attribute.ServerAttribute.RedisClient.Scan(0, "", 1000).Result()
		if err != nil {
			fmt.Println("Error Redis Scan : ", err.Error())
			return
		}

		if len(keys) == 0 {
			return
		}
		for i := 0; i < len(keys); i++ {
			rsp := attribute.ServerAttribute.RedisClient.Del(keys[i])
			if rsp.Err() != nil {
				fmt.Println(rsp.Err().Error())
				return
			}
		}
	}
}
