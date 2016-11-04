package data

import (
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
	"k8s.io/client-go/1.4/pkg/api/v1"
)

const (
	ServiceClassTPRName = "ServiceClass"
)

type ServiceClass struct {
	unversioned.TypeMeta `json:",inline"`
	v1.ObjectMeta        `json:"metadata,omitempty"`
	BrokerRef            api.ObjectReference `json:"broker_ref"`
	ID                   string              `json:"id"`
	BrokerName           string              `json:"broker_name"`
	Bindable             bool                `json:"bindable"`
	Plans                []ServicePlan       `json:"plans"`
	PlanUpdatable        bool                `json:"updatable"`
	Description          string              `json:"description"`
}

type ServicePlan struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
