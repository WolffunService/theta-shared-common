package common

import "k8s.io/client-go/rest"

var onCloud = false

func init() {
	onCloud = internalCheckOnCloud()
}

func internalCheckOnCloud() bool {
	// Check if the code is running inside a Kubernetes cluster
	_, err := rest.InClusterConfig()

	// Return true if running inside the K8s cluster, false otherwise
	return err == nil
}

func OnCloud() bool {
	return onCloud
}
