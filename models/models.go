package models

type Book struct {
	Book_id     int    `db:"BookId" json:"bookid"`
	AuthorID    int    `db:"AuthorID" json:"AuthorID"`
	PublisherID int    `db:"PublisherID" json:"PublisherID"`
	Book_Name   string `db:"BookName" json:"BookName"`
	Book_Url    string `db:"BookUrl" json:"BookURL"`
	IsDeleted   int    `db:"IsDeleted"`
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

type Inventory struct {
	BookId    int    `db:"BookId" json:"BookId"`
	NoOfBooks int    `db:"NoOfBooks" json:"NoOfBooks"`
	AddedBy   string `db:"AddedBy" json:"AddedBy"`
	IsDeleted int    `db:"IsDeleted"`
}
