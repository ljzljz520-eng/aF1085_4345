package review

func ActorAllowed(actor string) bool {
	return actor == "operator" || actor == "reviewer" || actor == "admin"
}
func RequiresSecondReview(detail string) bool { return len(detail) > 120 }
func ArchiveReason(ok bool) string {
	if ok {
		return "publication complete"
	}
	return "manual archive"
}
