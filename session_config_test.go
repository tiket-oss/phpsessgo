package phpsessgo_test

import (
	"testing"

	"github.com/imantung/phpsessgo"
	"github.com/stretchr/testify/require"
)

func TestSessionConfig_Default(t *testing.T) {
	config := phpsessgo.DefaultSessionConfig()
	require.Equal(t, phpsessgo.DefaultSessionName, config.Name)
}
