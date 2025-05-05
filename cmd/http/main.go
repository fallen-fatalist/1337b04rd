package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", temp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func temp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HELLO 1337bo4rd!!!")
}
