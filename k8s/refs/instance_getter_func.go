package refs

import (
	"github.com/deis/steward-framework/k8s/data"
	"github.com/deis/steward-framework/k8s/restutil"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/rest"
)

// InstanceGetterFunc is the function that attempts to fetch an instance at the given object ref
type InstanceGetterFunc func(api.ObjectReference) (*data.Instance, error)

// NewK8sInstanceGetterFunc returns an InstanceGetterFunc backed by a real kubernetes client
func NewK8sInstanceGetterFunc(restIface rest.Interface) InstanceGetterFunc {
	return func(ref api.ObjectReference) (*data.Instance, error) {
		ret := new(data.Instance)
		url := restutil.AbsPath(
			restutil.APIVersionBase,
			restutil.APIVersion,
			false,
			ref.Namespace,
			ref.Name,
		)
		if err := restIface.Get().AbsPath(url...).Do().Into(ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
}
