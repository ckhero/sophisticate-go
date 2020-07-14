package prototype

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	p1 := new(Prototype)
	p1.Name = "ck"
	p1.Age = 10
	p2 := p1.Clone()
	p2.Age = 11
	p2.Name = "ckk"
	fmt.Println(p2)
	fmt.Println(p1)
}