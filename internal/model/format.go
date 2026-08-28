package model

import "strings"

func CanonicalTeam(v string) string { return strings.ToUpper(strings.TrimSpace(v)) }
func CanonicalName(v string) string { return strings.TrimSpace(strings.Join(strings.Fields(v), " ")) }
func LabelsUnique(slots []Slot) bool {
	seen := map[string]bool{}
	for _, s := range slots {
		if seen[s.Label] {
			return false
		}
		seen[s.Label] = true
	}
	return true
}
