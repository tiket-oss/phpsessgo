package phpsessgo

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
)

func TestRedisSessionHandler(t *testing.T) {
	s, err := miniredis.Run()
	require.NoError(t, err)
	defer s.Close()

	handler := &RedisSessionHandler{
		Client: redis.NewClient(&redis.Options{
			Addr: s.Addr(),
		}),
		RedisKeyPrefix: "PHPREDIS_SESSION:",
	}
	defer handler.Close()

	t.Run("read data", func(t *testing.T) {
		s.Set("PHPREDIS_SESSION:some-sessionID", "some-data")

		data, err := handler.Read("some-sessionID")
		require.NoError(t, err)
		require.Equal(t, "some-data", data)
	})

	t.Run("read not existing data", func(t *testing.T) {
		s.Set("PHPREDIS_SESSION:some-sessionID", "some-data")

		data, err := handler.Read("not-exist")
		require.NoError(t, err)
		require.Equal(t, "", data)
	})

	t.Run("write data", func(t *testing.T) {
		err := handler.Write("some-sessionID-2", "some-data-2")
		require.NoError(t, err)

		val, _ := s.Get("PHPREDIS_SESSION:some-sessionID-2")
		require.Equal(t, "some-data-2", val)
	})

}
