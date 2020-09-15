/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/4 11:31 上午
* @Desc:
 */

package sort

/**
合并两个有序数组
 */
func MergeTwoSortedArr(arr1 []int, arr2 []int) []int {
	res := []int{}
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] <= arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}
	res = append(res, arr1...)
	res = append(res, arr2...)
	return res
}
