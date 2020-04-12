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
	// PORT defines the port value 
	// for the GoServer service
	PORT = ":8080"
)

// HeadSetValue sends a write request 
// to the HEAD of the chain
func HeadSetValue(HEAD string, store store.Entry) (int, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	// json encoded store
	bytesStore, err := json.Marshal(store)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// send Set POST request to HEAD
	url := "http://" + strings.TrimSpace(HEAD) + strings.TrimSpace(PORT) + "/store/set"
	response, err := client.Post(url, "application/json", bytes.NewBuffer(bytesStore))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return response.StatusCode, err
}

// TailGetValue sends a read request 
// to the TAIL of the chain
func TailGetValue(TAIL string, key string) (store.Entry, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	// send GET request to TAIL
	url := "http://" + strings.TrimSpace(TAIL) + strings.TrimSpace(PORT) + "/store/get/" + strings.TrimSpace(key)
	response, err := client.Get(url)
	if err != nil {
		return store.Entry{}, err
	}

	// json decoded store
	var store store.Entry
	_ = json.NewDecoder(response.Body).Decode(&store)

	return store, err
}

// GetPodHealth checks if the  connection 
// to a given pod in the chain and
// its redis service is alright
func GetPodHealth(Pod string) (NodeStatus, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	// send healthcheck GET request to node
	url := "http://" + strings.TrimSpace(Pod) + strings.TrimSpace(PORT) + "/health"
	response, err := client.Get(url)
	if err != nil {
		return NodeStatus{}, err
	}

	// create NodeStatus entry
	status := NodeStatus{
		Node:    Pod,
		Healthy: response.StatusCode == http.StatusOK,
	}

	return status, err
}

// UpdatePodNeighbor sends a node 
// the IP address of its 
// respective neighbor 
// to update locally 
func UpdatePodNeighbor(Pod string, Neighbor string) error {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	// send neighbor update GET request to node
	url := "http://" + strings.TrimSpace(Pod) + strings.TrimSpace(PORT) + "/neighbor/" + strings.TrimSpace(Neighbor)
	_, err := client.Get(url)
	if err != nil {
		return err
	}

	return nil
}
