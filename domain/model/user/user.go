package user

type User struct {
	UserId       *Id
	Password     *Password
	Notification string
	WatchList    []string
}
