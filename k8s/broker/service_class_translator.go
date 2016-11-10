package broker

import (
	"fmt"

	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
)

func translateServiceClasses(
	parentBroker *data.Broker,
	svcs []*framework.Service) ([]*data.ServiceClass, error) {

	brokerRef := getObjectReference(parentBroker)
	ret := make([]*data.ServiceClass, len(svcs))
	for i, svc := range svcs {
		ret[i] = &data.ServiceClass{
			TypeMeta: unversioned.TypeMeta{Kind: data.ServiceClassKind},
			ObjectMeta: v1.ObjectMeta{
				Name:      serviceClassName(parentBroker, svc),
				Namespace: parentBroker.Namespace,
			},
			BrokerRef:     *brokerRef,
			ID:            serviceClassID(parentBroker, svc),
			BrokerName:    parentBroker.Name,
			Bindable:      true,
			Plans:         translatePlans(svc.Plans),
			PlanUpdatable: svc.PlanUpdatable,
			Description:   svc.Description,
		}
	}
	return ret, nil
}

func serviceClassName(parentBroker *data.Broker, svc *framework.Service) string {
	return fmt.Sprintf("%s-%s", parentBroker.Name, svc.Name)
}

func serviceClassID(parentBroker *data.Broker, svc *framework.Service) string {
	return fmt.Sprintf("%s-%s", parentBroker.UID, svc.ID)
}

func translatePlans(plans []framework.ServicePlan) []data.ServicePlan {
	ret := make([]data.ServicePlan, len(plans))
	for i, plan := range plans {
		ret[i] = data.ServicePlan{ID: plan.ID, Name: plan.Name, Description: plan.Description}
	}
	return ret
}