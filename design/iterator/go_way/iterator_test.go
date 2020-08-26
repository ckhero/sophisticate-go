/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 7:01 下午
* @Desc: a
 */

package go_way

import (
	"fmt"
	"testing"
)

func TestInts_Iterator(t *testing.T) {
	ints := Ints{1, 2, 3}
	for v := range ints.Iterator() {
		fmt.Println(v)
	}
}