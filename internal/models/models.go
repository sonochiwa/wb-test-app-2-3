package models

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID       int
	Name     string
	Mark     int
	Category string
}
