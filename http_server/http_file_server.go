package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http file server is running...")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8090", nil)
}
