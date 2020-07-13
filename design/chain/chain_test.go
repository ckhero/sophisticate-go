package chain

import "testing"

func TestChain(t *testing.T)  {
	v1 := NewTeacher()
	v2 := NewHeadermaster()
	v1.SetNext(v2)
	v := v1
	v.Exec(1)
	v.Exec(3)
	v.Exec(2)
	v.Exec(4)
}
