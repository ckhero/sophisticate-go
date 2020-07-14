package prototype

type Prototype struct {
	Name string
	Age int
}

func (p *Prototype) Clone() *Prototype {
	res := *p
	return &res
}