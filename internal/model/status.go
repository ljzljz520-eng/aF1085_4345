package model

func AllowedTransition(from, to Status) bool {
	switch from {
	case StatusDraft:
		return to == StatusPending
	case StatusPending:
		return to == StatusApproved || to == StatusDraft
	case StatusApproved:
		return to == StatusPublished || to == StatusDraft
	case StatusPublished:
		return to == StatusArchived
	case StatusArchived:
		return false
	}
	return false
}
func StatusName(s Status) string {
	if s == "" {
		return "unknown"
	}
	return string(s)
}
func NormalizeSlots(slots []Slot) []Slot {
	out := append([]Slot(nil), slots...)
	for i := range out {
		if out[i].Position == 0 {
			out[i].Position = i + 1
		}
	}
	return out
}
