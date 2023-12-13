package dbybot

// 向大别野发送请求的主要逻辑
import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// 机器人向大别野 api 发送请求
// villaId 为按官方要求加入头中的参数
func (bot *Bot) sendMessagePostRequest(villaId uint64, json string) {
	url := "https://bbs-api.miyoushe.com/vila/api/bot/platform/sendMessage"

	body := strings.NewReader(json)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Build New Request Error:", err)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-rpc-bot_id", bot.auth.id)
	request.Header.Add("x-rpc-bot_secret", bot.auth.secretKey)
	request.Header.Add("x-rpc-bot_villa_id", strconv.Itoa(int(villaId)))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Request Error:", err)
		return
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bstring := string(bodyBytes)
	fmt.Println(bstring)
}
