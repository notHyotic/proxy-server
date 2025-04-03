package main

import (
	"io"
	"log"
	"net/http"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	port := ":8080"
	http.Handle("/", &Proxy{})
	log.Printf("Serving proxy on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
