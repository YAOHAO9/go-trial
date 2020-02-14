package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
	"trial/config"
	"trial/rpc/msg"
	"trial/rpc/msgtype"

	"github.com/gorilla/websocket"
)

var requestMap = make(map[int]func(data []byte))

// RPCClient websocket client 连接信息
type RPCClient struct {
	clientConn   *websocket.Conn
	serverConfig *config.ServerConfig
}

// SendRPCNotify send rpc message
func (client RPCClient) SendRPCNotify(data []byte) {
	client.clientConn.WriteMessage(msgtype.TextMessage, data)
}

// SendRPCRequest send rpc message
func (client RPCClient) SendRPCRequest(msgIndex int, data []byte, cb func(data []byte)) {
	client.clientConn.WriteMessage(msgtype.TextMessage, data)
	requestMap[msgIndex] = cb
}

// StartClient websocket client
func StartClient(serverConfig *config.ServerConfig, zkSessionTimeout time.Duration) *RPCClient {

	// Dialer
	dialer := websocket.Dialer{}
	urlString := url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprint(serverConfig.Host, ":", serverConfig.Port),
		Path:     "/rpc",
		RawQuery: fmt.Sprint("token=", serverConfig.Token),
	}

	var e error
	// 当前尝试次数
	tryTimes := 0
	// 最大尝试次数
	maxTryTimes := int(50 + zkSessionTimeout/100/time.Millisecond)

	// 尝试建立连接
	var clientConn *websocket.Conn
	for tryTimes = 0; tryTimes < maxTryTimes; tryTimes++ {

		clientConn, _, e = dialer.Dial(urlString.String(), nil)
		if e == nil {
			break
		}
		// 报错则休眠100毫秒
		time.Sleep(time.Millisecond * 100)
	}

	if tryTimes >= maxTryTimes {
		// 操过最大尝试次数则报错
		panic(fmt.Sprint("Cannot create connection with ", serverConfig.ID))
	}

	// 如果超过最大尝试次数，任然有错则报错
	if e != nil {
		panic(e)
	}

	// 连接成功！！！
	fmt.Println("连接到", serverConfig.ID, "成功！！！")

	// 接收消息
	go func() {
		for {
			_, data, err := clientConn.ReadMessage()

			if err != nil {
				clientConn.Close()
				clientConn.CloseHandler()(0, "")
				DelClientByID(serverConfig.ID)
				fmt.Println("服务", serverConfig.ID, "掉线")
				break
			}

			responseMessage := &msg.ResponseMessage{}
			err = json.Unmarshal(data, responseMessage)
			if err != nil {
				fmt.Println("Rpc request's response body parse fail")
				continue
			}
			if responseMessage.Index != 0 {
				requestFunc, ok := requestMap[responseMessage.Index]
				if ok {
					delete(requestMap, responseMessage.Index)
					requestFunc(data)
				}
				continue
			}
			fmt.Println("Notify消息:", string(data))
		}
	}()

	return &RPCClient{
		clientConn:   clientConn,
		serverConfig: serverConfig,
	}
}
