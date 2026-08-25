package importer

import (
	"example.com/arena/internal/store"
	"testing"
)

func TestImportReport(t *testing.T) {
	db, _ := store.Open(t.TempDir() + "/x.db")
	defer db.Close()
	rep, e := Import(db, Parse(DeterministicBatch()))
	if e != nil || rep.Imported != 2 || !Accepted(rep) {
		t.Fatalf("%v %+v", e, rep)
	}
}
