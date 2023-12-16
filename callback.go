package dbybot

type Callback struct {
	Event Event `json:"event"`
}

// Event 事件信息
type Event struct {
	Robot      Robot      `json:"robot"`
	Type       int        `json:"type"`        // 事件类型 todo: use const
	ExtendData ExtendData `json:"extend_data"` // 包含事件的具体数据
	CreatedAt  uint64     `json:"created_at"`
	SendAt     uint64     `json:"send_at"`
}

// Robot 机器人相关信息
type Robot struct {
	Template Template `json:"template"`
	VillaId  uint64   `json:"villa_id"`
}

// Template 机器人模板信息
type Template struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Icon     string    `json:"icon"`
	Commands []Command `json:"commands"`
}

// Command 机器人装载的命令
type Command struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// ExtendData 事件的具体数据
type ExtendData struct {
	EventData struct {
		SendMessage SendMessage `json:"SendMessage"`
	} `json:"EventData"`
}

// SendMessage 用户@机器人发送消息的回调数据
type SendMessage struct {
	ContentStr string `json:"content"`
	OContent   OContent
	FromUserID int      `json:"from_user_id"`
	SendAt     int      `json:"send_at"`
	RoomID     int      `json:"room_id"`
	ObjectName int      `json:"object_name"`
	Nickname   string   `json:"nickname"`
	MsgUID     string   `json:"msg_uid"`
	BotMsgID   string   `json:"bot_msg_id"`
	VillaID    int      `json:"villa_id"`
	QuoteMsg   QuoteMsg `json:"quote_msg"`
}

// QuoteMsg 回调消息引用消息的基础信息
type QuoteMsg struct {
	ContentStr       string `json:"content"`
	OContent         OContent
	MsgUID           string `json:"msg_uid"`
	BotMsgID         string `json:"bot_msg_id"`
	SendAt           int    `json:"send_at"`
	MsgType          string `json:"msg_type"`
	FromUserID       int    `json:"from_user_id"`
	FromUserNickname string `json:"from_user_nickname"`
	FromUserIDStr    string `json:"from_user_id_str"`
}

type OContent struct {
	Content Content `json:"content"`
}
