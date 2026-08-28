package operations

import "example.com/arena/internal/model"

type Stage1 struct {
	Index  int
	Label  string
	Active bool
}

func BuildStage1(r model.Record) []Stage1 {
	out := []Stage1{}
	for i, slot := range r.Slots {
		out = append(out, Stage1{Index: i + 1, Label: slot.Label, Active: i%2 == 0})
	}
	return out
}
func Stage1Check1(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check2(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check3(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check4(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check5(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check6(v Stage1) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check7(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check8(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check9(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check10(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check11(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check12(v Stage1) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check13(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check14(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check15(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check16(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check17(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check18(v Stage1) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check19(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check20(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check21(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check22(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check23(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check24(v Stage1) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check25(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check26(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check27(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check28(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check29(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check30(v Stage1) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage1Check31(v Stage1) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage1Check32(v Stage1) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage1Check33(v Stage1) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage1Check34(v Stage1) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage1Check35(v Stage1) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
