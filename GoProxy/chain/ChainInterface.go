package chain

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"../store"
)

const (
	// PORT ...
	PORT = ":8080"
)

// HeadSetValue ...
func HeadSetValue(HEAD string, store store.Entry) (int, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	bytesStore, err := json.Marshal(store)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	url := "http://" + strings.TrimSpace(HEAD) + strings.TrimSpace(PORT) + "/store/set"
	response, err := client.Post(url, "application/json", bytes.NewBuffer(bytesStore))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return response.StatusCode, err
}

// TailGetValue ...
func TailGetValue(TAIL string, key string) (store.Entry, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	url := "http://" + strings.TrimSpace(TAIL) + strings.TrimSpace(PORT) + "/store/get/" + strings.TrimSpace(key)
	response, err := client.Get(url)
	if err != nil {
		return store.Entry{}, err
	}

	var store store.Entry
	_ = json.NewDecoder(response.Body).Decode(&store)

	return store, err
}

// GetPodHealth ...
func GetPodHealth(Pod string) (NodeStatus, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	url := "http://" + strings.TrimSpace(Pod) + strings.TrimSpace(PORT) + "/health"
	response, err := client.Get(url)
	if err != nil {
		return NodeStatus{}, err
	}

	status := NodeStatus{
		Node:    Pod,
		Healthy: response.StatusCode == http.StatusOK,
	}

	return status, err
}

// UpdatePodNeighbor ...
func UpdatePodNeighbor(Pod string, Neighbor string) error {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	url := "http://" + strings.TrimSpace(Pod) + strings.TrimSpace(PORT) + "/neighbor/" + strings.TrimSpace(Neighbor)
	_, err := client.Get(url)
	if err != nil {
		return err
	}

	return nil
}
