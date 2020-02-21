package app

import (
	"trial/rpc/filter"
	"trial/rpc/filter/rpcfilter"
	"trial/rpc/handler"
	"trial/rpc/handler/rpchandler"
	"trial/rpc/msg"

	"github.com/gorilla/websocket"
)

// RegisterHandler 注册Handler
func RegisterHandler(name string, f func(respConn *websocket.Conn, fm *msg.ForwardMessage)) {
	handler.Manager().Register(name, f)
}

// RegisterRPCHandler 注册Handler
func RegisterRPCHandler(name string, f func(respConn *websocket.Conn, fm *msg.ForwardMessage)) {
	rpchandler.Manager().Register(name, f)
}

// RegisterHandlerBeforeFilter 注册before filter of handler
func RegisterHandlerBeforeFilter(f func(respConn *websocket.Conn, fm *msg.ForwardMessage) (next bool)) {
	filter.BeforeFilterManager().Register(f)
}

// RegisterHandlerAfterFilter 注册after filter of handler request
func RegisterHandlerAfterFilter(f func(rm *msg.ResponseMessage) (next bool)) {
	filter.AfterFilterManager().Register(f)
}

// RegisterRPCBeforeFilter 注册before filter of rpc
func RegisterRPCBeforeFilter(f func(respConn *websocket.Conn, fm *msg.ForwardMessage) (next bool)) {
	rpcfilter.BeforeFilterManager().Register(f)
}

// RegisterRPCAfterFilter 注册after filter of rpc request
func RegisterRPCAfterFilter(f func(rm *msg.ResponseMessage) (next bool)) {
	rpcfilter.AfterFilterManager().Register(f)
}
