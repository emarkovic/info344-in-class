package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// option 1 from the sides
func logReq(r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

// option 2 from slides
// takes in a handler from handlers.go and returns one
func logReqs(hfn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		start := time.Now()
		hfn(w, r)
		fmt.Printf("%s\n", time.Since(start))
	}
}

// option 3 - the best way
// this wraps the entire mux
// http.Handler is an interface
// func logRequests(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// preprocessing
// 		log.Printf("%s %s", r.Method, r.URL.Path)
// 		start := time.Now()

// 		// mux
// 		handler.ServeHTTP(w, r)

// 		// post processing
// 		fmt.Printf("%s\n", time.Since(start))
// 	})
// }

func logRequests(logger *log.Logger) Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// preprocessing
			logger.Printf("%s %s", r.Method, r.URL.Path)
			start := time.Now()

			// mux
			handler.ServeHTTP(w, r)

			// post processing
			logger.Printf("%s\n", time.Since(start))
		})
	}
}
