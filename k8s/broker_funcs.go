package k8s

import (
	"github.com/deis/steward-framework/k8s/data"
)

type BrokerUpdateFunc func(*data.Broker) error
