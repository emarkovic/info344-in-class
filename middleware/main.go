package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "localhost:4000"

	// http.HandleFunc("/v1/hello1", HelloHandler1)
	// option2 from slides, wrapped in logReqs. Adding a function that wraps around to mux
	// http.HandleFunc("/v1/hello1", logReqs(HelloHandler1))
	// http.HandleFunc("/v1/hello2", HelloHandler2)
	// http.HandleFunc("/v1/hello3", HelloHandler3)

	// only logging a few calls and not all of them
	mux := http.NewServeMux()
	muxLogged := http.NewServeMux()
	//adding handlers to our mux
	muxLogged.HandleFunc("/v1/hello1", logReqs(HelloHandler1))
	muxLogged.HandleFunc("/v1/hello2", HelloHandler2)

	mux.HandleFunc("/v1/hello3", HelloHandler3)
	// adding one mux to another mux
	// mux.Handle("/v1/", logRequests(muxLogged))

	logger := log.New(os.Stdout, "", log.LstdFlags)
	// not usuing adapter
	// mux.Handle("/v1/", logRequests(logger)(muxLogged))
	// using adapter
	mux.Handle("/v1/", Adapt(muxLogged, logRequests(logger), throttleRequests(2, time.Minute)))

	fmt.Printf("listening at %s...\n", addr)
	// log.Fatal(http.ListenAndServe(addr, nil))

	// option 3 - default serve mux had the 3 handlers above inside of it.
	log.Fatal(http.ListenAndServe(addr, mux))
}
