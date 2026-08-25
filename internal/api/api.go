package api

import (
	"encoding/json"
	"example.com/arena/internal/registry"
	"net/http"
)

type Handler struct{ reg *registry.Registry }

func New(reg *registry.Registry) *Handler { return &Handler{reg: reg} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/records" {
		http.NotFound(w, r)
		return
	}
	team := r.URL.Query().Get("team")
	items, e := h.reg.Search(team)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func Health(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) }
