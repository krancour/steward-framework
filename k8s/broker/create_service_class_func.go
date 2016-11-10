package broker

import (
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/rest"
)

// CreateServiceClassFunc is the function that can successfully create a ServiceClass
type CreateServiceClassFunc func(*data.ServiceClass) error

func NewK8sCreateServiceClassFunc(restIface rest.Interface) CreateServiceClassFunc {
	return func(sClass *data.ServiceClass) error {
		return nil
	}
}
