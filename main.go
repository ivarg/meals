package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting web server...")
	http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("web/dist")))
}
