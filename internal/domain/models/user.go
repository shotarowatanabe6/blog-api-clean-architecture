package models

type User struct {
	ID        string // UUIDv4
	Name      string
	Email     string
	CreatedAt float64 // UnixTime, mili second
}
