package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello, meals and docker")
	http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("web/dist")))
	fmt.Println("goodbye")
}
