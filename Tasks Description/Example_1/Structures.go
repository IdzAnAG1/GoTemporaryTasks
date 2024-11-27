package Example_1

import "time"

type Book struct {
	ID          string
	Title       string
	Author      string
	PublishYear int
	isAvailable bool
	BookIssued  *Client
	IssuedDate  time.Time
	ReturnDate  time.Time
}
type Client struct {
	ID       string
	Name     string
	Email    string
	Books    *[]Book
	JoinDate time.Time
}
type Library struct {
	Books              map[string]*Book
	Clients            map[string]*Client
	MaxBooksPerMembers int
}
