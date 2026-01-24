package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	chat_app "github.com/VSBrilyakov/chat-api"
)

const (
	defaultLimitValue = 20
	maxLimitValue     = 200
)

func extractURLParamStr(r *http.Request, position int) (string, error) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) <= position {
		return "", fmt.Errorf("parameter not found at position %d", position)
	}

	return parts[position], nil
}

func extractURLParamInt(r *http.Request, position int) (int, error) {
	var paramStr string
	var paramInt int
	var err error

	if paramStr, err = extractURLParamStr(r, position); err != nil {
		return 0, err
	}

	if paramInt, err = strconv.Atoi(paramStr); err != nil || paramInt < 0 {
		return 0, fmt.Errorf("invalid id at position %d", position)
	}

	return paramInt, nil
}

func (h *HTTPHandler) getChatData(w http.ResponseWriter, r *http.Request) {
	var chatId int
	var err error
	if chatId, err = extractURLParamInt(r, 1); err != nil {
		responseError(w, http.StatusBadRequest, "invalid chat id", err)
		return
	}

	var limitValue int
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitValue = defaultLimitValue
	} else {
		limitValue, err = strconv.Atoi(limitStr)
		if err != nil || limitValue < 1 || limitValue > maxLimitValue {
			responseError(w, http.StatusBadRequest, "error getting data", fmt.Errorf("invalid limit value"))
			return
		}
	}

	var chatMessages *chat_app.ChatMessages
	if chatMessages, err = h.services.GetChat(chatId, limitValue); err != nil {
		responseError(w, http.StatusBadRequest, "error getting data", err)
		return
	}

	responseOK(w, chatMessages)
}

func (h *HTTPHandler) postChatData(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	switch len(pathParts) {
	case 2:
		{
			var newChat chat_app.Chat
			if err := json.NewDecoder(r.Body).Decode(&newChat); err != nil {
				responseError(w, http.StatusBadRequest, "invalid JSON", err)
				return
			}

			if err := h.validator.Struct(newChat); err != nil {
				responseError(w, http.StatusBadRequest, "validating error", err)
				return
			}

			if err := h.services.AddChat(&newChat); err != nil {
				responseError(w, http.StatusBadRequest, "error adding chat", err)
				return
			}

			responseOK(w, newChat)
		}
	case 4:
		{
			if part, err := extractURLParamStr(r, 2); err != nil || part != "messages" {
				responseError(w, http.StatusNotFound, "invalid path", err)
				return
			}

			var newMsg chat_app.Message
			var err error
			if newMsg.ChatId, err = extractURLParamInt(r, 1); err != nil {
				responseError(w, http.StatusBadRequest, "invalid chat id", err)
				return
			}

			if err = json.NewDecoder(r.Body).Decode(&newMsg); err != nil {
				responseError(w, http.StatusBadRequest, "invalid JSON", err)
				return
			}

			if err = h.validator.Struct(newMsg); err != nil {
				responseError(w, http.StatusBadRequest, "validating error", err)
				return
			}

			if err = h.services.AddMessage(&newMsg); err != nil {
				responseError(w, http.StatusBadRequest, "error adding message", err)
				return
			}

			responseOK(w, newMsg)
		}
	default:
		{
			responseError(w, http.StatusNotFound, "invalid path", nil)
			//http.Error(w, "Invalid path", http.StatusNotFound)
		}
	}
}

func (h *HTTPHandler) deleteChat(w http.ResponseWriter, r *http.Request) {
	var chatId int
	var err error
	if chatId, err = extractURLParamInt(r, 1); err != nil {
		responseError(w, http.StatusBadRequest, "invalid chat id", err)
		return
	}
	
	if err = h.services.DeleteChat(chatId); err != nil {
		responseError(w, http.StatusBadRequest, "error getting data", err)
		return
	}

	responseOK(w, nil)
}
