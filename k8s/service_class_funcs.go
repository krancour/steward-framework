package k8s

import (
	"encoding/json"

	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/1.4/rest"
)

type ServiceClassListFunc func(string) ([]data.ServiceClass, error)

type ServiceClassCreateFunc func(*data.ServiceClass) (*data.ServiceClass, error)

func NewK8sServiceClassListFunc(rc *rest.Client) ServiceClassListFunc {
	return func(ns string) ([]data.ServiceClass, error) {
		req := rc.Get().AbsPath(getTPRPath(ns, data.ServiceClassTPDName))
		logger.Debugf("making request to %s", req.URL().String())
	}
}
