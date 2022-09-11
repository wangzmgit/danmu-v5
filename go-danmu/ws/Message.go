package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	messageClient  = make(map[interface{}]map[interface{}]*websocket.Conn)  // 消息通道
	messageChannel = make(map[interface{}]map[interface{}]chan interface{}) // websocket客户端链接池
	messageMux     sync.Mutex                                               // 互斥锁
)

func MsgWsHandler(w http.ResponseWriter, r *http.Request, id uint) {
	//groupId = "0" 表示不分组
	WsHandler(w, r, id, "0", &messageMux, messageClient, messageChannel)
}

/*********************************************************
** 函数功能: 向用户发送消息
** 日    期: 2022年2月25日09:37:48
**********************************************************/
func SendMsgToUser(id uint, msg interface{}) {
	sendMsg(id, "0", &messageMux, messageChannel, msg)
}
