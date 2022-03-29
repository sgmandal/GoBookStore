package models

type Book struct {
	Book_id     int    `db:"BookId" json:"BookID"`
	AuthorID    int    `db:"AuthorID" json:"AuthorID"`
	PublisherID int    `db:"PublisherID" json:"PublisherID"`
	Book_Name   string `db:"BookName" json:"BookName"`
	Book_Url    string `db:"BookUrl" json:"BookURL"`
}

type Author struct {
	AuthorId   int    `db:"AuthorID" json:"AuthorID"`
	AuthorName string `db:"AuthorName" json:"AuthorName"`
	AuthorAge  int    `db:"AuthorAge" json:"AuthorAge"`
}

type Publisher struct {
	PublisherID      int    `db:"PublisherID" json:"PublisherID"`
	PublisherName    string `db:"PublisherName" json:"PublisherName"`
	PublisherAddress string `db:"PublisherAddress" json:"PublisherAddress"`
}
