package restutil

import (
	"strings"
	"testing"

	"github.com/arschles/assert"
)

func TestAbsPath(t *testing.T) {
	const (
		apiVersionBase = "testbase"
		apiVersion     = "testversion"
		namespace      = "testNS"
		pluralName     = "testplural"
	)
	elts := AbsPath(apiVersionBase, apiVersion, true, namespace, pluralName)
	assert.Equal(t, len(elts), 5, "number of path elts")
	assert.Equal(t, elts[0], "apis", "first path elt")
	assert.Equal(t, elts[1], apiVersionBase, "base API version")
	assert.Equal(t, elts[2], apiVersion, "api version")
	assert.Equal(t, elts[3], "true", "watch path element")
	assert.Equal(t, elts[4], strings.ToLower(pluralName), "plural name")
}
