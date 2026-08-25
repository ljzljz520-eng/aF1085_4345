package operations

import "example.com/arena/internal/model"

type Stage3 struct {
	Index  int
	Label  string
	Active bool
}

func BuildStage3(r model.Record) []Stage3 {
	out := []Stage3{}
	for i, slot := range r.Slots {
		out = append(out, Stage3{Index: i + 1, Label: slot.Label, Active: i%2 == 0})
	}
	return out
}
func Stage3Check1(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check2(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check3(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check4(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check5(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check6(v Stage3) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check7(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check8(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check9(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check10(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check11(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check12(v Stage3) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check13(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check14(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check15(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check16(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check17(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check18(v Stage3) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check19(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check20(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check21(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check22(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check23(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check24(v Stage3) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check25(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check26(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check27(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check28(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check29(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check30(v Stage3) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage3Check31(v Stage3) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage3Check32(v Stage3) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage3Check33(v Stage3) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage3Check34(v Stage3) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage3Check35(v Stage3) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
