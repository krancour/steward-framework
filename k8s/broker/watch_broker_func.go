package broker

import (
	"k8s.io/client-go/pkg/watch"
	"k8s.io/client-go/rest"
)

// WatchBrokerFunc is the function that returns a watch interface for broker resources
type WatchBrokerFunc func() (watch.Interface, error)

// NewK8sWatchBrokerFunc returns a WatchBrokerFunc backed by a Kubernetes client
func NewK8sWatchBrokerFunc(restIface rest.Interface) WatchBrokerFunc {
	return func() (watch.Interface, error) {
		return nil, nil
	}
}
