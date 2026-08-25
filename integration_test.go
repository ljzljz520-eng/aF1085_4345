package arena_test

import (
	"example.com/arena/internal/importer"
	"example.com/arena/internal/model"
	"example.com/arena/internal/registry"
	"example.com/arena/internal/review"
	"example.com/arena/internal/store"
	"testing"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	reg := registry.New(db)
	reg.Register(model.NewRecord("r", "Falcons", "Logo", []model.Slot{{1, "open"}, {2, "final"}}))
	rv := review.New(db)
	rv.Submit("r", "operator")
	rv.Approve("r", "reviewer")
	rv.Publish("r", "reviewer")
	rv.Archive("r", "operator")
	got, _ := db.GetRecord("r")
	if got.Status != model.StatusArchived {
		t.Fatal(got.Status)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	registry.New(db).Register(model.NewRecord("r", "Falcons", "Logo", []model.Slot{{1, "open"}}))
	if n, _ := db.Count("records"); n != 1 {
		t.Fatal(n)
	}
}
func TestWorkflowImportReport(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	rep, _ := importer.Import(db, importer.Parse(importer.DeterministicBatch()))
	if rep.Imported != 2 {
		t.Fatal(rep)
	}
}
