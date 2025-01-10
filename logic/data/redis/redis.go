package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"code-comment-analyzer/config"
)

type SessionManager interface {
	SetSession(sessionID string, userID uint64) error
	GetSession(sessionID string) (uint64, error)
	ClearSession(sessionID string) error
	RefreshSession(sessionID string) error
	Close()
}

func NewSessionManager(cfg config.RedisConfig, sessionDuration uint32) SessionManager {
	return newRedisMaster(cfg, sessionDuration)
}

type redisMaster struct {
	ctx             context.Context
	client          *redis.Client
	sessionPrefix   string
	sessionDuration time.Duration
}

func newRedisMaster(config config.RedisConfig, sessionDuration uint32) *redisMaster {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DBNum,
	})
	return &redisMaster{
		ctx:             context.Background(),
		client:          client,
		sessionPrefix:   config.PrefixSession,
		sessionDuration: time.Duration(sessionDuration) * time.Minute,
	}
}

func (r *redisMaster) Close() {
	err := r.client.Close()
	if err != nil {
		panic(err)
	}
}

func (r *redisMaster) SetSession(sessionID string, userID uint64) error {
	sessionID = hashKey(sessionID)
	key := r.sessionPrefix + sessionID
	value := fmt.Sprintf("%d", userID)
	status := r.client.Set(r.ctx, key, value, r.sessionDuration)
	return status.Err()
}

func (r *redisMaster) GetSession(sessionID string) (uint64, error) {
	sessionID = hashKey(sessionID)
	key := r.sessionPrefix + sessionID
	result, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, fmt.Errorf("session not found")
		}
		return 0, err
	}
	userID, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing userID from session: %v", err)
	}
	return userID, nil
}

func (r *redisMaster) ClearSession(sessionID string) error {
	sessionID = hashKey(sessionID)
	key := r.sessionPrefix + sessionID
	result, err := r.client.Del(r.ctx, key).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return fmt.Errorf("no session was found to delete")
	}
	return nil
}

func (r *redisMaster) RefreshSession(sessionID string) error {
	sessionID = hashKey(sessionID)
	key := r.sessionPrefix + sessionID
	result, err := r.client.Expire(r.ctx, key, r.sessionDuration).Result()
	if err != nil {
		return err
	}
	if !result {
		return fmt.Errorf("no session found to refresh")
	}
	return nil
}
