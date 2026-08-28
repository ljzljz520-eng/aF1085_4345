package importer

import (
	"example.com/arena/internal/model"
	"example.com/arena/internal/store"
	"fmt"
	"strings"
)

type Item struct {
	ID, Team, Name string
	Labels         []string
}
type Report struct {
	Imported int
	Rejected int
	Messages []string
}

func Parse(lines []string) []Item {
	out := []Item{}
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) >= 4 {
			out = append(out, Item{ID: parts[0], Team: parts[1], Name: parts[2], Labels: strings.Split(parts[3], ",")})
		}
	}
	return out
}
func Validate(items []Item) []string {
	errs := []string{}
	for _, it := range items {
		if it.ID == "" || it.Team == "" || len(it.Labels) == 0 {
			errs = append(errs, it.ID+":invalid")
		}
	}
	return errs
}
func Import(db *store.Store, items []Item) (Report, error) {
	rep := Report{}
	for _, it := range items {
		if len(it.Labels) == 0 {
			rep.Rejected++
			continue
		}
		slots := make([]model.Slot, len(it.Labels))
		for i, l := range it.Labels {
			slots[i] = model.Slot{Position: i + 1, Label: l}
		}
		r := model.NewRecord(it.ID, it.Team, it.Name, slots)
		if e := r.Validate(); e != nil {
			rep.Rejected++
			rep.Messages = append(rep.Messages, e.Error())
			continue
		}
		if e := db.SaveRecord(r); e != nil {
			return rep, e
		}
		rep.Imported++
	}
	return rep, nil
}
func Summary(rep Report) string {
	return fmt.Sprintf("imported=%d rejected=%d", rep.Imported, rep.Rejected)
}
