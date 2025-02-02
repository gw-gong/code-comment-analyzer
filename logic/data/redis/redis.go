package redis

import (
	"code-comment-analyzer/util"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"code-comment-analyzer/config"
)

type redisClient struct {
	ctx             context.Context
	client          *redis.Client
	sessionPrefix   string
	sessionDuration time.Duration
}

func newRedisMaster(config config.RedisConfig, sessionDuration uint32) *redisClient {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DBNum,
	})
	return &redisClient{
		ctx:             context.Background(),
		client:          client,
		sessionPrefix:   config.PrefixSession,
		sessionDuration: time.Duration(sessionDuration) * time.Minute,
	}
}

func (r *redisClient) Close() {
	err := r.client.Close()
	if err != nil {
		panic(err)
	}
}

func (r *redisClient) SetSession(userID uint64, token string) error {
	key := r.transformUserIDToKey(userID)
	status := r.client.Set(r.ctx, key, token, r.sessionDuration)
	return status.Err()
}

func (r *redisClient) GetSession(userID uint64) (token string, err error) {
	key := r.transformUserIDToKey(userID)
	token, err = r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("session not found")
		}
		return "", err
	}
	return token, nil
}

func (r *redisClient) ClearSession(userID uint64) error {
	key := r.transformUserIDToKey(userID)
	result, err := r.client.Del(r.ctx, key).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return fmt.Errorf("no session was found to delete")
	}
	return nil
}

func (r *redisClient) RefreshSession(userID uint64) error {
	key := r.transformUserIDToKey(userID)
	result, err := r.client.Expire(r.ctx, key, r.sessionDuration).Result()
	if err != nil {
		return err
	}
	if !result {
		return fmt.Errorf("no session found to refresh")
	}
	return nil
}

func (r *redisClient) transformUserIDToKey(userID uint64) string {
	return r.sessionPrefix + util.FormatUserIDStr(userID)
}
