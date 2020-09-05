package context

import (
	"fmt"
	"sync"

	"github.com/YAOHAO9/yh/rpc/message"
	"github.com/YAOHAO9/yh/rpc/session"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var mutex sync.Mutex

// RPCCtx response context
type RPCCtx struct {
	conn      *websocket.Conn
	kind      int
	requestID int
	handler   string
	Data      interface{} `json:",omitempty"`
	Session   *session.Session
}

// GenRespCtx 创建一个response上下文
func GenRespCtx(conn *websocket.Conn, rpcMsg *message.RPCMsg) *RPCCtx {
	return &RPCCtx{
		conn:      conn,
		kind:      rpcMsg.Kind,
		requestID: rpcMsg.RequestID,
		handler:   rpcMsg.Handler,
		Data:      rpcMsg.Data,
		Session:   rpcMsg.Session,
	}
}

// GetHandler 获取请求的Handler
func (rpcCtx *RPCCtx) GetHandler() string {
	return rpcCtx.handler
}

// GetRequestID 获取请求的GetRequestID
func (rpcCtx RPCCtx) GetRequestID() int {
	return rpcCtx.requestID
}

// SendMsg 消息发送失败
func (rpcCtx *RPCCtx) SendMsg(data interface{}, code int) {
	// Notify的消息，不通知成功
	if rpcCtx.requestID == 0 {
		logrus.Warn("Notify不需要回复消息")
		return
	}
	// 重复回复
	if rpcCtx.requestID == -1 {
		logrus.Warn("请勿重复回复消息")
		return
	}

	// response
	rpcResp := message.RPCResp{
		Kind:      rpcCtx.kind + 10000,
		RequestID: rpcCtx.requestID,
		Code:      code,
		Data:      data,
	}
	// 标记为已回复消息
	rpcCtx.requestID = -1

	mutex.Lock()
	err := rpcCtx.conn.WriteMessage(message.TypeEnum.TextMessage, rpcResp.ToBytes())
	if err != nil {
		logrus.Error(err)
	}
	mutex.Unlock()
}

// ToString 格式化消息
func (rpcCtx RPCCtx) ToString() string {
	if rpcCtx.kind == message.KindEnum.RPC {
		return fmt.Sprintf("RPC RequestID: %d, Data: %+v", rpcCtx.requestID, rpcCtx.Data)
	}
	return fmt.Sprintf("Handler RequestID: %d, Data: %+v", rpcCtx.requestID, rpcCtx.Data)
}
