/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/14 4:50 下午
* @Desc: 原型模式(Prototype)[url:http://c.biancheng.net/view/1343.html]
*/
package prototype

type Prototype struct {
	Name string
	Age int
}


func (p *Prototype) Clone() *Prototype {
	res := *p
	return &res
}