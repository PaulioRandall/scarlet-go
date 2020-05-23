package z_matching

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkPanic(t *testing.T, f func()) {
	require.Panics(t, f, "Expected a panic")
}
