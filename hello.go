package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "from golang apps demo"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}
func main() {
	log.Print("Helloworld: starting server ...")

	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
