package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	func() error {
		server := &http.Server{Addr: "8080", Handler: http.Handler(nil)}
		return server.ListenAndServe()
	}()

	
}