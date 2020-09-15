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
	arr2 := []int{3,9,1,11,12,15}
	sort.QuickSort2(arr2, 0, len(arr2) - 1)
	fmt.Println(arr2)
}