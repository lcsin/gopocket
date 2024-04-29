package judge

import "regexp"

// IsNumber 判断文本是否为纯数字
func IsNumber(text string) bool {
	matched, _ := regexp.MatchString("^[0-9]*$", text)
	return matched
}

// ContainChinese 判断文本是否包含中文
func ContainChinese(text string) bool {
	matched, _ := regexp.MatchString("\\p{Han}", text)
	return matched
}
