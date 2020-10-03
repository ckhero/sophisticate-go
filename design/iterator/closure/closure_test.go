/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 6:54 下午
* @Desc: a
 */

package closure

import (
	"fmt"
	"testing"
)

func TestInts_Iterator(t *testing.T) {
	ints := Ints{1, 2, 3}
	it := ints.Next()
	for {
		if val, ok := it(); ok {
			fmt.Println(val)
		} else {
			fmt.Println("end")

			break
		}
	}
}