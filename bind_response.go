package framework

import (
	"fmt"
)

// BindResponse represents a response to a BindRequest. It contains a map of credentials and other
// connection details for a service instance.
type BindResponse struct {
	Creds map[string]interface{}
}

// ResponseCredsToSecretData converts a map[string]interface{}, such as one in BindResponse's
// Creds field, into a Kubernetes Secret-compatible map[string][]byte
func ResponseCredsToSecretData(creds map[string]interface{}) map[string][]byte {
	// TODO: make sure the values of the secret are valid DNS subdomains.
	// https://tools.ietf.org/html/rfc4648#section-4 for the spec
	// (steal logic from the k8s client or steward here)
	secretData := make(map[string][]byte, len(creds))
	for k, v := range creds {
		vStr := fmt.Sprintf("%s", v)
		secretData[k] = []byte(vStr)
	}
	return secretData
}
