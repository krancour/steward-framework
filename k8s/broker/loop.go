package broker

import (
	"context"

	"github.com/deis/steward-framework"
	"k8s.io/client-go/pkg/watch"
)

// WatchBrokerFunc is the function that returns a watch interface for broker resources
type WatchBrokerFunc func() (watch.Interface, error)

// RunLoop starts a blocking control loop that watches and takes action on broker resources
func RunLoop(ctx context.Context, fn WatchBrokerFunc, lifecycler framework.Lifecycler) error {
	return nil
}
