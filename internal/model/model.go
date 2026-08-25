package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPending   Status = "pending"
	StatusApproved  Status = "approved"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type Slot struct {
	Position int    `json:"position"`
	Label    string `json:"label"`
}
type Record struct {
	ID        string `json:"id"`
	Team      string `json:"team"`
	Name      string `json:"name"`
	Status    Status `json:"status"`
	Slots     []Slot `json:"slots"`
	CreatedAt string `json:"created_at"`
}
type AuditEvent struct {
	ID       string `json:"id"`
	RecordID string `json:"record_id"`
	Action   string `json:"action"`
	Actor    string `json:"actor"`
	Detail   string `json:"detail"`
}
type Workflow struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	State     string   `json:"state"`
	RecordIDs []string `json:"record_ids"`
}
type Attachment struct {
	ID       string `json:"id"`
	RecordID string `json:"record_id"`
	Name     string `json:"name"`
	Digest   string `json:"digest"`
}

func NewRecord(id, team, name string, slots []Slot) Record {
	return Record{ID: id, Team: strings.TrimSpace(team), Name: strings.TrimSpace(name), Status: StatusDraft, Slots: append([]Slot(nil), slots...), CreatedAt: "2026-01-01T00:00:00Z"}
}
func (r Record) Validate() error {
	if r.ID == "" || r.Team == "" || r.Name == "" {
		return fmt.Errorf("record identity required")
	}
	if len(r.Slots) == 0 {
		return fmt.Errorf("slots required")
	}
	seen := map[int]bool{}
	for _, s := range r.Slots {
		if s.Position < 1 || s.Label == "" || seen[s.Position] {
			return fmt.Errorf("invalid slot %d", s.Position)
		}
		seen[s.Position] = true
	}
	return nil
}
func (r Record) Clone() Record        { r.Slots = append([]Slot(nil), r.Slots...); return r }
func Encode(v any) ([]byte, error)    { return json.Marshal(v) }
func Decode(data []byte, v any) error { return json.Unmarshal(data, v) }
func (r Record) IsMutable() bool {
	return r.Status == StatusDraft || r.Status == StatusPending || r.Status == StatusApproved
}
func (r Record) OrderedLabels() []string {
	out := make([]string, len(r.Slots))
	for i, s := range r.Slots {
		out[i] = s.Label
	}
	return out
}
