package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

const (
	IsCorruptEnabled = true
	IsDelayEnabled   = true
	DelayDuration    = 1 * time.Second
	ServerUrl        = "http://localhost:8081"
)

func main() {
	target, err := url.Parse(ServerUrl)
	if err != nil {
		log.Fatalf("Invalid server URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ModifyResponse = func(resp *http.Response) error {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()

		corrupted := append([]byte("CORRUPTDATA"), bodyBytes...)
		resp.Body = io.NopCloser(bytes.NewBuffer(corrupted))

		resp.ContentLength = int64(len(corrupted))
		resp.Header.Set("Content-Length", strconv.Itoa(len(corrupted)))

		return nil
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if IsDelayEnabled {
			delay()
		}
		proxy.ServeHTTP(w, r)
	})

	log.Println("Proxy running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func delay() {
	log.Printf("Delaying request by %v seconds", DelayDuration)
	time.Sleep(DelayDuration)
}
