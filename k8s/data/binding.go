package data

import (
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/api/v1"
)

const (
	BindingTPRName = "Binding"
)

type BindingState string

const (
	BindingStatePending BindingState = "Pending"
	BindingStateBound   BindingState = "Bound"
	BindingStateFailed  BindingState = "Failed"
)

type Binding struct {
	kunversioned.TypeMeta
	kapi.ObjectMeta

	Spec   BindingSpec
	Status BindingStatus
}

type BindingSpec struct {
	InstanceRef api.ObjectReference    `json:"instance_ref"`
	Parameters  map[string]interface{} `json:"parameters"`
	SecretName  string                 `json:"secret_name"`
}

type BindingStatus struct {
	State BindingState
}
