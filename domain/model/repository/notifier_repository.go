package repository

type Notifier interface {
	Notify(text string)
}
