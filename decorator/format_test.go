package decorator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestFormat_EmptyCurrentContext(t *testing.T) {
	require.Equal(t, "", Format(api.Config{}, "none"))
}

func TestFormat_EmptyContextNotFound(t *testing.T) {
	require.Equal(t, "", Format(api.Config{
		CurrentContext: "foo",
		Contexts: map[string]*api.Context{
			"bar": {Cluster: "bar", Namespace: "fizzbuz"},
		},
	}, "none"))
}

func TestFormat_DefaultIfNamespaceNotSet(t *testing.T) {
	require.Equal(t, "[☸️ foo:default]", Format(api.Config{
		CurrentContext: "foo",
		Contexts:       map[string]*api.Context{"foo": {Cluster: "foo"}},
	}, "none"))
}

func TestFormat(t *testing.T) {
	require.Equal(t, "[☸️ foo:fizzbuzz]", Format(api.Config{
		CurrentContext: "foo",
		Contexts: map[string]*api.Context{
			"foo": {Cluster: "foo", Namespace: "fizzbuzz"},
		},
	}, "none"))
}
