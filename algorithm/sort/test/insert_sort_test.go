/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/10 4:17 下午
* @Desc: a
 */

package test

import (
	"fmt"
	"main/algorithm/sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr2 := []int{3,9,1,11,12,15}
	sort.InsertSort(arr2)
	fmt.Println(arr2)
}