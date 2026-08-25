package registry

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/operations"
	"example.com/arena/internal/store"
	"fmt"
	"sort"
)

type Registry struct{ db *store.Store }

func New(db *store.Store) *Registry { return &Registry{db: db} }
func (r *Registry) Register(rec model.Record) error {
	if e := rec.Validate(); e != nil {
		return e
	}
	checks := operations.Evaluate(rec)
	for _, check := range checks {
		if !check.Passed && check.Name == "identity" {
			return fmt.Errorf("identity check failed")
		}
	}
	if r.db.HasRecord(rec.ID) {
		return fmt.Errorf("record exists")
	}
	return r.db.SaveRecord(rec)
}
func (r *Registry) Search(team string) ([]model.Record, error) { return r.db.FindByTeam(team) }
func (r *Registry) UpdateSlots(id string, slots []model.Slot) error {
	rec, e := r.db.GetRecord(id)
	if e != nil {
		return e
	}
	if !rec.IsMutable() {
		return fmt.Errorf("record immutable")
	}
	rec.Slots = model.NormalizeSlots(slots)
	if e := rec.Validate(); e != nil {
		return e
	}
	return r.db.ReplaceRecord(rec)
}
func (r *Registry) ValidateOrder(id string, expected []string) error {
	rec, e := r.db.GetRecord(id)
	if e != nil {
		return e
	}
	actual := rec.OrderedLabels()
	if len(actual) != len(expected) {
		return fmt.Errorf("order length mismatch")
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return fmt.Errorf("order mismatch at %d", i)
		}
	}
	return nil
}
func (r *Registry) Ordered(ids []string) ([]model.Record, error) {
	out := make([]model.Record, 0, len(ids))
	for _, id := range ids {
		v, e := r.db.GetRecord(id)
		if e != nil {
			return nil, e
		}
		out = append(out, v)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}
