package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// define types and structs after import
type zip struct {
	// the fields that we want to fill with JSON must be exported for the JSON file to actually see it
	Zip   string `json:"zip"`
	City  string `json:"city"`
	State string `json:"state"`
}

type zipSlice []*zip              // [*zip, *zip, ...]
type zipIndex map[string]zipSlice // {"seattle":[*zip, *zip, ...], "la":[*zip, *zip, ...]}

// sign quotes are used for a go "rune" which is a single character (which covers characters of multiple bytes)
// back ticks are used to also preverse line breaks

//varName type 		* means it is sending a reference and not a copy
// 					we are passing a pointer to the request
// 					in go, any type can be passed by value and any type can be passed by reference
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// must add headers before you start writing the body
	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("hello " + name))
}

func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {
	// /zips/city/seattle
	_, city := path.Split(r.URL.Path)
	lcity := strings.ToLower(city)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(zi[lcity]); err != nil {
		http.Error(w, "error encoding json: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		log.Fatal("please set addr env var")
	}

	// f, err := os.Open("../data/zips.json")
	f, err := os.Open("zips.json")
	if err != nil {
		log.Fatal("Error opening zips file: " + err.Error())
	}

	zips := make(zipSlice, 0, 43000) // [*zip, *zip, ...]
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&zips); err != nil {
		log.Fatal("error decoding zips json: " + err.Error())
	}
	fmt.Printf("loaded %d zips\n", len(zips))

	zi := make(zipIndex) // {"seattle":[*zip, *zip, ...], "la":[*zip, *zip, ...]}

	// index, *zip
	for _, value := range zips {
		lower := strings.ToLower(value.City)
		// (array of zips).append(new zip)
		zi[lower] = append(zi[lower], value)
	}
	fmt.Printf("there are %d zips in seattle\n", len(zi["seattle"]))

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/zips/city/", zi.zipsForCityHandler)

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

/*
compile with "go install" and run with "zipsvr"
$ go install && zipsvr
*/
