package middleware

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/data/redis"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/util"
)

type option func(*optionParams)

type optionParams struct {
	sessionManager redis.SessionManager
}

func WithSessionManager(s redis.SessionManager) option {
	return func(params *optionParams) {
		params.sessionManager = s
	}
}

type routerGroup struct {
	basePath       string
	mux            *http.ServeMux
	sessionManager redis.SessionManager
}

func NewRouterGroup(basePath string, mux *http.ServeMux, opts ...option) *routerGroup {
	optParams := &optionParams{}
	for _, opt := range opts {
		opt(optParams)
	}
	return &routerGroup{
		basePath:       basePath,
		mux:            mux,
		sessionManager: optParams.sessionManager,
	}
}

func (rg *routerGroup) Get(relativePath string, getHandler GetHandler, middleOps ...MiddleOpFunc) {
	rg.registerRouter(Get, relativePath, getHandler, middleOps...)
}

func (rg *routerGroup) Post(relativePath string, getHandler GetHandler, middleOps ...MiddleOpFunc) {
	rg.registerRouter(Post, relativePath, getHandler, middleOps...)
}

func (rg *routerGroup) registerRouter(method string, relativePath string, getHandler GetHandler, middleOps ...MiddleOpFunc) {
	if rg.mux == nil {
		panic("mux is nil")
	}
	realPath := rg.basePath + relativePath
	// 请注意调用顺序，先调用最外层的中间件，再调用最内层的中间件，这里需要“反过来”写
	handlerFunc := defaultMiddleOp(getHandler)
	for i := len(middleOps) - 1; i >= 0; i-- {
		handlerFunc = middleOps[i](rg, handlerFunc)
	}
	formatHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		util.WithRecover(func() {
			switch method {
			case Get:
				handlerFunc = enforceGet(handlerFunc)
			case Post:
				handlerFunc = enforcePost(handlerFunc)
			default:
				log.Printf("invalid method: %s", method)
			}
			handlerFunc(w, r, nil)
		}, util.WithPanicHandler(func(err interface{}) {
			panicHandler(err, method, realPath)
		}))
	}
	rg.mux.HandleFunc(realPath, formatHandlerFunc)
}

func defaultMiddleOp(getHandler GetHandler) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		h := getHandler(w, r, extractor)
		h.Handle()
	}
}

func enforceGet(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != Get {
			log.Printf("request method must be %s", Get)
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeMustBeGet, fmt.Sprintf("request method must be %s", Get))
			return
		}
		handlerFunc(w, r, extractor)
	}
}

func enforcePost(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != Post {
			log.Printf("request method must be %s", Post)
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeMustBePost, fmt.Sprintf("request method must be %s", Post))
			return
		}
		handlerFunc(w, r, extractor)
	}
}

func (rg *routerGroup) getSessionManager() redis.SessionManager {
	if rg.sessionManager == nil {
		panic("session manager is not set")
	}
	return rg.sessionManager
}

func panicHandler(err interface{}, method string, realPath string) {
	log.Printf("panic: %v, method: %s, realPath: %s", err, method, realPath)
}
