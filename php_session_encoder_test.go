package phpsessgo_test

import (
	"testing"

	"github.com/tiket-oss/phpsessgo"
	"github.com/stretchr/testify/require"
)

func TestPHPSessionEncoder(t *testing.T) {
	raw := `spike01|s:6:"data01";angka|i:987654321;array01|a:5:{i:1;s:2:"i2";i:2;s:2:"i3";i:3;s:2:"i4";i:4;s:2:"i5";i:0;s:2:"i1";}person|a:2:{s:10:"first_name";s:4:"iman";s:9:"last_name";s:4:"tung";}bool01|b:1;`

	encoder := phpsessgo.PHPSessionEncoder{}
	session, err := encoder.Decode(raw)
	require.NoError(t, err)

	encoded, err := encoder.Encode(session)
	require.NoError(t, err)
	require.Equal(t, len(raw), len(encoded))

}
