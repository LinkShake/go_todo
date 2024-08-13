package redis

import (
	"context"
	"errors"

	"github.com/gofiber/storage/redis/v3"
	"github.com/google/uuid"
)

var RedisClient = redis.New()

func GetUserId(sid string) (string, error) {
	val, err := RedisClient.Conn().Get(context.TODO(), sid).Result()
	if err != nil {
		if val != "" {
			panic(err)
		}
		return "", errors.New("invalid session id")
	}

	return val, nil
}

func UpdateUserSessionIdx(userId string) (string, string, error) {
	sid := uuid.New()
	val, err := RedisClient.Conn().Set(context.TODO() ,sid.String(), userId, 0).Result()
	if err != nil {
		panic(err)
	}

	return sid.String(), val, nil
}

func RemoveSessionId(sid string) error {
	_, err := RedisClient.Conn().Del(context.TODO(), sid).Result()
	return err
}