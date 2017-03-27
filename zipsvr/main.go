package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//varName type 		* means it is sending a reference and not a copy
// 					we are passing a pointer to the request
// 					in go, any type can be passed by value and any type can be passed by reference
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// must add headers before you start writing the body
	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("hello " + name))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		log.Fatal("please set addr env var")
	}

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

/*
compile with "go install" and run with "zipsvr"
$ go install && zipsvr
*/
