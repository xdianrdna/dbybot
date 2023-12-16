package dbybot

// 向大别野发送请求的主要逻辑
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/xdianrdna/dbybot/model"
)

const baseApi = "https://bbs-api.miyoushe.com"

// 基础响应
type baseResp struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
}

// 统一响应体
type mhyResp struct {
	baseResp
	Data struct {
		Member model.Member `json:"member"`
	} `json:"data"`
}

// MemberResp 用户信息响应
type MemberResp struct {
	baseResp
	Member model.Member
}

// 处理公共的加头与请求部分
func (bot *Bot) commonHanlder(req *http.Request, villaId uint64) ([]byte, bool) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-rpc-bot_id", bot.auth.id)
	req.Header.Add("x-rpc-bot_secret", bot.auth.secretKey)
	req.Header.Add("x-rpc-bot_villa_id", strconv.Itoa(int(villaId)))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error:", err)
		return nil, false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// todo: xq你看一下
		}
	}(response.Body)

	bodyBytes, _ := io.ReadAll(response.Body)

	return bodyBytes, true
}

// 机器人向大别野 api 发送请求
// villaId 为按官方要求加入头中的参数
func (bot *Bot) sendMessagePostRequest(villaId uint64, json string) {
	url := baseApi + "/vila/api/bot/platform/sendMessage"

	request, err := http.NewRequest("POST", url, strings.NewReader(json))
	if err != nil {
		fmt.Println("Build New Request Error:", err)
		return
	}

	bot.commonHanlder(request, villaId)
}

// 调用官方 getMember 接口，获取用户信息
func (bot *Bot) getMember(villaId, uid uint64) *MemberResp {
	url := baseApi + "/vila/api/bot/platform/getMember?uid=" + strconv.Itoa(int(uid))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Build New Request Error:", err)
		return nil
	}

	b, ok := bot.commonHanlder(request, villaId)
	if !ok {
		fmt.Println("Build New Request Error")
		return nil
	}

	r := &mhyResp{}
	err = json.Unmarshal(b, r)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &MemberResp{
		baseResp: r.baseResp,
		Member:   r.Data.Member,
	}
}
