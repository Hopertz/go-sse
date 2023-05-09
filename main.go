// /main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	name   string
	events chan uint
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	client := &Client{name: r.RemoteAddr, events: make(chan uint, 10)}
	go updateNumber(client)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(1 * time.Second)
	select {
	case ev := <-client.events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updateNumber(client *Client) {
	for {
		db := uint(rand.Uint32())
		client.events <- db
	}

}

func main() {
	http.HandleFunc("/sse", sseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
