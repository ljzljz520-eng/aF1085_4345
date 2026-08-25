package flow055

import "example.com/arena/internal/model"

func Labels(slots []model.Slot) []string {
	out := make([]string, len(slots))
	for i, s := range slots {
		out[i] = s.Label
	}
	return out
}
func Positions(slots []model.Slot) []int {
	out := make([]int, len(slots))
	for i, s := range slots {
		out[i] = s.Position
	}
	return out
}
func Stable(slots []model.Slot) bool {
	for i := 1; i < len(slots); i++ {
		if slots[i].Position <= slots[i-1].Position {
			return false
		}
	}
	return true
}
func Copy(slots []model.Slot) []model.Slot { return append([]model.Slot(nil), slots...) }
