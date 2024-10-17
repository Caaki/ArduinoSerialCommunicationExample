package initializers

import (
	"SerialArduinoCommunication/configuration"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

//var EmailSender *emailSender.EmailSender

//func InitializeEmailSender() {
//
//	config := configuration.GlobalConfiguration
//
//	EmailSender = &emailSender.EmailSender{
//		Username: config.EmailSender.Username,
//		Server:   config.EmailSender.Server,
//		Port:     config.EmailSender.Port,
//		Password: config.EmailSender.Password,
//		From:     config.EmailSender.From,
//	}
//
//}

func InitializeRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configuration.GlobalConfiguration.Redis.Address,
		Password: configuration.GlobalConfiguration.Redis.Password,
		DB:       1,
	})

	ping, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return
	}

	fmt.Println(ping)

}
