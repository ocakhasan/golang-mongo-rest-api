package models

type Book struct {
	Title  string `bson:"title" json:"title"`
	Author Author `bson:"author" json:"author"`
	Likes  int    `bson:"likes" json:"likes"`
}

type BookWithComments struct {
	Title    string    `bson:"title" json:"title"`
	Author   Author    `bson:"author" json:"author"`
	Likes    int       `bson:"likes" json:"likes"`
	Comments []Comment `json:"comments"`
}
