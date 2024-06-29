package fooer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assert is similar to t.Error*
func TestFooerWithTestify(t *testing.T) {
	assert.Equal(t, "Foo", Fooer(0), "0 is divisible by 3, should return Foo")
	assert.NotEqual(t, "Foo", Fooer(1), "1 is not divisible by 3, should not return Foo")
}

// require is similar to t.Fatal*
func TestMapWithTestify(t *testing.T) {
	require.Equal(t, map[int]string{1: "1", 2: "2"}, map[int]string{1: "1", 2: "2"})
}
