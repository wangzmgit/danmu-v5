package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/vo"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理ws请求
func WsHandler(
	w http.ResponseWriter,
	r *http.Request,
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	clientMap map[interface{}]map[interface{}]*websocket.Conn,
	messageChannel map[interface{}]map[interface{}]chan interface{},
) {

	var conn *websocket.Conn
	var err error
	// 创建一个定时器用于服务端心跳
	pingTicker := time.NewTicker(time.Second * 10)
	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		util.LogError("Failed to set websocket upgrade " + err.Error())
		return
	}
	// 把与客户端的链接添加到客户端链接池中
	addClient(id, groupId, mux, conn, clientMap)
	// 获取该客户端的消息通道
	m, exist := getMessageChannel(id, groupId, mux, messageChannel)
	if !exist {
		m = make(chan interface{})
		addMessageChannel(id, groupId, mux, messageChannel, m)
	}

	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteClient(id, groupId, mux, clientMap, messageChannel)
		// fmt.Println(code)
		return nil
	})

	//如果不是默认分组广播房间人数
	if groupId != "0" {
		BroadcastNumber(groupId, mux, clientMap, messageChannel)
	}

	for {
		select {
		case content, ok := <-m:
			// 从消息通道接收消息，然后推送给前端
			err = conn.WriteJSON(content)
			if err != nil {
				util.LogError("Send MEssage Error" + err.Error())
				if ok {
					go func() {
						m <- content
					}()
				}

				conn.Close()
				deleteClient(id, groupId, mux, clientMap, messageChannel)
				return
			}
		case <-pingTicker.C:
			// 服务端心跳:每20秒ping一次客户端，查看其是否在线
			conn.SetWriteDeadline(time.Now().Add(time.Second * 20))
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				util.LogError("Send ping error " + err.Error())
				conn.Close()
				deleteClient(id, groupId, mux, clientMap, messageChannel)
				return
			}
		}
	}

}

//添加客户端
func addClient(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	conn *websocket.Conn,
	clientMap map[interface{}]map[interface{}]*websocket.Conn,
) {
	mux.Lock()
	if clientMap[groupId] == nil {
		clientMap[groupId] = make(map[interface{}]*websocket.Conn)
	}
	clientMap[groupId][id] = conn
	mux.Unlock()
}

//获取消息管道
func getMessageChannel(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	messageChannel map[interface{}]map[interface{}]chan interface{},
) (m chan interface{}, exist bool) {

	mux.Lock()
	m, exist = messageChannel[groupId][id]
	mux.Unlock()
	return
}

//添加消息管道
func addMessageChannel(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	messageChannel map[interface{}]map[interface{}]chan interface{},
	m chan interface{},
) {
	mux.Lock()
	if messageChannel[groupId] == nil {
		messageChannel[groupId] = make(map[interface{}]chan interface{})
	}
	messageChannel[groupId][id] = m
	mux.Unlock()
}

//移除客户端
func deleteClient(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	clientMap map[interface{}]map[interface{}]*websocket.Conn,
	messageChannel map[interface{}]map[interface{}]chan interface{},
) {
	mux.Lock()
	delete(clientMap[groupId], id)
	delete(messageChannel[groupId], id)
	mux.Unlock()
	BroadcastNumber(groupId, mux, clientMap, messageChannel) //广播房间人数
}

//设置消息到房间内所有客户端
func setMessageAllClient(
	groupId interface{},
	mux *sync.Mutex,
	messageChannel map[interface{}]map[interface{}]chan interface{},
	content interface{},
) {
	mux.Lock()
	all := messageChannel[groupId]
	mux.Unlock()
	go func() {
		for _, m := range all {
			m <- content
		}
	}()
}

//设置消息
func setMessage(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	messageChannel map[interface{}]map[interface{}]chan interface{},
	content interface{},
) {
	mux.Lock()
	if m, exist := messageChannel[groupId][id]; exist {
		go func() {
			m <- content
		}()
	}
	mux.Unlock()
}

//广播房间人数
func BroadcastNumber(
	groupId interface{},
	mux *sync.Mutex,
	clientMap map[interface{}]map[interface{}]*websocket.Conn,
	messageChannel map[interface{}]map[interface{}]chan interface{},
) {
	setMessageAllClient(groupId, mux, messageChannel, &vo.RoomVo{
		Number: len(clientMap[groupId]),
	})
	return
}

/*********************************************************
** 函数功能: 向用户发送消息
** 日    期: 2022年2月25日09:37:48
**********************************************************/
func sendMsg(
	id interface{},
	groupId interface{},
	mux *sync.Mutex,
	messageChannel map[interface{}]map[interface{}]chan interface{},
	content interface{},
) {
	if id != 0 {
		_, exist := getMessageChannel(id, groupId, mux, messageChannel)
		if !exist {
			return
		}
	}
	setMessage(id, groupId, mux, messageChannel, content)
	return
}
