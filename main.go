package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/potch8228/go-echo-ip/models"
)

func main() {
	http.HandleFunc("/", routerHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Listen port will be %s", port)
	}

	log.Printf("Set listen port to %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func routerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=utf-8")

	if r.URL.Path == "/ip" {
		ipHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Printf("Header Key: %v = %v", k, v)
	}

	ip := models.Ip{V4: "IPv4 address", V6: "IPv6 address"}

	var data []byte
	var err error
	if data, err = json.Marshal(ip); err != nil {
		msg := fmt.Sprintf("Json Marshaling Failure: %v", err)

		log.SetOutput(os.Stderr)
		log.Printf(msg)

		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(data); err != nil {
		msg := fmt.Sprintf("Write Response Failure: %v", err)

		log.SetOutput(os.Stderr)
		log.Printf(msg)

		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
