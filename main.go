package main

import (
	"io"
	"log"
	"net/http"
)

var cunstomTransport = http.DefaultTransport

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleReq),
	}

	log.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy on: ", err)
	}
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	targetUrl := r.URL

	proxyReq, err := http.NewRequest(r.Method, targetUrl.String(), nil)
	if err != nil {
		http.Error(w, "Error creating a proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, v := range values {
			proxyReq.Header.Add(name, v)
		}
	}

	resp, err := cunstomTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(name, v)
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
