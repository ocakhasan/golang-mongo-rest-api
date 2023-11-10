package integrationtest

import (
	"context"
	"log"

	"github.com/ocakhasan/mongoapi/internal/models"
)

func populateDB() {
	var authors = []interface{}{
		models.Author{
			Id:   "654e618a60034d917aa0ae63",
			Name: "Dostoyevski",
		},
		models.Author{
			Id:   "654e619760034d917aa0ae64",
			Name: "Marcus Aurelius",
		},
	}

	var books = []interface{}{
		models.Book{
			Title: "Crime and Punishment",
			Author: models.Author{
				Name: "Dostoyevski",
				Id:   "654e618a60034d917aa0ae63",
			},
			Likes: 12,
		},
		models.Book{
			Title: "Notes From The Underground",
			Author: models.Author{
				Name: "Dostoyevski",
				Id:   "654e618a60034d917aa0ae63",
			},
			Likes: 100,
		},
		models.Book{
			Title: "Meditations",
			Author: models.Author{
				Name: "Marcus Aurelius",
				Id:   "654e619760034d917aa0ae64",
			},
			Likes: 200,
		},
	}

	var comments = []interface{}{
		models.Comment{
			Comment:   "great read",
			Likes:     3,
			PostTitle: "Crime and Punishment",
		},
		models.Comment{
			Comment:   "good info",
			Likes:     0,
			PostTitle: "Notes From The Underground",
		},
		models.Comment{
			Comment:   "I liked this post",
			Likes:     12,
			PostTitle: "Notes From The Underground",
		},
		models.Comment{
			Comment:   "very nice book",
			Likes:     8,
			PostTitle: "Meditations",
		},
	}

	_, err := testDbInstance.Collection("books").InsertMany(context.Background(), books)
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := testDbInstance.Collection("comments").InsertMany(context.Background(), comments)
	if err2 != nil {
		log.Fatal(err2)
	}

	_, err3 := testDbInstance.Collection("authors").InsertMany(context.Background(), authors)
	if err3 != nil {
		log.Fatal(err3)
	}
}
