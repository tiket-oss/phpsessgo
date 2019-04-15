package phpsessgo_test

import (
	"regexp"
	"testing"

	"github.com/tiket-oss/phpsessgo"
	"github.com/stretchr/testify/require"
)

func TestUUIDCreator(t *testing.T) {
	creator := phpsessgo.UUIDCreator{}
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	require.True(t, r.MatchString(creator.CreateSID()))
}
