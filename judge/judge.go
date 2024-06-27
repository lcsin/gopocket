package judge

import (
	"regexp"
)

// Number 判断文本是否为纯数字
func Number(text string) bool {
	if len(text) == 0 {
		return false
	}
	matched, _ := regexp.MatchString("^[0-9]*$", text)
	return matched
}

// ContainChinese 判断文本是否包含中文
func ContainChinese(text string) bool {
	if len(text) == 0 {
		return false
	}
	matched, _ := regexp.MatchString("\\p{Han}", text)
	return matched
}

// BalancedChars 校验文本中的字符序列是否左右匹配
func BalancedChars(text string, chars ...string) bool {
	if len(text) == 0 {
		return false
	}

	stack := make([]rune, 0)
	lchar := make(map[rune]bool, len(chars)/2)
	rchar := make(map[rune]bool, len(chars)/2)
	rlchar := make(map[rune]rune, len(chars)/2)

	for _, char := range chars {
		if len(char)%2 != 0 {
			return false
		}

		runes := []rune(char)
		lchar[runes[0]] = true
		rchar[runes[1]] = true
		rlchar[runes[1]] = runes[0]
	}

	for _, char := range text {
		switch true {
		case lchar[char]: // 左边，直接入栈
			stack = append(stack, char)
		case rchar[char]: // 右边，匹配出栈，否则返回false
			if len(stack) == 0 || stack[len(stack)-1] != rlchar[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
