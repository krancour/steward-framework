package broker

import (
	"context"
	"testing"

	"github.com/arschles/assert"
	"github.com/deis/steward-framework/fake"
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
	assert.Equal(t, len(createdSvcClasses), 0, "number of create svc classes")
}
