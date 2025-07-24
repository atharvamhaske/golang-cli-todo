// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type HomeHandler struct {

// }

// func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "welcome to the go server")
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/", HomeHandler{})
// 	mux.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintln(w, "Hello from /hello route")
// 	})

// 	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request){
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"status" : "ok}`)
// 	})

// 	fmt.Println("server starting on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }
