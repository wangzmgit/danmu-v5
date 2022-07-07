package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"kuukaa.fun/danmu-v5/util"
	"kuukaa.fun/danmu-v5/vo"
)

var (
	room        = make(map[string]map[uuid.UUID]*websocket.Conn)  //房间
	roomMessage = make(map[string]map[uuid.UUID]chan interface{}) // 房间消息通道
	roomMux     sync.Mutex                                        // 互斥锁
)

// 处理ws请求
func WsHandler(w http.ResponseWriter, r *http.Request, roomId string) {
	var conn *websocket.Conn
	var err error
	clientId := uuid.New()
	// 创建一个定时器用于服务端心跳
	pingTicker := time.NewTicker(time.Second * 10)
	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		util.LogError("room failed to set websocket upgrade " + err.Error())
		return
	}
	// 把与客户端的链接添加到客户端链接池中
	addClientToRoom(roomId, clientId, conn)
	// 获取该客户端的消息通道
	m, exist := getRoomMessageChannel(roomId, clientId)
	if !exist {
		m = make(chan interface{})
		addRoomMessageChannel(roomId, clientId, m)
	}

	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteRoomClient(roomId, clientId)
		return nil
	})
	//广播房间人数
	BroadcastNumber(roomId)
	for {
		select {
		case content, ok := <-m:
			// 从消息通道接收消息，然后推送给前端
			err = conn.WriteJSON(content)
			if err != nil {
				util.LogError("room send Message Error " + err.Error())
				if ok {
					go func() {
						m <- content
					}()
				}

				conn.Close()
				deleteRoomClient(roomId, clientId)
				return
			}
		case <-pingTicker.C:
			// 服务端心跳:每20秒ping一次客户端，查看其是否在线
			conn.SetWriteDeadline(time.Now().Add(time.Second * 20))
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				util.LogError("room send ping error " + err.Error())
				conn.Close()
				deleteRoomClient(roomId, clientId)
				return
			}
		}
	}

}

//广播房间人数
func BroadcastNumber(roomId string) {
	setMessageAllClient(roomId, &vo.RoomVo{
		Number: len(room[roomId]),
	})
	return
}

//添加客户端到房间
func addClientToRoom(roomId string, clientId uuid.UUID, conn *websocket.Conn) {
	roomMux.Lock()
	if room[roomId] == nil {
		room[roomId] = make(map[uuid.UUID]*websocket.Conn)
	}
	room[roomId][clientId] = conn
	roomMux.Unlock()
}

//移除客户端
func deleteRoomClient(roomId string, clientId uuid.UUID) {
	roomMux.Lock()
	delete(room[roomId], clientId)
	delete(roomMessage[roomId], clientId)
	roomMux.Unlock()
	BroadcastNumber(roomId) //广播房间人数
}

//添加消息管道
func addRoomMessageChannel(roomId string, clientId uuid.UUID, m chan interface{}) {
	roomMux.Lock()
	if roomMessage[roomId] == nil {
		roomMessage[roomId] = make(map[uuid.UUID]chan interface{})
	}
	roomMessage[roomId][clientId] = m
	roomMux.Unlock()
}

// 获取消息管道
func getRoomMessageChannel(roomId string, clientId uuid.UUID) (m chan interface{}, exist bool) {
	roomMux.Lock()
	m, exist = roomMessage[roomId][clientId]
	roomMux.Unlock()
	return
}

//设置消息到房间内所有客户端
func setMessageAllClient(roomId string, content interface{}) {
	roomMux.Lock()
	all := roomMessage[roomId]
	roomMux.Unlock()
	go func() {
		for _, m := range all {
			m <- content
		}
	}()
}
