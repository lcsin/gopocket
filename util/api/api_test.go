package api

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	req := make(map[string]string)
	req["d"] = "v1"
	req["a"] = "v2"
	req["c"] = ""

	key := "sign key"
	fmt.Println(Sign(req, key))
}
