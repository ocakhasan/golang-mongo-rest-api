package integrationtest

import (
	"context"
	"log"

	"github.com/ocakhasan/mongoapi/internal/models"
)

func populateDB() {
	var posts = []interface{}{
		models.Post{
			Title: "Crime and Punishment",
			Author: models.Author{
				Name: "Dostoyevski",
				Id:   1,
			},
			Likes: 12,
		},
		models.Post{
			Title: "Notes From The Underground",
			Author: models.Author{
				Name: "Dostoyevski",
				Id:   1,
			},
			Likes: 100,
		},
		models.Post{
			Title: "Meditations",
			Author: models.Author{
				Name: "Marcus Aurelius",
				Id:   2,
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

	_, err := testDbInstance.Collection("posts").InsertMany(context.Background(), posts)
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := testDbInstance.Collection("comments").InsertMany(context.Background(), comments)
	if err2 != nil {
		log.Fatal(err2)
	}
}
