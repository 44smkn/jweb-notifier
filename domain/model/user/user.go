package user

type User struct {
	UserId       Id
	Password     string
	Notification string
	WatchList    []string
}
