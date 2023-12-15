# dbybot

简易的大别野机器人的 sdk。

当前只支持回复文本消息。

```go
package main

import "github.com/xdianrdna/dbybot"

func main() {
	// 注册机器人信息，填写在大别野后台取得的机器人 id 与 sercet key
	dbybot.Register("bot_id", "bot_key")

	// 规则引擎
	engine := dbybot.Engine()

	// 侦测以 /nn 开头的消息 (此处的消息是经过处理的)
	// 当消息以 ping 开头时，回复 pong
	engine.StartWith("ping", func(ctx *dbybot.Context) {
		ctx.ReplyMessage(dbybot.NewTextMsg().Text("pong!"))
	})

	// 开始在 12345 端口 /api/callback 接收回调报文
	dbybot.Serve("/api/callback", 12345)
}
```
