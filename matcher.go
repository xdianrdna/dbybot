package dbybot

import "strings"

type HandleFunc func(*Context)

// matcher 是指令匹配器，用于消息是否命中关键词
type Matcher interface {
	Judge(*Context, string) bool
}

// 前缀匹配器
type prefixMatcher struct{}

func (matcher prefixMatcher) Judge(ctx *Context, key string) bool {
	return strings.HasPrefix(ctx.DisposedMsg, key)
}

// 后缀匹配器
type suffixMatcher struct{}

func (matcher suffixMatcher) Judge(ctx *Context, key string) bool {
	return strings.HasSuffix(ctx.DisposedMsg, key)
}

// 包含匹配器
type containsMatcher struct{}

func (matcher containsMatcher) Judge(ctx *Context, key string) bool {
	return strings.Contains(ctx.DisposedMsg, key)
}
