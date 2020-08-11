package data_struct

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	fmt.Printf("NumCpu: %d \n", runtime.NumCPU())
	fmt.Printf("PreProcs: %d \n", 	runtime.GOMAXPROCS(2))
	fmt.Printf("PreProcs: %d \n", 	runtime.GOMAXPROCS(3))
	fmt.Printf("PreProcs: %d \n", 	runtime.GOMAXPROCS(4))
}