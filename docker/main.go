package main

import (
	"docker/route"
	"net/http"
)

type aa interface {
}
func main()  {
	route.Route()
	http.ListenAndServe("0.0.0.0:1001", nil)
}


