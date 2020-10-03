package abs

/**
一个工厂生产多个产品，产品族的概念。比如这个公长既生产手机又生产电脑
*/
type Phone interface {
	Size() int
}


type PC interface {
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


type XiaomiPC struct {

}

type HuaWeiPC struct {

}
func(c *XiaomiPC) Size() int {
	return 1
}

func(b *HuaWeiPC) Size() int {
	return 2
}


type Factory interface {
	GetPhone() Phone
	GetPC() Phone
}

type XiaoMi struct {

}

type HuaWei struct {

}
func(x *XiaoMi) GetPhone() Phone {
	return &XiaomiPhone{}
}


func(x *HuaWei) GetPhone() Phone {
	return &HuaWeiPhone{}
}


func(x *XiaoMi) GetPC() Phone {
	return &XiaomiPhone{}
}


func(x *HuaWei) GetPC() Phone {
	return &HuaWeiPhone{}
}