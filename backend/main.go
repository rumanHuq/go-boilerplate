package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/backend/mirrors"
)

type response struct {
	FastestURL string        `json:fastest_url`
	Latency    time.Duration `json:latency`
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	response := findFastest(mirrors.MirrorList)
	respJSON, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest-mirror", handleFunc)
	server := &http.Server{
		Addr:           ":3000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Listening on port 3000")
	log.Fatal(server.ListenAndServe())
}
