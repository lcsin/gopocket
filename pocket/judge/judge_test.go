package judge

import (
	"fmt"
	"testing"
)

func TestIsNumber(t *testing.T) {
	texts := []string{
		"123124324",
		"hello，world",
		"2342，hello",
		"-1",
		"0",
	}

	for _, text := range texts {
		fmt.Println(IsNumber(text))
	}
}

func TestContainChinese(t *testing.T) {
	texts := []string{
		"你好，世界",
		"你好，world",
		"hello, world",
	}

	for _, text := range texts {
		fmt.Println(ContainChinese(text))
	}
}

func TestIsBalancedChars(t *testing.T) {
	text := "{【sfsf】}cxfdsaf[23432],(sfsdf)sfsdf"
	fmt.Println(BalancedChars(text, "{}", "[]", "()", "【】"))
	fmt.Println(IsBalancedChars(text))
	text = "{【sfsf】}cxfdsaf[23432],(sfsdf)sfsdf)"
	fmt.Println(BalancedChars(text, "{}", "[]", "()", "【】"))
	fmt.Println(IsBalancedChars(text))
}
