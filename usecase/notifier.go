package usecase

type Notifier interface {
	Notify(string) error
}
