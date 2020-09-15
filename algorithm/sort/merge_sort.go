/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/7 2:57 下午
* @Desc: a
 */

package sort

func MergeSort(arr []int) []int{
	if len(arr) < 2 {
		return arr
	}
	mid := (0 + len(arr)) >> 1
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return MergeTwoSortedArr(left, right)
}
