package data_struct

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Printf("cap:%d \n", cap(ch))
	fmt.Printf("size:%d \n", len(ch))

}
