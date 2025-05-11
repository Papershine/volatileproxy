package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	for {
		log.Println("Sending request")
		resp, err := client.Get("http://localhost:8080")
		if err != nil {
			log.Fatal("Request failed:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		log.Println("Status:", resp.Status)
		log.Println("Body:", string(body))

		log.Println("Press Enter to restart")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		log.Println("Restarting")
	}
}
