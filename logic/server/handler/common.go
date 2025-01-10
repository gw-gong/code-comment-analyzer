package handler

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/server/middleware"
)

func getUserIDFromRequestCtx(w http.ResponseWriter, r *http.Request) (uint64, error) {
	userID := r.Context().Value(middleware.CtxKeyUserID)
	if userID == nil {
		return 0, fmt.Errorf("missing user id")
	}
	return userID.(uint64), nil
}
