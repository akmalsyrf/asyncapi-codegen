//go:generate go run ../../../../cmd/asyncapi-codegen -p issue300 -g types -i ./asyncapi.yaml -o ./asyncapi.gen.go

package issue300

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpenObject_IsMap(t *testing.T) {
	// This should compile if OpenObjectSchema underlying type is map[string]any.
	var _ OpenObjectSchema = map[string]any{
		"foo": "bar",
		"n":   123,
	}
}

func TestOneOf_Generated(t *testing.T) {
	var u UnionSchema
	s := "abc"
	u.AnyOf0 = &s
	require.NotNil(t, u.AnyOf0)
}
