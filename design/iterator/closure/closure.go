/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 6:48 下午
* @Desc: a
 */

package closure

type Ints []int

func(i Ints) Iterator() func() (int, bool) {
	idx := 0
	return func() (int, bool) {
		if idx >= len(i) {
			return -1, false
		}
		val := i[idx]
		idx++
		return val, true
	}
}
