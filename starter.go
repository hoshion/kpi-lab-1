package main

import (
	"net/http"
	"strconv"
)

func Start(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	HandleError(err)
}
