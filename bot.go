package dbybot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bot struct {
	auth    auth          // 认证信息
	engine  commandEngine // 指令管理器
	Context *Context      // 上下文，将会携带每次生成 callback 时上下文信息
}

type auth struct {
	id        string // bot_id
	secretKey string // bot_serceyKey
}

var globalBot *Bot

func init() {
	globalBot = &Bot{
		auth: auth{},
		engine: commandEngine{
			commands: make([]command, 0),
		},
		Context: &Context{},
	}
}

// 注册别野机器人
func Register(id, secretKey string) {
	globalBot.auth.id = id
	globalBot.auth.secretKey = secretKey
}

func Get() *Bot {
	return globalBot
}

func Engine() *commandEngine {
	return &globalBot.engine
}

// 机器人开始上监听端口，接收回调的地址由 path 指定
// 如 bot.Serve("/mhy", 12345) 将启动在 12345 端口上监听回调请求
func Serve(path string, port uint16) {
	g := gin.Default()

	g.POST(path, func(c *gin.Context) {
		// 解析 callback
		callback := &Callback{}
		c.ShouldBindJSON(&callback)
		err := json.Unmarshal([]byte(callback.Event.ExtendData.EventData.SendMessage.ContentStr),
			&callback.Event.ExtendData.EventData.SendMessage.OContent)
		if err != nil {
			fmt.Println("Callback convert to json error:", err)
		}

		// 每次接收 callback，会重新刷 Context 信息
		// 处理所有 botctx 的信息
		globalBot.Context = NewBotContext(callback)

		// 交由总控进行下一步处理
		globalBot.handleCallback()

		// TODO 这里先固定返回成功
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"retcode": 0,
		})
	})

	portStr := fmt.Sprintf(":%d", port)
	g.Run(portStr)
}

// bot 开始处理一条回调
// 遍历所有命令，匹配时执行对应命令并退出
// 如果找到至少一条匹配项，则执行返回 true，否则返回 false
func (bot *Bot) handleCallback() bool {
	ctx := bot.Context
	for _, command := range bot.engine.commands {
		if command.matcher.Judge(ctx, command.keyword) {
			command.fn(ctx)
			return true
		}
	}

	return false
}
