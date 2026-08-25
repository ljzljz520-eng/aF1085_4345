package review

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/store"
	"fmt"
)

type Service struct{ db *store.Store }

func New(db *store.Store) *Service { return &Service{db: db} }
func (s *Service) Submit(id, actor string) error {
	return s.transition(id, model.StatusPending, actor, "submitted")
}
func (s *Service) Approve(id, actor string) error {
	return s.transition(id, model.StatusApproved, actor, "approved")
}
func (s *Service) Publish(id, actor string) error {
	return s.transition(id, model.StatusPublished, actor, "published")
}
func (s *Service) Archive(id, actor string) error {
	return s.transition(id, model.StatusArchived, actor, "archived")
}
func (s *Service) transition(id string, to model.Status, actor, action string) error {
	r, e := s.db.GetRecord(id)
	if e != nil {
		return e
	}
	if !model.AllowedTransition(r.Status, to) {
		return fmt.Errorf("transition %s to %s denied", r.Status, to)
	}
	r.Status = to
	if e = s.db.SaveRecord(r); e != nil {
		return e
	}
	return s.db.SaveAudit(model.AuditEvent{ID: id + "-" + action, RecordID: id, Action: action, Actor: actor, Detail: model.StatusName(to)})
}
func (s *Service) AuditCount() (int, error) { return s.db.Count("audits") }
