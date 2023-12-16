package model

// 用户 api 为开发者提供有关用户操作的接口，包括获取用户信息、获取大别野成员列表和踢出大别野用户。

type color string

const (
	colorRed color = "#F47884"
	//...
)

// Member 用户信息
type Member struct {
	MemberBasic membetBasic `json:"basic"`

	RoleIdList []string   `json:"role_id_list"` // 用户加入的身份组 id 列表
	JoinedAt   string     `json:"joined_at"`    // 用户加入时间  NOTE: 返回的是 string 类型
	RoleList   memberRole `json:"roleList"`     // 用户已加入的身份组列表
}

// 用户基本信息
type membetBasic struct {
	Uid       string `json:"uid"`        // uid
	Nickname  string `json:"nickname"`   // 昵称
	Introduct string `json:"introduce"`  // 个性签名
	AvatarUrl string `json:"avatar_url"` // 头像地址
}

// 身份组信息
type memberRole struct {
	Id        uint64   `json:"id"`          // 身份组 id
	Name      string   `json:"name"`        // 身份组名称
	VillaId   string   `json:"villa_id"`    // 大别野 id
	Color     color    `json:"color"`       // 身份组颜色  todo: 参考颜色
	RoleType  string   `json:"role_type"`   //	身份组类型
	IsAllRoom bool     `json:"is_all_room"` //	是否选择全部房间
	RoomIDs   []uint64 `json:"room_ids"`    //	指定的房间列表
}
