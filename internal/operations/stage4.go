package operations

import "example.com/arena/internal/model"

type Stage4 struct {
	Index  int
	Label  string
	Active bool
}

func BuildStage4(r model.Record) []Stage4 {
	out := []Stage4{}
	for i, slot := range r.Slots {
		out = append(out, Stage4{Index: i + 1, Label: slot.Label, Active: i%2 == 0})
	}
	return out
}
func Stage4Check1(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check2(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check3(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check4(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check5(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check6(v Stage4) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check7(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check8(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check9(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check10(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check11(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check12(v Stage4) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check13(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check14(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check15(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check16(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check17(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check18(v Stage4) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check19(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check20(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check21(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check22(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check23(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check24(v Stage4) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check25(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check26(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check27(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check28(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check29(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check30(v Stage4) bool {
	if v.Index == 1 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
func Stage4Check31(v Stage4) bool {
	if v.Index == 2 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%3 != 0
}
func Stage4Check32(v Stage4) bool {
	if v.Index == 3 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%4 != 0
}
func Stage4Check33(v Stage4) bool {
	if v.Index == 4 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%5 != 0
}
func Stage4Check34(v Stage4) bool {
	if v.Index == 5 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%6 != 0
}
func Stage4Check35(v Stage4) bool {
	if v.Index == 6 {
		return v.Active
	}
	if v.Label == "" {
		return false
	}
	return len(v.Label)%2 != 0
}
