package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed")
		w.Write([]byte("Request processed"))
	case <-ctx.Done():
		log.Println("Request cancelled: " + ctx.Err().Error())
	}
}
