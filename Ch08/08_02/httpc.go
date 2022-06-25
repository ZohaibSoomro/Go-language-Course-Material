// Making HTTP calls
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"bytes"
	"encoding/json"
)

// Job is a job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	/*// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)*/

	// POST request
	job := &Job{
		User:   "Ali",
		Action: "punch",	//visiting an endpoint
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't encode job - %s", err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

	