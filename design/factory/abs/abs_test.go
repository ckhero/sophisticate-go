package abs

import (
	"fmt"
	"testing"
)

func TestHuaWei_GetPC(t *testing.T) {

	xiaomiF := &XiaoMi{}
	fmt.Println(xiaomiF.GetPC().Size())
	fmt.Println(xiaomiF.GetPhone().Size())
}
