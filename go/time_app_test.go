package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	quit := make(chan bool)

	go func() {
		TimeServer(quit)
	}()
	return func(t *testing.T) {
		quit <- true
	}
}

func TestGetTime(t *testing.T) {
	tearDownTest := setupTestCase(t)
	defer tearDownTest(t)

	requestURL := "http://localhost:8080"

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(requestURL)
	if err != nil {
		t.Fatalf("error testing time server get request: %s\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("time server get request status code error: %s\n", fmt.Sprint(resp.StatusCode))
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading time server response body: %s\n", err)
	}

	bodyString := string(bodyBytes)

	if !strings.Contains(bodyString, "New York") {
		t.Fatalf("time app response missing New York data\n")
	}
}

func TestGetHealth(t *testing.T) {
	tearDownTest := setupTestCase(t)
	defer tearDownTest(t)

	requestURL := "http://localhost:8080/health"

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(requestURL)
	if err != nil {
		t.Fatalf("error testing time server get request: %s\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("time server get request status code error: %s\n", fmt.Sprint(resp.StatusCode))
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading time server response body: %s\n", err)
	}

	bodyString := string(bodyBytes)

	if !strings.Contains(bodyString, "status_code") {
		t.Fatalf("time app health response missing information\n")
	}
}
