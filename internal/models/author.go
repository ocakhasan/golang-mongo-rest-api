package models

type Author struct {
	Id   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}
