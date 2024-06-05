package judge

import "regexp"

// IsNumber 判断文本是否为纯数字
func IsNumber(text string) bool {
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

// 校验占位符是否配对
func isBalancedChar(text string) bool {
	stack := make([]rune, 0)

	for _, char := range text {
		switch char {
		case '{', '[', '<', '「':
			stack = append(stack, char)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '>':
			if len(stack) == 0 || stack[len(stack)-1] != '<' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '」':
			if len(stack) == 0 || stack[len(stack)-1] != '「' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
