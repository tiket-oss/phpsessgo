package phpsessgo_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/imantung/phpsessgo"
	"github.com/imantung/phpsessgo/mock"
	"github.com/imantung/phpsessgo/phpencode"
	"github.com/stretchr/testify/require"
)

func TestSessionManager_Start_GenerateSessionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sidCreator := mock.NewMockSessionIDCreator(ctrl)
	sidCreator.EXPECT().CreateSID().Return("random-hash")

	handler := mock.NewMockSessionHandler(ctrl)

	manager := phpsessgo.SessionManager{
		SessionName: "some-session-name",
		SIDCreator:  sidCreator,
		Handler:     handler,
	}

	req, _ := http.NewRequest(http.MethodGet, "some-url", nil)
	rr := httptest.NewRecorder()

	session, err := manager.Start(rr, req)
	require.NoError(t, err)
	require.Equal(t, "random-hash", session.SessionID)
	require.Equal(t, "some-session-name=random-hash", rr.HeaderMap.Get("Set-Cookie"))
}

func TestSessionManager_Start_ExistingSessionID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sidCreator := mock.NewMockSessionIDCreator(ctrl)
	handler := mock.NewMockSessionHandler(ctrl)
	encoder := mock.NewMockSessionEncoder(ctrl)

	manager := phpsessgo.SessionManager{
		SessionName: "some-session-name",
		SIDCreator:  sidCreator,
		Handler:     handler,
		Encoder:     encoder,
	}

	req, _ := http.NewRequest(http.MethodGet, "some-url", nil)
	req.AddCookie(&http.Cookie{
		Name:  "some-session-name",
		Value: "some-session-id",
	})

	t.Run("handler failed to read", func(t *testing.T) {
		handler.EXPECT().Read("some-session-id").Return("", fmt.Errorf("some-error"))

		_, err := manager.Start(nil, req)
		require.EqualError(t, err, "some-error")
	})

	t.Run("decode success", func(t *testing.T) {
		handler.EXPECT().Read("some-session-id").Return("some-data", nil)
		encoder.EXPECT().Decode("some-data").Return(phpencode.PhpSession{}, nil)

		session, err := manager.Start(nil, req)
		require.NoError(t, err)
		require.Equal(t, "some-session-id", session.SessionID)
	})

	t.Run("decode failed", func(t *testing.T) {
		handler.EXPECT().Read("some-session-id").Return("some-data", nil)
		encoder.EXPECT().Decode("some-data").Return(nil, fmt.Errorf("some-error"))
		_, err := manager.Start(nil, req)
		require.EqualError(t, err, "some-error")
	})

}

func TestSessionManager_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sidCreator := mock.NewMockSessionIDCreator(ctrl)
	handler := mock.NewMockSessionHandler(ctrl)
	encoder := mock.NewMockSessionEncoder(ctrl)

	manager := phpsessgo.SessionManager{
		SessionName: "some-session-name",
		SIDCreator:  sidCreator,
		Handler:     handler,
		Encoder:     encoder,
	}

	session := phpsessgo.NewSession()
	session.SessionID = "some-session-id"

	t.Run("encode success", func(t *testing.T) {
		encoder.EXPECT().Encode(session.Value).Return("encoded-data", nil)
		handler.EXPECT().Write("some-session-id", "encoded-data").Return(nil)

		err := manager.Save(session)
		require.NoError(t, err)
	})

	t.Run("decode error", func(t *testing.T) {
		encoder.EXPECT().Encode(session.Value).Return("encoded-data", fmt.Errorf("some-error"))

		err := manager.Save(session)
		require.EqualError(t, err, "some-error")
	})

}
