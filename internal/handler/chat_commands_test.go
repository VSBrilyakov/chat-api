package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/VSBrilyakov/chat-api/internal/service"
	mock_service "github.com/VSBrilyakov/chat-api/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_AddChat(t *testing.T) {
	type mockBehavior func(s *mock_service.MockChatCommands, chat *chatApp.Chat)
	testTable := []struct {
		name               string
		inputBody          string
		mockBehavior       mockBehavior
		inputChat          *chatApp.Chat
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"title": "test chat"}`,
			mockBehavior: func(s *mock_service.MockChatCommands, chat *chatApp.Chat) {
				s.EXPECT().AddChat(chat).Return(nil)
			},
			inputChat: &chatApp.Chat{
				Title: "test chat",
			},
			expectedStatusCode: 200,
		},
		{
			name:      "Space title",
			inputBody: `{"title":" "}`,
			mockBehavior: func(s *mock_service.MockChatCommands, chat *chatApp.Chat) {
				s.EXPECT().AddChat(chat).Return(errors.New("service failure"))
			},
			inputChat:          &chatApp.Chat{Title: " "},
			expectedStatusCode: 400,
		},
	}

	var needInitRoutes bool = true
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			chat := mock_service.NewMockChatCommands(c)
			testCase.mockBehavior(chat, testCase.inputChat)

			services := &service.Service{ChatCommands: chat}
			handler := NewHTTPHandler(services)
			if needInitRoutes {
				handler.InitRoutes()
				needInitRoutes = false
			}

			w := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPost, "/chats", bytes.NewBufferString(testCase.inputBody))

			handler.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
