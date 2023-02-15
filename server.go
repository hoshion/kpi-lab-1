package main

import "net/http"

func main() {
	http.HandleFunc("/time", HandleTime)

	Start(8795)
}
