package k8s

import (
	"context"

	"github.com/deis/steward-framework"
	"github.com/deis/steward-framework/k8s/broker"
	"github.com/deis/steward-framework/k8s/instance"
	"github.com/deis/steward-framework/k8s/refs"
	"k8s.io/client-go/kubernetes"
)

// StartControlLoops is a convenience function that initiates all of Steward's individual control
// loops, of which there is one per relevant resource type-- Broker, Instance, and Binding.
//
// In order to stop all loops, pass a cancellable context to this function and call its cancel
// function.
func StartControlLoops(
	ctx context.Context,
	k8sClient *kubernetes.Clientset,
	cataloger framework.Cataloger,
	lifecycler framework.Lifecycler,
	globalNamespace string,
	errCh chan<- error,
) {
	restClient := k8sClient.CoreClient.RESTClient()

	// Start broker loop
	go func() {
		updateBrokerFn := broker.NewK8sUpdateBrokerFunc(restClient)
		watchBrokerFn := broker.NewK8sWatchBrokerFunc(restClient)
		createSvcClassFn := broker.NewK8sCreateServiceClassFunc(restClient)
		if err := broker.RunLoop(
			ctx,
			globalNamespace,
			watchBrokerFn,
			updateBrokerFn,
			cataloger,
			createSvcClassFn,
		); err != nil {
			errCh <- err
		}
	}()

	// Start instance loop
	go func() {
		if err := instance.RunLoop(
			ctx,
			instance.NewK8sWatchInstanceFunc(restClient),
			instance.NewK8sUpdateInstanceFunc(restClient),
			refs.NewK8sServiceClassGetterFunc(restClient),
			refs.NewK8sBrokerGetterFunc(restClient),
			lifecycler,
		); err != nil {
			errCh <- err
		}
	}()

	// TODO: Start binding loop

	<-ctx.Done()
}
