package ws

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	roomClient  = make(map[interface{}]map[interface{}]*websocket.Conn)  //房间
	roomChannel = make(map[interface{}]map[interface{}]chan interface{}) // 房间消息通道
	roomMux     sync.Mutex                                               // 互斥锁
)

// 处理ws请求
func RoomWsHandler(w http.ResponseWriter, r *http.Request, roomId uint) {
	clientId := uuid.New()
	WsHandler(w, r, clientId, roomId, &roomMux, roomClient, roomChannel)
}
