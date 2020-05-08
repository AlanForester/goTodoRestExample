package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type TestStruct struct {
	Args    struct{}          `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func main() {
	s := new(TestStruct)
	if err := getJson("https://httpbin.org/get", s); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Result: %#v", s)
}

func getJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
