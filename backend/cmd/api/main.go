package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Programa iniciado")
	var response string = "Hello you requested: "

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request iniciado")
		response += "\n" + r.URL.Path
		fmt.Fprintf(w, "%s", response)
	})

	http.ListenAndServe(":80", nil)
}
