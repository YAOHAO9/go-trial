package handler

import (
	"fmt"
	"trial/rpc/response"
)

// Map handler函数仓库
type Map map[string]func(respCtx *response.RespCtx)

// Handler Handler
type Handler struct {
	Map Map
}

// Register handler
func (handler Handler) Register(name string, f func(respCtx *response.RespCtx)) {
	handler.Map[name] = f
}

// Exec 执行handler
func (handler Handler) Exec(respCtx *response.RespCtx) {

	f, ok := handler.Map[respCtx.Fm.Handler]
	if ok {
		f(respCtx)
	} else {
		respCtx.SendFailMessage(fmt.Sprintf("SysHandler %v 不存在", respCtx.Fm.Handler))
	}
}

var a = Handler{make(Map)}
