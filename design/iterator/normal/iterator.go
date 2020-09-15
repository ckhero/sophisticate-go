/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 6:32 下午
* @Desc: a
 */

package normal

type Ints []int

type Iterator struct {
	data Ints
	idx int
}
func(i Ints) First() *Iterator {
	return &Iterator{
		data: i,
		idx:  0,
	}
}

func(i *Iterator) Val() int {
	return i.data[i.idx]
}

func(i *Iterator) HasNext() bool {
	return i.idx < len(i.data)
}

func(i *Iterator) Next() *Iterator {
	return &Iterator{
		data: i.data,
		idx:  i.idx + 1,
	}
}

