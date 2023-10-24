package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
  log.Println("Hello Im starting now I think....")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


