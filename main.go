package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HeaderLoggerServer)
	http.ListenAndServe(":8080", nil)
}

func HeaderLoggerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK %s!", r.URL.Path[1:])

	// Loop over header names
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
}
