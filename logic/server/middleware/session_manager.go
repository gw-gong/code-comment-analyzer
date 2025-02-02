package middleware

import (
	"code-comment-analyzer/data/redis"
)

var sessionManager redis.SessionManager

func RegisterSessionManager(s redis.SessionManager) {
	sessionManager = s
}
