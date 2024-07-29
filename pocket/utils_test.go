package pocket

import (
	"fmt"
	"testing"
)

func TestSizeStr2Bytes(t *testing.T) {
	b1, err := SizeStr2Bytes("25.2MB")
	if err != nil {
		panic(err)
	}
	fmt.Println(b1, b1/1024/1024)
}

func TestPaginate(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := Paginate(numbers, 1, 3)
	fmt.Println(list)
}
