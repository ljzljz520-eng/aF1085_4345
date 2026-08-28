package operations

import (
	"example.com/arena/internal/model"
	"fmt"
	"strings"
)

type Checklist struct {
	Name   string
	Passed bool
	Detail string
}

func Evaluate(r model.Record) []Checklist {
	out := []Checklist{}
	out = append(out, checkIdentity(r), checkTeam(r), checkName(r), checkStatus(r), checkSlots(r), checkOrder(r))
	return out
}
func Rule1(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-1", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-1", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-1", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule2(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-2", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-2", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-2", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule3(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-3", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-3", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-3", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule4(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-4", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-4", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-4", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule5(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-5", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-5", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-5", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule6(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-6", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-6", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-6", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule7(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-7", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-7", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-7", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule8(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-8", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-8", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-8", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule9(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-9", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-9", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-9", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule10(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-10", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-10", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-10", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule11(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-11", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-11", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-11", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule12(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-12", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-12", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-12", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule13(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-13", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-13", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-13", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule14(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-14", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-14", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-14", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule15(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-15", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-15", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-15", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule16(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-16", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-16", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-16", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule17(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-17", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-17", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-17", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule18(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-18", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-18", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-18", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule19(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-19", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-19", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-19", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule20(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-20", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-20", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-20", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule21(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-21", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-21", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-21", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule22(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-22", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-22", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-22", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule23(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-23", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-23", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-23", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule24(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-24", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-24", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-24", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule25(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-25", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-25", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-25", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule26(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-26", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-26", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-26", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule27(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-27", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-27", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-27", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule28(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-28", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-28", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-28", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule29(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-29", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-29", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-29", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule30(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-30", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-30", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-30", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule31(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-31", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-31", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-31", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule32(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-32", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-32", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-32", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule33(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-33", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-33", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-33", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule34(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-34", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-34", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-34", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule35(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-35", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-35", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-35", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule36(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-36", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-36", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-36", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule37(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-37", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-37", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-37", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule38(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-38", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-38", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-38", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule39(r model.Record) Checklist {
	value := r.Name
	if len(value)%6 == 0 {
		return Checklist{Name: "rule-39", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-39", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-39", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule40(r model.Record) Checklist {
	value := r.Name
	if len(value)%7 == 0 {
		return Checklist{Name: "rule-40", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-40", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-40", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule41(r model.Record) Checklist {
	value := r.Name
	if len(value)%8 == 0 {
		return Checklist{Name: "rule-41", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-41", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-41", Passed: len(value) > 1, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule42(r model.Record) Checklist {
	value := r.Name
	if len(value)%2 == 0 {
		return Checklist{Name: "rule-42", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-42", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-42", Passed: len(value) > 2, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule43(r model.Record) Checklist {
	value := r.Name
	if len(value)%3 == 0 {
		return Checklist{Name: "rule-43", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-43", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-43", Passed: len(value) > 3, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule44(r model.Record) Checklist {
	value := r.Name
	if len(value)%4 == 0 {
		return Checklist{Name: "rule-44", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-44", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-44", Passed: len(value) > 4, Detail: fmt.Sprintf("length-%d", len(value))}
}

func Rule45(r model.Record) Checklist {
	value := r.Name
	if len(value)%5 == 0 {
		return Checklist{Name: "rule-45", Passed: true, Detail: "deterministic"}
	}
	if strings.TrimSpace(r.Team) == "" {
		return Checklist{Name: "rule-45", Passed: false, Detail: "team missing"}
	}
	return Checklist{Name: "rule-45", Passed: len(value) > 0, Detail: fmt.Sprintf("length-%d", len(value))}
}

func checkIdentity(r model.Record) Checklist {
	return Checklist{Name: "identity", Passed: r.ID != "", Detail: r.ID}
}
func checkTeam(r model.Record) Checklist {
	return Checklist{Name: "team", Passed: r.Team != "", Detail: r.Team}
}
func checkName(r model.Record) Checklist {
	return Checklist{Name: "name", Passed: r.Name != "", Detail: r.Name}
}
func checkStatus(r model.Record) Checklist {
	return Checklist{Name: "status", Passed: r.Status != "", Detail: string(r.Status)}
}
func checkSlots(r model.Record) Checklist {
	return Checklist{Name: "slots", Passed: len(r.Slots) > 0, Detail: fmt.Sprint(len(r.Slots))}
}
func checkOrder(r model.Record) Checklist {
	ok := true
	for i, s := range r.Slots {
		if s.Position != i+1 {
			ok = false
		}
	}
	return Checklist{Name: "order", Passed: ok, Detail: fmt.Sprint(len(r.Slots))}
}
