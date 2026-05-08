package models

type User struct {
	ID        string  // UUIDv4
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	CreatedAt float64 // UnixTime, mili second
}
