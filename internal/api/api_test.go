package api

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/registry"
	"example.com/arena/internal/store"
	"net/http/httptest"
	"testing"
)

func TestHandlerRecords(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	db.SaveRecord(model.NewRecord("r", "Falcons", "Logo", []model.Slot{{1, "a"}}))
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/records?team=Falcons", nil)
	New(registry.New(db)).ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
