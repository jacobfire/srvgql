package book

type BookEntity struct {
	Id        string
	Name       string
	Author string
}

type AuthorEntity struct {
	Id string
	Name string
	Book int
}

