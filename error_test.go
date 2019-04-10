package phpsessgo_test

import (
	"testing"

	"github.com/imantung/phpsessgo"
	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	err := phpsessgo.Error("some-error")
	require.Equal(t, "some-error", err.Error())
}
