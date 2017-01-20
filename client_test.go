package sib

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {

	noKey := ""
	client, err := NewClient(noKey)

	if err == nil {
		t.Error("Expected NewClient function to fail without API Key.")
	}

	key := "123"
	client, err = NewClient(key)

	if err != nil {
		t.Error("Expected NewClient to complete without error.")
	}

	if client.apiKey != key {
		t.Error("API key is not being set.")
	}
	if client.Client == nil {
		t.Error("Http client is not being set.")
	}

	if client.Client.Timeout == time.Duration(0) {
		t.Error("Request timeout is not being set.")
	}
}
