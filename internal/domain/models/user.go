package models

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt float64 // UnixTime, mili second
}
