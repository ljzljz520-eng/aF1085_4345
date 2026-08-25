package flow055

import (
	"example.com/arena/internal/model"
	"testing"
)

func Test1085BusinessRegression(t *testing.T) {
	p := New()
	in := []model.Slot{{1, "opening"}, {2, "second"}, {3, "final"}}
	out, e := p.Plan(in)
	if e != nil {
		t.Fatal(e)
	}
	if e = p.Validate(in, out); e != nil {
		t.Fatal(e)
	}
	for i := range in {
		if out[i].Label != in[i].Label {
			t.Fatalf("slot %d label %s", i, out[i].Label)
		}
	}
}
func TestPlanStable(t *testing.T) {
	p := New()
	out, _ := p.Plan([]model.Slot{{1, "a"}, {2, "b"}})
	if !Stable(out) {
		t.Fatal("unstable")
	}
}
