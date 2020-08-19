package rpchandler

import (
	"trial/rpc/handler"
)

// Handler RPCHandler
type Handler struct {
	*handler.Handler
}

// Manager return RPCHandler
var Manager = &Handler{&handler.Handler{Map: make(handler.Map)}}
