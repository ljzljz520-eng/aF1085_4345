package model

import "testing"

func TestRecordValidation(t *testing.T) {
	r := NewRecord("x", "Falcons", "Logo", []Slot{{Position: 1, Label: "one"}})
	if e := r.Validate(); e != nil {
		t.Fatal(e)
	}
	if !LabelsUnique(r.Slots) {
		t.Fatal("labels")
	}
}
