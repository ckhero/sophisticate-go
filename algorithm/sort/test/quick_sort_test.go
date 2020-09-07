/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/7 3:42 下午
* @Desc: a
 */

package test

import (
	"fmt"
	"main/algorithm/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3,9,1}
	fmt.Println(sort.QuickSort(arr))
}