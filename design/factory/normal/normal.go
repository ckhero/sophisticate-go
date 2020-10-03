package normal

/**
一个共产只生产一种产品
 */
type Product interface {
	Size() int
}

type XiaomiPhone struct {

}

type HuaWeiPhone struct {

}

func(c *XiaomiPhone) Size() int {
	return 1
}

func(b *HuaWeiPhone) Size() int {
	return 2
}

type Factory interface {
	GetProduct() Product
}

type XiaoMi struct {

}

type HuaWei struct {

}
func(x *XiaoMi) GetProduct() Product {
	return &XiaomiPhone{}
}


func(x *HuaWei) GetProduct() Product {
	return &HuaWeiPhone{}
}
