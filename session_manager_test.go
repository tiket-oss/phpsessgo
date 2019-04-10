package phpsessgo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/imantung/phpsessgo"
	"github.com/imantung/phpsessgo/mock"
	"github.com/stretchr/testify/require"
)

func TestSessionManager_New(t *testing.T) {
	config := phpsessgo.SessionConfig{
		Name: "some-session-name",
	}
	manager, _ := phpsessgo.NewSessionManager(config)
	require.Equal(t, "some-session-name", manager.SessionName)
}

func TestSessionManager_Start(t *testing.T) {
	manager := phpsessgo.SessionManager{
		SessionName: "some-session-name",
	}

	t.Run("generate Session ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sidCreator := mock.NewMockSessionIDCreator(ctrl)
		sidCreator.EXPECT().CreateSID().Return("random-hash")
		manager.SIDCreator = sidCreator

		req, _ := http.NewRequest(http.MethodGet, "some-url", nil)
		rr := httptest.NewRecorder()

		session, err := manager.Start(rr, req)
		require.NoError(t, err)
		require.Equal(t, "random-hash", session.SessionID())
		require.Equal(t, "some-session-name=random-hash", rr.HeaderMap.Get("Set-Cookie"))
	})

	t.Run("get Session ID from cookies", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "some-url", nil)
		req.AddCookie(&http.Cookie{
			Name:  "some-session-name",
			Value: "some-session-id",
		})

		session, err := manager.Start(nil, req)
		require.NoError(t, err)
		require.Equal(t, "some-session-id", session.SessionID())
	})
}
