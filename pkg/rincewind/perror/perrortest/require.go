package perrortest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func RequireNil(t *testing.T, e error) {

	if e == nil {
		return
	}

	s := fmt.Sprintf("%+v", e)
	require.Fail(t, s)
}
