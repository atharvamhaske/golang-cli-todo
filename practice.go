package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/gay", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello you are gay, stfu and you are accessing %s and using user agent %s\n", r.URL.Path, r.Header.Get("User-Agent"))
	})

	mux.HandleFunc("/athx", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w,"this is atharva here" )
	})

	mux.HandleFunc("/status", func (w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status" : "okayish}`)
	})

	fmt.Println("server running on 8000")
	http.ListenAndServe(":8000", mux)
}