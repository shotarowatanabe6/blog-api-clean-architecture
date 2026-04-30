package models

type Article struct {
	ID        int
	UserID    int
	Title     string
	Body      string
	Published bool
	CreatedAt float64 // UnixTime, mili second
	UpdatedAt float64 // UnixTime, mili second
}
