package models

type Book struct {
	Book_id   int    `db:"BookId"`
	Book_Name string `db:"BookName"`
	Book_Url  string `db:"BookUrl"`
}
