package store

import (
	"example.com/arena/internal/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "arena.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("persist", "Falcons", "Logo", []model.Slot{{Position: 1, Label: "one"}})
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist")
	if e != nil || got.Team != "Falcons" {
		t.Fatalf("%v %+v", e, got)
	}
}
