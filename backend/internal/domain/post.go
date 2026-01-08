package domain

type Post struct {
	ID         int 
	created_at int 
	Title      string `JSON:"title"`      
	Content    string `JSON:"content"`
}