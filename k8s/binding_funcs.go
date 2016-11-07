package k8s

import (
	"github.com/deis/steward-framework/k8s/data"
)

type BindingUpdateFunc func(*data.Binding) error
