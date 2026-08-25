package store

import (
	"example.com/arena/internal/model"
	"sort"
)

func (s *Store) FindByTeam(team string) ([]model.Record, error) {
	all, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := make([]model.Record, 0)
	for _, r := range all {
		if r.Team == team {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}
func (s *Store) ReplaceRecord(r model.Record) error { return s.SaveRecord(r) }
func (s *Store) HasRecord(id string) bool           { _, e := s.GetRecord(id); return e == nil }
func (s *Store) Snapshot() ([]model.Record, error)  { return s.ListRecords() }
