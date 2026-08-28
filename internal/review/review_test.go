package review

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/store"
	"testing"
)

func TestReviewLifecycle(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	db.SaveRecord(model.NewRecord("r", "T", "N", []model.Slot{{1, "a"}}))
	s := New(db)
	for _, f := range []func(string, string) error{s.Submit, s.Approve, s.Publish, s.Archive} {
		if e := f("r", "operator"); e != nil {
			t.Fatal(e)
		}
	}
	r, _ := db.GetRecord("r")
	if r.Status != model.StatusArchived {
		t.Fatal(r.Status)
	}
}
