/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 4:05 下午
* @Desc: a
 */

package memory_align

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestOptimizeMemory(t *testing.T) {
	//OptimizeMemory()
	fmt.Println(unsafe.Sizeof(Test1{}))
	fmt.Println(unsafe.Sizeof(Test2{}))
}