/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/10 3:41 下午
* @Desc: a
 */

package sort

func InsertSort(arr []int) {
	i := 0
	for i < len(arr) {
		index := i
		tmp := arr[i]
		end := i - 1
		for end >= 0 {
			if arr[end] > tmp {
				arr[end + 1] = arr[end]
				index = end
			} else {
				break
			}
			end--
		}
		if index != i {
			arr[index] = tmp
		}
		i++
	}
}