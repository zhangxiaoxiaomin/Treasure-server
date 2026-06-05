package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"treasure-server/models"
	"treasure-server/repository"
)

// CollectionHandler handles HTTP requests for collections
type CollectionHandler struct {
	repo *repository.CollectionRepo
}

// NewCollectionHandler creates a new CollectionHandler
func NewCollectionHandler(repo *repository.CollectionRepo) *CollectionHandler {
	return &CollectionHandler{repo: repo}
}

// HandleCollections routes requests based on method
func (h *CollectionHandler) HandleCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extract ID from path: /api/collections/{id}
	path := strings.TrimPrefix(r.URL.Path, "/api/collections")
	path = strings.TrimPrefix(path, "/")

	switch r.Method {
	case http.MethodGet:
		if path == "" {
			h.listCollections(w, r)
		} else {
			h.getCollection(w, path)
		}
	case http.MethodPost:
		if path == "" {
			h.createCollection(w, r)
		} else {
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		}
	case http.MethodPut:
		if path != "" {
			h.updateCollection(w, r, path)
		} else {
			http.Error(w, `{"error":"id is required"}`, http.StatusBadRequest)
		}
	case http.MethodDelete:
		if path != "" {
			h.deleteCollection(w, path)
		} else {
			http.Error(w, `{"error":"id is required"}`, http.StatusBadRequest)
		}
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (h *CollectionHandler) listCollections(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	keyword := r.URL.Query().Get("keyword")
	page := parseInt(r.URL.Query().Get("page"), 1)
	pageSize := parseInt(r.URL.Query().Get("pageSize"), 20)

	resp, err := h.repo.List(category, keyword, page, pageSize)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *CollectionHandler) getCollection(w http.ResponseWriter, id string) {
	collection, err := h.repo.GetByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "collection not found"})
		return
	}

	writeJSON(w, http.StatusOK, collection)
}

func (h *CollectionHandler) createCollection(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	collection, err := h.repo.Create(&req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, collection)
}

func (h *CollectionHandler) updateCollection(w http.ResponseWriter, r *http.Request, id string) {
	var req models.UpdateCollectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	collection, err := h.repo.Update(id, &req)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, collection)
}

func (h *CollectionHandler) deleteCollection(w http.ResponseWriter, id string) {
	if err := h.repo.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// parseInt parses an integer from a string, returning defaultVal on failure
func parseInt(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	var val int
	if _, err := fmt.Sscanf(s, "%d", &val); err != nil {
		return defaultVal
	}
	return val
}