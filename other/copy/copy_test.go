package copy

import (
	"fmt"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	letters := map[string]string{
		"a" : "aaa",
	}
	letters2 := map[string]string{}
	_ = DeepCopy(&letters2, letters)
	letters2["a"] = "cake"
	fmt.Println(letters)
	fmt.Println(letters2)
	if _, ok := letters2["dsafasdfa"]; !ok {
		fmt.Println("no letter")
	}
}