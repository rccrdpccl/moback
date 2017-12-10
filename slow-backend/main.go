package main

import (
    "fmt"
	"log"
    "net/http"
	"math/rand"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {	
	time.Sleep(time.Duration(rand.Int31n(10) + 1000) * time.Millisecond)
	w.Write([]byte(fmt.Sprintf("pong")))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Serving slow-backend 0.0.0.0:8080");
	http.ListenAndServe("0.0.0.0:8080", nil)
}
