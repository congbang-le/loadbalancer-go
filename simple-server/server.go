package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, "This is "+hostname+"\n")
	return
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]string)
	res["You good?"] = "Cannot be better"
	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonRes)
	return
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", healthz)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
