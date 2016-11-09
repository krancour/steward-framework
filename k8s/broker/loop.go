package broker

import (
	"context"
	"errors"

	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/k8s/data"
	"k8s.io/client-go/pkg/watch"
)

var (
	ErrCancelled  = errors.New("stopped")
	ErrNotABroker = errors.New("not a broker")
)

// WatchBrokerFunc is the function that returns a watch interface for broker resources
type WatchBrokerFunc func() (watch.Interface, error)

// RunLoop starts a blocking control loop that watches and takes action on broker resources
func RunLoop(ctx context.Context, fn WatchBrokerFunc, cataloger framework.Cataloger) error {
	watcher, err := fn()
	if err != nil {
		return err
	}
	ch := watcher.ResultChan()
	defer watcher.Stop()
	for {
		select {
		case <-ctx.Done():
			return ErrCancelled
		case evt := <-ch:
			switch evt.Type {
			case watch.Added:
				if err := handleAddBroker(cataloger, evt); err != nil {
					return err
				}
			}
		}
	}
}

func handleAddBroker(ctx, cataloger framework.Cataloger, evt watch.Event) error {
	broker, ok := evt.Object.(data.Broker)
	if !ok {
		return ErrNotABroker
	}
	svcs, err := cataloger.List(ctx, broker.Spec)
	if err != nil {
		return err
	}

	// sClasses := translateServiceClasses(svcs)
	// TODO: write sClasses to k8s
	return nil

}

func translateServiceClasses(svcs []*framework.Service) []*data.ServiceClass {
	return nil
}
