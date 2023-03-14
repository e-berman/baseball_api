package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello Jackie Robinson")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})
	
	http.HandleFunc("/barry_bonds", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Barry Bonds")
	})

	http.ListenAndServe(":4242", nil)
}
