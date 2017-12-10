package main

import (
    "fmt"
	"log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {	
	w.Write([]byte(fmt.Sprintf("pong")))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Serving fast-backend 0.0.0.0:8080");
	http.ListenAndServe("0.0.0.0:8080", nil)
}
