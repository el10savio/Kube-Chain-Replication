package chain

import (
	"sort"

	kube "../kubeinterface"
)

// GetPodsList ...
func GetPodsList() ([]string, error) {
	pods, err := kube.GetPods("goredis", "craq")
	if err != nil {
		return []string{}, err
	}

	sort.Strings(pods)

	return pods, nil
}

// GetHead ...
func GetHead() (string, error) {
	pods, err := GetPodsList()
	if err != nil {
		return "", err
	}

	return pods[0], nil
}

// GetTail ...
func GetTail() (string, error) {
	pods, err := GetPodsList()
	if err != nil {
		return "", err
	}

	return pods[len(pods)-1], nil
}
