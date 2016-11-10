package broker

import (
	"context"
	"testing"

	"github.com/arschles/assert"
	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/fake"
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/watch"
)

func TestHandleAddBrokerNotABroker(t *testing.T) {
	ctx := context.Background()
	cataloger := fake.Cataloger{}
	createSvcClass, createdSvcClasses := newFakeCreateServiceClassFunc(nil)
	evt := watch.Event{
		Type:   watch.Added,
		Object: &api.Pod{},
	}
	err := handleAddBroker(ctx, cataloger, createSvcClass, evt)
	assert.Err(t, ErrNotABroker, err)
	assert.Equal(t, len(*createdSvcClasses), 0, "number of create svc classes")
}

func TestHandleAddBrokerSuccess(t *testing.T) {
	ctx := context.Background()
	cataloger := fake.Cataloger{
		Services: []*framework.Service{
			&framework.Service{
				ServiceInfo: framework.ServiceInfo{
					Name:          "testName",
					ID:            "testID",
					Description:   "testDescr",
					PlanUpdatable: true,
				},
				Plans: []framework.ServicePlan{
					framework.ServicePlan{ID: "tid1", Name: "tName1", Description: "tDesc1", Free: true},
					framework.ServicePlan{ID: "tid2", Name: "tName2", Description: "tDesc2", Free: false},
					framework.ServicePlan{ID: "tid3", Name: "tName3", Description: "tDesc3", Free: true},
				},
			},
		},
	}
	createSvcClass, createdSvcClasses := newFakeCreateServiceClassFunc(nil)
	evt := watch.Event{
		Type:   watch.Added,
		Object: &data.Broker{},
	}
	err := handleAddBroker(ctx, cataloger, createSvcClass, evt)
	assert.NoErr(t, err)
	assert.Equal(t, len(*createdSvcClasses), len(cataloger.Services), "number of create svc classes")

}
