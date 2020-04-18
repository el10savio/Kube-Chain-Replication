package kubeinterface

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// configCluster wrapper configures
// the Kubernetes cluster
func configCluster() (config *rest.Config, err error) {
	config, err = rest.InClusterConfig()
	if err != nil {
		return
	}
	return
}

// createClientset creates the clientset
// with config as the input parameter
func createClientset(config *rest.Config) (clientset *kubernetes.Clientset, err error) {
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	return
}

// GetPods interface gets the podIPs of a given
// pod name & namespace in the cluster
func GetPods(podName string, namespace string) ([]string, error) {
	config, err := configCluster()
	if err != nil {
		return []string{}, err
	}

	clientset, err := createClientset(config)
	if err != nil {
		return []string{}, err
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return []string{}, err
	}

	podIPs := make([]string, 0)

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, podName) && pod.Status.Phase == "Running" {
			podIPs = append(podIPs, pod.Status.PodIP)
		}
	}

	return podIPs, nil
}
