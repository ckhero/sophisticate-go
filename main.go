package main

import "main/design/chain"

func main()  {
	v1 := chain.NewTeacher()
	v2 := chain.NewHeadermaster()
	v1.SetNext(v2)
	v := v1
	v.Exec(1)
	v.Exec(3)
	v.Exec(2)
	v.Exec(4)
}