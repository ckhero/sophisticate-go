/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/4 11:41 上午
* @Desc: a
 */

package test

import (
	"fmt"
	"main/algorithm/sort"
	"testing"
)

func TestMergeTwoSortedArr(t *testing.T) {
	arr1 := []int{1,4,5}
	arr2 := []int{2,3}
	res := sort.MergeTwoSortedArr(arr1, arr2)
	fmt.Print(res)
}