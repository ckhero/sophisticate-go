package route

import (
	"docker/handle"
	"net/http"
)
func Route()  {

	http.HandleFunc("/", handle.IndexHandler)
}