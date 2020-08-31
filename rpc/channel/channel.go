package channel

import (
	"github.com/YAOHAO9/yh/application"
	"github.com/YAOHAO9/yh/rpc"
	"github.com/YAOHAO9/yh/rpc/message"
	"github.com/YAOHAO9/yh/rpc/session"
	"github.com/YAOHAO9/yh/util/beeku"
)

// Channel ChannelService
type Channel map[string]*session.Session

// PushMessageToUser 推送消息给指定玩家
func (channel Channel) PushMessageToUser(uid string, route string, data interface{}) {
	session, ok := channel[uid]
	if !ok {
		return
	}

	PushMessageBySession(session, route, data)
}

// PushMessage 推送消息给所有人
func (channel Channel) PushMessage(uids []string, route string, data interface{}) {
	for _, uid := range uids {
		channel.PushMessageToUser(uid, route, data)
	}
}

// PushMessageToOthers 推送消息给其他人
func (channel Channel) PushMessageToOthers(uids []string, route string, data interface{}) {
	for _, uid := range uids {
		if beeku.InSlice(uid, uids) == -1 {
			channel.PushMessageToUser(uid, route, data)
		}
	}
}

// Add 推送消息给其他人
func (channel Channel) Add(uid string, session *session.Session) {
	channel[uid] = session
}

// PushMessageBySession 通过session推送消息
func PushMessageBySession(session *session.Session, route string, data interface{}) {
	notify := message.RPCNotify{
		Route: route,
		Data:  data,
	}
	application.RPC.Notify.ToServer(session.CID, session, rpc.SysRPCEnum.PushMessage, notify)
}
