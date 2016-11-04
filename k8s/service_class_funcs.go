package k8s

import (
	"github.com/deis/steward-framework/k8s/data"
)

type ServiceClassListFunc func() ([]data.ServiceClass, error)

type ServiceClassCreateFunc func(*data.ServiceClass) (*data.ServiceClass, error)
