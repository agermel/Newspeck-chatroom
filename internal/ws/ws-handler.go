package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

// 保存所有连接
var clients = make(map[*websocket.Conn]string)
var clientsLock sync.Mutex

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		httpx.Error(w, err)
		return
	}

	clientsLock.Lock()
	clients[conn] = "active"
	clientsLock.Unlock()

	// 处理客户端消息或心跳
	go func(c *websocket.Conn) {
		defer func() {
			clientsLock.Lock()
			delete(clients, c)
			clientsLock.Unlock()
			c.Close()
		}()
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
		}
	}(conn)
}

// 审查完成后广播推送
func BroadcastFilteredMessage(msg []byte) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	for c := range clients {
		err := c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			c.Close()
			delete(clients, c)
		}
	}
}
