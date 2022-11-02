package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"red-server/global"
	"red-server/utils"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type MessageCmd int

const (
	MESSAGE_CMD_HEART  = 0
	MESSAGE_CMD_SINGLE = 1
	MESSAGE_CMD_ROOM   = 2
)

type Media int

const (
	MEDIA_TYPE_TEXT  Media = iota // 0 文本
	MEDIA_TYPE_NEWS               // 1 新闻
	MEDIA_TYPE_VOICE              // 2 语音
	MEDIA_TYPE_IMG                // 3 图片
	MEDIA_TYPE_VIDEO              // 4 视频
	MEDIA_TYPE_MUSIC              // 5 音乐
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
}

type Message struct {
	Id       int64      `json:"id,omitempty" form:"id"`             // 消息ID
	UserId   int64      `json:"userId,omitempty" form:"userId"`     // 谁发的
	Cmd      MessageCmd `json:"cmd,omitempty" form:"cmd"`           // 群聊还是私聊
	TargetId int64      `json:"targetId,omitempty" form:"targetId"` // 对端用户ID/群ID
	Media    int        `json:"media,omitempty" form:"media"`       // 消息按照什么样式展示
	Content  string     `json:"content,omitempty" form:"content"`   // 消息的内容
	Pic      string     `json:"pic,omitempty" form:"pic"`           // 预览图片
	Url      string     `json:"url,omitempty" form:"url"`           // 服务的URL
	Memo     string     `json:"memo,omitempty" form:"memo"`         // 简单描述
	Amount   int        `json:"amount,omitempty" form:"amount"`     // 其他和数字相关的
}

type ChatController struct {
	rwLocker  sync.RWMutex
	clientMap map[int64]*Node
}

func NewChatController() *ChatController {
	return &ChatController{
		clientMap: make(map[int64]*Node, 0),
	}
}

func (c *ChatController) Chat(ctx *gin.Context) {
	currentUser := utils.GetCurrentUser(ctx)
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		global.Logger.Error(err)
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
	}
	c.rwLocker.Lock()
	c.clientMap[currentUser.Id] = node
	c.rwLocker.Unlock()
	go c.sendProc(node)
	go c.recvProc(node)
}

func (c *ChatController) sendProc(node *Node) {
	for data := range node.DataQueue {
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			global.Logger.Error(err)
			return
		}
	}
}

func (c *ChatController) recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			global.Logger.Error(err)
			return
		}
		c.dispatch(data)
		fmt.Printf("接收消息: %s", data)
	}
}

func (c *ChatController) dispatch(data []byte) {
	msg := Message{}

	err := json.Unmarshal(data, &msg)
	if err != nil {
		global.Logger.Error(err)
		return
	}

	switch msg.Cmd {
	case MESSAGE_CMD_SINGLE:
		c.sendMsg(msg.TargetId, data)
	case MESSAGE_CMD_ROOM:
		// 群聊消息
	case MESSAGE_CMD_HEART:
		// 心跳
	}
}

func (c *ChatController) sendMsg(userId int64, data []byte) {
	c.rwLocker.RLock()
	node, ok := c.clientMap[userId]
	c.rwLocker.RUnlock()
	if ok {
		node.DataQueue <- data
	}
}

func (c *ChatController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/ws", c.Chat)
}
