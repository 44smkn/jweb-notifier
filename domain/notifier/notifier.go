package entity

type Notifier interface {
	Notify(string) error
}
