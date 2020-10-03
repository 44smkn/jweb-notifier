package update

type UpdatesRepository interface {
	Fetch() *Updates
}
