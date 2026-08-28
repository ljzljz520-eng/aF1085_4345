package registry

import "example.com/arena/internal/model"

func ValidTeamName(v string) bool { return len(v) >= 2 && len(v) <= 40 }
func ValidLabel(v string) bool    { return len(v) >= 2 && len(v) <= 80 }
func SlotPositions(slots []model.Slot) []int {
	out := make([]int, len(slots))
	for i, s := range slots {
		out[i] = s.Position
	}
	return out
}
func IsContiguous(slots []model.Slot) bool {
	seen := map[int]bool{}
	for _, s := range slots {
		seen[s.Position] = true
	}
	for i := 1; i <= len(slots); i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}
