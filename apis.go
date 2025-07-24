package main 

import (
	"fmt"
	"net/http"
)

type Homehandler struct {}

func (h Homehandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to the go server")
}

func main() {
	
}