/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 6:44 下午
* @Desc: a
 */

package normal

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T)  {
	ints := Ints{1, 2, 3}
	for it := ints.First(); it.HasNext(); it = it.Next() {
		fmt.Println(it.Val())
	}
}