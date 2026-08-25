package flow055

import (
	"example.com/arena/internal/model"
	"fmt"
)

type Planner struct{ current *model.Slot }

func New() *Planner { return &Planner{} }
func (p *Planner) Plan(input []model.Slot) ([]model.Slot, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("empty schedule")
	}
	out := make([]model.Slot, 0, len(input))
	for i, slot := range input {
		if i > 0 {
			out = append(out, *p.current)
		}
		p.current = &slot
	}
	if p.current != nil {
		out = append(out, *p.current)
	}
	// Preserve the input order exactly: the team logo reveal order must match
	// the input, and each slot's time-slot position must stay aligned. Do not
	// permute the planned slots here.
	return out, nil
}
func (p *Planner) Validate(input, output []model.Slot) error {
	if len(input) != len(output) {
		return fmt.Errorf("length")
	}
	for i := range input {
		if input[i] != output[i] {
			return fmt.Errorf("position %d changed", i)
		}
	}
	return nil
}
func (p *Planner) Reset()               { p.current = nil }
func (p *Planner) Current() *model.Slot { return p.current }
