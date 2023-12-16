package dbybot

import (
	"regexp"
	"strings"
)

// Context 是包含了一系列机器执行指令时上下文信息的结构体
type Context struct {
	bot *Bot

	NickName      string
	OriginMessage string // 原始消息
	DisposedMsg   string // 经处理的消息

	VillaId uint64
	RoomId  uint64

	SenderNickName string // 消息发送者昵称
	SenderId       int    // 消息发送者 uid

	callback *Callback // 一些特别的场景当中，可以主动获取 callback 进行进一步的处理
}

func NewBotContext(callback *Callback) *Context {
	ctx := &Context{}
	ctx.bot = globalBot

	ctx.NickName = callback.Event.Robot.Template.Name

	ctx.OriginMessage = callback.Event.ExtendData.EventData.
		SendMessage.OContent.Content.Text
	ctx.DisposedMsg = disposeMessage(callback) // 处理后的消息

	ctx.VillaId = callback.Event.Robot.VillaId
	ctx.RoomId = uint64(callback.Event.ExtendData.EventData.SendMessage.RoomID)

	ctx.SenderNickName = callback.Event.ExtendData.EventData.SendMessage.Nickname
	ctx.SenderId = callback.Event.ExtendData.EventData.SendMessage.FromUserID

	return ctx
}

// 处理原始消息，给出更便于处理的文本消息形式。
//
// 该函数会去除消息中所有 "@机器人" 的内容，然后移除前导空格与后导空格，并且只保留最多一个连续空格。
// 如原始消息为(机器人名为 foo) "   @foo hello    world  @foo  ", 处理后为 "hello world"
func disposeMessage(callback *Callback) string {
	message := callback.Event.ExtendData.EventData.
		SendMessage.OContent.Content.Text
	nickname := callback.Event.Robot.Template.Name

	removeStr := "@" + nickname + " "
	message = strings.ReplaceAll(message, removeStr, "")
	message = strings.Trim(message, " ")
	reg := regexp.MustCompile(`( )+|(\n)+`)
	message = reg.ReplaceAllString(message, "$1$2")
	return message
}

// FetchCallback 提供给使用方获取 callback 的方法
func (ctx *Context) FetchCallback() *Callback {
	return ctx.callback
}
