package main

import (
	"net/http"
	"strconv"
)

func Start(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
