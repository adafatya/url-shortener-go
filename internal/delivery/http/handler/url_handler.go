package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"url-shortener-go/internal/repository/memory"
	"url-shortener-go/internal/usecase"
)

type URLHandler struct {
	service *usecase.URLService
}

type shortenRequest struct {
	TargetURL string `json:"target_url"`
}

func NewURLHandler(urlService *usecase.URLService) *URLHandler {
	return &URLHandler{service: urlService}
}

func (h *URLHandler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/urls", h.shorten)
	mux.HandleFunc("/", h.resolve)
	return mux
}

func (h *URLHandler) shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	shortURL, err := h.service.Shorten(r.Context(), request.TargetURL)
	if errors.Is(err, usecase.ErrInvalidTargetURL) {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, "could not shorten url")
		return
	}

	h.writeJSON(w, http.StatusCreated, shortURL)
}

func (h *URLHandler) resolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	if shortCode == "" {
		h.writeError(w, http.StatusNotFound, "url not found")
		return
	}

	shortURL, err := h.service.Resolve(r.Context(), shortCode)
	if errors.Is(err, memory.ErrURLNotFound) {
		h.writeError(w, http.StatusNotFound, "url not found")
		return
	}
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, "could not resolve url")
		return
	}

	http.Redirect(w, r, shortURL.TargetURL, http.StatusFound)
}

func (h *URLHandler) writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func (h *URLHandler) writeError(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, map[string]string{"error": message})
}
