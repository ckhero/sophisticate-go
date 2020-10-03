package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	multiJson()
}

func multiJson()  {
	m := map[string]string{}
	m["key"] = "value"
	data, _ := json.Marshal(m)
	fmt.Println(data)
	data2, _ := json.Marshal(m)
	fmt.Println(data2)
}


