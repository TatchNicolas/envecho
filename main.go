package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, pair := range os.Environ() {
		key := strings.SplitN(pair, "=", 2)[0]
		prefix := strings.SplitN(key, "_", 2)[0]
		if prefix == "ENVECHO" {
			fmt.Printf("%s\n", pair)
		}
	}
	host, ok := os.LookupEnv("ENVECHO_HOST")
	if !ok {
		host = "0.0.0.0"
	}
	port, ok := os.LookupEnv("ENVECHO_PORT")
	if !ok {
		port = "80"
	}
	message, ok := os.LookupEnv("ENVECHO_MESSAGE")
	if !ok {
		message = "Set ENVECHO_MESSAGE to configure response text"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
                log.Println(r.RequestURI)
	})
	http.ListenAndServe(address, nil)
}
