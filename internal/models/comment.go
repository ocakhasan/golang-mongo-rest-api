package models

type Comment struct {
	PostTitle string `bson:"postTitle" json:"postTitle"`
	Comment   string `bson:"comment" json:"comment"`
	Likes     int    `bson:"likes" json:"likes"`
}
