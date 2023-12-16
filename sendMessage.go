package dbybot

import (
	"encoding/json"
	"log"
)

// 包含了发送消息的基本请求实例

// 发送实例结构声明
type SendMessageReq struct {
	RoomId        uint64 `json:"room_id"`
	ObjectName    string `json:"object_name"`
	MsgContentStr string `json:"msg_content"`
}

type MsgContent struct {
	Content Content `json:"content"`
}

type Content struct {
	Text     string         `json:"text"`
	Entities []EntitiesItem `json:"entities"`
}

type EntitiesItem struct {
	Entity Entity `json:"entity"`
	Length int    `json:"length"`
	Offset int    `json:"offset"`
}

type Entity struct {
	Type   string `json:"type"`
	UserId string `json:"user_id"`
}

// TextMsg 向服务器发送的消息
type TextMsg struct {
	req        SendMessageReq // 向服务器发送的请求实体
	msgContent MsgContent
}

// NewTextMsg 新建一条消息
func NewTextMsg() *TextMsg {
	return &TextMsg{
		req: SendMessageReq{
			ObjectName: "MHY:Text",
		},
		msgContent: MsgContent{
			Content: Content{
				Entities: []EntitiesItem{},
			},
		},
	}
}

// Text 消息追加文本
func (tb *TextMsg) Text(text string) *TextMsg {
	tb.msgContent.Content.Text += text
	return tb
}

// AtAll @全体成员
func (tb *TextMsg) AtAll() *TextMsg {
	appendText := "@全体成员 "
	length := lengthAtUtf16(appendText)

	tb.msgContent.Content.Entities = append(tb.msgContent.Content.Entities, EntitiesItem{
		Entity: Entity{
			Type: "mentioned_all",
		},
		Length: length,
		Offset: lengthAtUtf16(tb.msgContent.Content.Text),
	})
	tb.msgContent.Content.Text += appendText

	return tb
}

// AtMember @单个成员
// todo: 这个地方应该要通过 userId 查到 nickname ...
func (tb *TextMsg) AtMember(nickname, userId string) *TextMsg {

	appendText := "@" + nickname + " "
	length := lengthAtUtf16(appendText)
	tb.msgContent.Content.Entities = append(tb.msgContent.Content.Entities, EntitiesItem{
		Entity: Entity{
			Type:   "mentioned_user",
			UserId: userId,
		},
		Length: length,
		Offset: lengthAtUtf16(tb.msgContent.Content.Text),
	})

	tb.msgContent.Content.Text += appendText

	return tb
}

// 机器人发送消息到指定房间
// todo: 错误处理
func (bot *Bot) sendMessageTo(tb *TextMsg, villaId, roomId uint64) {
	tb.req.RoomId = roomId

	b, err := json.Marshal(tb.msgContent)
	if err != nil {
		log.Fatal()
	}

	// 将 msg_content 对应的结构序列化后组装到发信请求中
	tb.req.MsgContentStr = string(b)

	b, err = json.Marshal(tb.req)
	if err != nil {
		log.Fatal()
	}

	bot.sendMessagePostRequest(villaId, string(b))
}

// SendMessageTo 向指定别野或房间发送消息由机器人本体触发
func (bot *Bot) SendMessageTo(tb *TextMsg, villaId, roomId uint64) {
	bot.sendMessageTo(tb, villaId, roomId)
}

// ReplyMessage 回复消息由上下文直接触发
func (ctx *Context) ReplyMessage(tb *TextMsg) {
	ctx.bot.sendMessageTo(tb, ctx.bot.Context.VillaId, ctx.bot.Context.RoomId)
}

// 获取 utf 字符长度
func lengthAtUtf16(s string) int {
	return len([]rune(s))
}
