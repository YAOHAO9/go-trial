package rpc

import (
	"github.com/YAOHAO9/yh/rpc/client/clientmanager"
	"github.com/YAOHAO9/yh/rpc/message"
	"github.com/YAOHAO9/yh/rpc/session"
	"github.com/sirupsen/logrus"
)

type notify struct{}

// ToServer Rpc到指定的Server
func (n notify) ToServer(serverID string, session *session.Session, handler string, data interface{}) {

	rpcMsg := &message.RPCMsg{
		Kind:    message.KindEnum.RPC,
		Handler: handler,
		Data:    data,
		Session: session,
	}

	rpcClient := clientmanager.GetClientByID(serverID)
	if rpcClient == nil {
		logrus.Error("Rpc Notify(ToServer) 消息发送失败，没有找到对应的服务器 handler:", handler)
		return
	}

	rpcClient.SendRPCNotify(session, rpcMsg)
}

// ByKind Rpc到指定的Server
func (n notify) ByKind(serverKind string, session *session.Session, handler string, data interface{}) {
	rpcMsg := &message.RPCMsg{
		Kind:    message.KindEnum.RPC,
		Handler: handler,
		Data:    data,
		Session: session,
	}

	// 根据类型转发
	rpcClient := clientmanager.GetClientByRouter(serverKind, rpcMsg)
	if rpcClient == nil {
		logrus.Error("Rpc Notify(ByKind) 消息发送失败，没有找到对应的服务器 handler:", handler)
		return
	}
	rpcClient.SendRPCNotify(session, rpcMsg)
}

type request struct{}

// ToServer Rpc到指定的Server
func (req request) ToServer(serverID string, session *session.Session, handler string, data interface{}, f func(rpcResp *message.RPCResp)) {

	rpcMsg := &message.RPCMsg{
		Kind:    message.KindEnum.RPC,
		Handler: handler,
		Data:    data,
		Session: session,
	}

	rpcClient := clientmanager.GetClientByID(serverID)
	if rpcClient == nil {
		logrus.Error("Rpc Request(ToServer) 消息发送失败，没有找到对应的服务器 handler:", handler)
		return
	}

	rpcClient.SendRPCRequest(session, rpcMsg, f)
}

// ByKind Rpc到指定的Server
func (req request) ByKind(serverKind string, session *session.Session, handler string, data interface{}, f func(rpcResp *message.RPCResp)) {
	rpcMsg := &message.RPCMsg{
		Kind:    message.KindEnum.RPC,
		Handler: handler,
		Data:    data,
		Session: session,
	}
	// 根据类型转发
	rpcClient := clientmanager.GetClientByRouter(serverKind, rpcMsg)
	if rpcClient == nil {
		logrus.Error("Rpc Request(ByKind) 消息发送失败，没有找到对应的服务器 handler:", handler)
		return
	}
	rpcClient.SendRPCRequest(session, rpcMsg, f)
}

// Notify 实例
var Notify notify

// Request 实例
var Request request
