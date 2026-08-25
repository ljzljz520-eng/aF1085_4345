package importer

func Accepted(rep Report) bool { return rep.Imported > 0 && rep.Rejected == 0 }
func Merge(a, b Report) Report {
	a.Imported += b.Imported
	a.Rejected += b.Rejected
	a.Messages = append(a.Messages, b.Messages...)
	return a
}
func DeterministicBatch() []string {
	return []string{"r1|Falcons|Falcon Logo|opening,roster,final", "r2|Titans|Titan Logo|opening,highlight,final"}
}
