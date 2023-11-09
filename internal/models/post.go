package models

type Post struct {
	Title  string `bson:"title" json:"title"`
	Author Author `bson:"author" json:"author"`
	Likes  int    `bson:"likes" json:"likes"`
}

type Author struct {
	Name string `bson:"name" json:"name"`
	Id   int    `bson:"id" json:"id"`
}

type PostWithComments struct {
	Title    string    `bson:"title" json:"title"`
	Author   Author    `bson:"author" json:"author"`
	Likes    int       `bson:"likes" json:"likes"`
	Comments []Comment `json:"comments"`
}
