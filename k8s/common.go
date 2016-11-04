package k8s

import (
	"fmt"
)

const (
	// TODO: consider making this configurable
	resourceAPIVersionBase = "steward.deis.io"
	apiVersionV1           = "v1"
)

func resourceAPIVersion(v string) string {
	return fmt.Sprintf("%s/%s", resourceAPIVersionBase, v)
}

func getTPRPath(namespace, tprName string) []string {
	return []string{
		"apis",
		resourceAPIVersionBase,
		apiVersionV1,
		"namespaces",
		namespace,
		tprName,
	}
}
