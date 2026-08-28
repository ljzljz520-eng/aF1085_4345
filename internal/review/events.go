package review

import "example.com/arena/internal/model"

func Event(id, action, actor string) model.AuditEvent {
	return model.AuditEvent{ID: id + ":" + action, RecordID: id, Action: action, Actor: actor, Detail: "workflow transition"}
}
func IsTerminal(s model.Status) bool { return s == model.StatusArchived }
func Next(s model.Status) model.Status {
	switch s {
	case model.StatusDraft:
		return model.StatusPending
	case model.StatusPending:
		return model.StatusApproved
	case model.StatusApproved:
		return model.StatusPublished
	case model.StatusPublished:
		return model.StatusArchived
	}
	return s
}
