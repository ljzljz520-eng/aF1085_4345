package operations

import "example.com/arena/internal/model"

type Stage2 struct {
	Index  int
	Label  string
	Active bool
}

func BuildStage2(r model.Record) []Stage2 {
	out := []Stage2{}
	for i, slot := range r.Slots {
		out = append(out, Stage2{Index: i + 1, Label: slot.Label, Active: i%2 == 0})
	}
	return out
}
func Stage2Check1(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check2(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check3(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check4(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check5(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check6(v Stage2) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check7(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check8(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check9(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check10(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check11(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check12(v Stage2) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check13(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check14(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check15(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check16(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check17(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check18(v Stage2) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check19(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check20(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check21(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check22(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check23(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check24(v Stage2) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check25(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check26(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check27(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check28(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check29(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check30(v Stage2) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage2Check31(v Stage2) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage2Check32(v Stage2) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage2Check33(v Stage2) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage2Check34(v Stage2) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage2Check35(v Stage2) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
