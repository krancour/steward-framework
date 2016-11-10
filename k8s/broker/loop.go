package broker

import (
	"context"
	"errors"

	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/pkg/watch"
)

var (
	ErrCancelled   = errors.New("stopped")
	ErrNotABroker  = errors.New("not a broker")
	ErrWatchClosed = errors.New("watch closed")
)

// RunLoop starts a blocking control loop that watches and takes action on broker resources
func RunLoop(
	ctx context.Context,
	namespace string,
	fn WatchBrokerFunc,
	cataloger framework.Cataloger,
	createSvcClassFunc CreateServiceClassFunc) error {
	watcher, err := fn(namespace)
	if err != nil {
		return err
	}
	ch := watcher.ResultChan()
	defer watcher.Stop()
	for {
		select {
		case <-ctx.Done():
			return ErrCancelled
		case evt, open := <-ch:
			if !open {
				return ErrWatchClosed
			}
			switch evt.Type {
			case watch.Added:
				if err := handleAddBroker(ctx, cataloger, createSvcClassFunc, evt); err != nil {
					return err
				}
			}
		}
	}
}

func handleAddBroker(
	ctx context.Context,
	cataloger framework.Cataloger,
	createServiceClass CreateServiceClassFunc,
	evt watch.Event) error {
	broker, ok := evt.Object.(*data.Broker)
	if !ok {
		return ErrNotABroker
	}
	svcs, err := cataloger.List(ctx, broker.Spec)
	if err != nil {
		return err
	}

	sClasses, err := translateServiceClasses(broker, svcs)
	if err != nil {
		return err
	}
	for _, sClass := range sClasses {
		if err := createServiceClass(sClass); err != nil {
			return err
		}
	}
	return nil
}
