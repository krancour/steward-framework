package broker

import (
	"github.com/deis/steward-framework/k8s/data"
	"github.com/deis/steward-framework/k8s/restutil"
	"k8s.io/client-go/rest"
)

// CreateServiceClassFunc is the function that can successfully create a ServiceClass
type CreateServiceClassFunc func(*data.ServiceClass) error

// NewK8sCreateServiceClassFunc returns a CreateServiceClassFunc implemented with restIFace
func NewK8sCreateServiceClassFunc(restIface rest.Interface) CreateServiceClassFunc {
	return func(sClass *data.ServiceClass) error {
		url := restutil.AbsPath(
			restutil.APIVersionBase,
			restutil.APIVersion,
			false,
			sClass.Namespace,
			data.ServiceClassKindPlural,
		)
		return restIface.Post().AbsPath(url...).Do().Error()
	}
}

// returns the function and a mutable slice of classes that were created. if retErr != nil,
// it is always returned by the function and the returned slice is never modified
func newFakeCreateServiceClassFunc(retErr error) (CreateServiceClassFunc, []*data.ServiceClass) {
	createdClasses := []*data.ServiceClass{}
	retFn := func(sClass *data.ServiceClass) error {
		if retErr != nil {
			return retErr
		}
		createdClasses = append(createdClasses, sClass)
		return nil
	}
	return retFn, createdClasses
}
