package chain

import (
	"sort"

	kube "../kubeinterface"
)

// GetPodsList interacts withe kube client 
// to obtain the goredis node IPs 
// in the craq namespace
func GetPodsList() ([]string, error) {
	pods, err := kube.GetPods("goredis", "craq")
	if err != nil {
		return []string{}, err
	}

	sort.Strings(pods)

	return pods, nil
}

// GetHead gets the HEAD of the chain
// It is string sorted so this means 
// 172.2.0.17 is lesser than 172.2.0.8
func GetHead() (string, error) {
	pods, err := GetPodsList()
	if err != nil {
		return "", err
	}

	return pods[0], nil
}

// GetTail gets the TAIL of the chain
func GetTail() (string, error) {
	pods, err := GetPodsList()
	if err != nil {
		return "", err
	}

	return pods[len(pods)-1], nil
}
