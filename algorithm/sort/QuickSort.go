/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/7 3:33 下午
* @Desc: a
 */

package sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := arr[0]
	left, right := []int{}, []int{}
	for k, v := range arr {
		if k == 0 {
			continue
		}
		if v <= mid {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	res := append([]int{},left...)
	res = append(res,mid)
	return append(res,right...)
}