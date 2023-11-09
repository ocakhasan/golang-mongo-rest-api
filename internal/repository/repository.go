package repository

import (
	"context"

	"github.com/ocakhasan/mongoapi/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "users"
)

type Repository interface {
	GetPostsWithComments(ctx context.Context, filter PostFilter) ([]models.PostWithComments, error)
}

func New(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

type mongoRepository struct {
	db *mongo.Database
}

type PostFilter struct {
	AuthorId *int
}

func (m *mongoRepository) GetPostsWithComments(ctx context.Context, filter PostFilter) ([]models.PostWithComments, error) {
	pipeline := mongo.Pipeline{}

	if filter.AuthorId != nil {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"author.id", *filter.AuthorId}}}})
	}

	lookupStage := bson.D{
		{"$lookup",
			bson.D{
				{"from", "comments"},
				{"localField", "title"},
				{"foreignField", "postTitle"},
				{"as", "comments"},
			},
		},
	}

	pipeline = append(pipeline, lookupStage)

	cur, err := m.db.Collection("posts").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var res = make([]models.PostWithComments, 0)
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
