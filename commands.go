package dbybot

type command struct {
	keyword string     // 触发指令的关键词
	matcher Matcher    // 匹配器
	fn      HandleFunc // 执行的函数
}

type commandEngine struct {
	commands []command
}

// 内部使用的追加命令功能
func (e *commandEngine) addCommand(matcher Matcher, keyword string, fn HandleFunc) {
	command := command{
		keyword: keyword,
		matcher: matcher,
		fn:      fn,
	}

	e.commands = append(e.commands, command)
}

// StartWith 以前缀匹配新增一条命令
func (e *commandEngine) StartWith(keyword string, fn HandleFunc) {
	e.addCommand(prefixMatcher{}, keyword, fn)
}

// EndWith 以后缀匹配新增一条命令
func (e *commandEngine) EndWith(keyword string, fn HandleFunc) {
	e.addCommand(suffixMatcher{}, keyword, fn)
}

// Contains 以包含匹配新增一条命令
func (e *commandEngine) Contains(keyword string, fn HandleFunc) {
	e.addCommand(containsMatcher{}, keyword, fn)
}

// UseMatcher 自定义
func (e *commandEngine) UseMatcher(matcher Matcher, keyword string, fn HandleFunc) {
	e.addCommand(matcher, keyword, fn)
}
