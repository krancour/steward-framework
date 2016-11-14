package refs

import (
	"github.com/deis/steward-framework/k8s/data"
	"github.com/deis/steward-framework/k8s/restutil"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/rest"
)

// NewK8sBrokerGetterFunc returns a BrokerGetterFunc  backed by a real kubernetes client
func NewK8sBrokerGetterFunc(restIface rest.Interface) BrokerGetterFunc {
	return func(ref api.ObjectReference) (*data.Broker, error) {
		ret := new(data.Broker)
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
