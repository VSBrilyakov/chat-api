package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/VSBrilyakov/chat-api/internal/service"
	"github.com/go-playground/validator/v10"
)

type HTTPHandler struct {
	services  *service.Service
	validator *validator.Validate
}

func NewHTTPHandler(service *service.Service) *HTTPHandler {
	newHandler := &HTTPHandler{
		services:  service,
		validator: validator.New(),
	}

	return newHandler
}

func (h *HTTPHandler) InitRoutes() {
	http.HandleFunc("/chats", h.ServeHTTP)
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getChatData(w, r)

	case "POST":
		h.postChatData(w, r)

	case "DELETE":
		h.deleteChat(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func responseOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encoding error: %s", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

type ErrorResponse struct {
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Details   string    `json:"details"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

func responseError(w http.ResponseWriter, code int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResp := ErrorResponse{
		Error:     http.StatusText(code),
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
	}

	if err != nil {
		errorResp.Details = err.Error()
	}

	json.NewEncoder(w).Encode(errorResp)
}
