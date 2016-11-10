package broker

import (
	"fmt"
	"testing"

	"github.com/arschles/assert"
	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/pkg/api/v1"
)

func TestServiceClassName(t *testing.T) {
	broker := &data.Broker{
		ObjectMeta: v1.ObjectMeta{Name: "testBroker"},
	}
	svc := &framework.Service{
		ServiceInfo: framework.ServiceInfo{Name: "testSvc"},
	}
	name := serviceClassName(broker, svc)
	assert.Equal(t, name, fmt.Sprintf("%s-%s", broker.Name, svc.Name), "service class name")
}

func TestServiceClassID(t *testing.T) {
	broker := &data.Broker{
		ObjectMeta: v1.ObjectMeta{Name: "testBroker", UID: "testUID"},
	}
	svc := &framework.Service{
		ServiceInfo: framework.ServiceInfo{Name: "testSvc", ID: "testID"},
	}
	id := serviceClassID(broker, svc)
	assert.Equal(t, id, fmt.Sprintf("%s-%s", broker.UID, svc.ID), "service class ID")
}

func TestTranslatePlans(t *testing.T) {

}
