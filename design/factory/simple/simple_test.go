package main

import (
	"fmt"
	"testing"
)

func TestNewFactory(t *testing.T) {
	f := NewFactory()
	p1 := f.MakeProduct(1)
	p2 := f.MakeProduct(2)
	fmt.Println(p1.Size())
	fmt.Println(p2.Size())
}