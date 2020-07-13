package singleton

import "testing"

func TestGetIns(t *testing.T)  {
	go GetIns()
	go GetIns()
	go GetIns()
}