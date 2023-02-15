package main

import "log"

func HandleError(err error) {
	log.Fatal("ListenAndServe: ", err)
}
