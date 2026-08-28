package api

import "example.com/arena/internal/model"

func DecodeSlots(labels []string) []model.Slot {
	out := make([]model.Slot, len(labels))
	for i, l := range labels {
		out[i] = model.Slot{Position: i + 1, Label: l}
	}
	return out
}
func Ready(status model.Status) bool {
	return status == model.StatusApproved || status == model.StatusPublished
}
