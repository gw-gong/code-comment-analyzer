package redis

import "code-comment-analyzer/config"

type SessionManager interface {
	SetSession(userID uint64, token string) error
	GetSession(userID uint64) (token string, err error)
	ClearSession(userID uint64) error
	RefreshSession(userID uint64) error
	Close()
}

func NewSessionManager(cfg config.RedisConfig, sessionDuration uint32) SessionManager {
	return newRedisMaster(cfg, sessionDuration)
}
