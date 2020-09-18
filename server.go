package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	type Res struct {
		Time string `json:"time"`
	}

	http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) {
		res := Res{time.Now().Format(time.RFC3339)}
		json, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	})

	log.Println("Starting HTTP server on localhost:8795")

	http.ListenAndServe(":8795", nil)
}
