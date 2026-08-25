package registry

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/store"
	"testing"
)

func TestRegisterSearchUpdate(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	r := New(db)
	x := model.NewRecord("r", "Falcons", "Logo", []model.Slot{{1, "a"}, {2, "b"}})
	if e := r.Register(x); e != nil {
		t.Fatal(e)
	}
	if e := r.UpdateSlots("r", []model.Slot{{1, "b"}, {2, "a"}}); e != nil {
		t.Fatal(e)
	}
	if e := r.ValidateOrder("r", []string{"b", "a"}); e != nil {
		t.Fatal(e)
	}
}
