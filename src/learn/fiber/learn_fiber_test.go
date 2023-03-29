package main

import (
	"io"
	"net/http"
	"testing"
)

func doRequest(url string) (responseContent string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	respContent := string(body)
	return respContent, nil
}

func TestDirect(t *testing.T) {
	ws := NewWebServer("localhost", 10081)
	defer ws.Stop()
	go ws.Run()

	respContent, err := doRequest("http://localhost:10081/direct")
	if err != nil {
		t.Fatalf("failed do request, error: %s", err.Error())
	}
	expected := "Hello, World!"
	if respContent != expected {
		t.Fatalf("'%s' is expected, but '%s' is received", expected, respContent)
	}
}

func TestSimple(t *testing.T) {
	ws := NewWebServer("localhost", 10081)
	defer ws.Stop()
	go ws.Run()

	type TestCase struct {
		name     string
		expected string
	}
	testCases := []TestCase{
		{
			name:     "Jack",
			expected: "Hello, Jack!",
		},
		{
			name:     "Lily",
			expected: "Hello, Lily!",
		},
	}

	for _, testCase := range testCases {
		respContent, err := doRequest("http://localhost:10081/simple/" + testCase.name)
		if err != nil {
			t.Fatalf("failed do request, error: %s", err.Error())
		}
		if respContent != testCase.expected {
			t.Fatalf("'%s' is expected, but '%s' is received", testCase.expected, respContent)
		}
	}
}
