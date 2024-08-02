package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client sends a request to the server and processes the response
func Client() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-MyClient", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	//fmt.Println(res.Header.Get("Content-Type"))

	var response struct {
		Message string `json:"message"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response from server: %s\n", response.Message)
}
