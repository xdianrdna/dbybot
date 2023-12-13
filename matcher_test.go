package dbybot

import "testing"

type matcherCase struct {
	keyword string
	result  bool
}

var textContext = &Context{DisposedMsg: "foo bar"}

// 测试前缀匹配器
func TestPrefixMatcher(t *testing.T) {
	matcher := prefixMatcher{}

	testCases := []matcherCase{
		{"", true},
		{"foo", true},
		{"foo bar", true},

		{"foo bar bar", false},
		{"bar", false},
		{"hello", false},
	}

	for _, testCase := range testCases {
		if testCase.result != matcher.Judge(textContext, testCase.keyword) {
			t.Fail()
		}
	}
}

// 测试后缀匹配器
func TestSuffixMatcher(t *testing.T) {
	matcher := suffixMatcher{}

	testCases := []matcherCase{
		{"", true},
		{"bar", true},
		{"foo bar", true},

		{"fo", false},
		{"foo bar bar", false},
		{"hello", false},
	}

	for _, testCase := range testCases {
		if testCase.result != matcher.Judge(textContext, testCase.keyword) {
			t.Fail()
		}
	}
}

// 测试包含匹配器
func TestContainsMatcher(t *testing.T) {
	matcher := containsMatcher{}

	testCases := []matcherCase{
		{"", true},
		{" ", true},
		{"foo", true},
		{"bar", true},

		{"foobar", false},
		{"dbybot", false},
	}

	for _, testCase := range testCases {
		if testCase.result != matcher.Judge(textContext, testCase.keyword) {
			t.Fail()
		}
	}
}
