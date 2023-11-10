package repository

import (
	"context"

	"github.com/ocakhasan/mongoapi/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "users"
)

type Repository interface {
	GetBooksWithComments(ctx context.Context, filter PostFilter) ([]models.BookWithComments, error)
	CreateBook(ctx context.Context, book models.Book) (models.Book, error)
	GetAuthorById(ctx context.Context, id primitive.ObjectID) (models.Author, error)
}

func New(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

type mongoRepository struct {
	db *mongo.Database
}

func (m *mongoRepository) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {
	_, err := m.db.Collection("books").InsertOne(ctx, book)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (m *mongoRepository) GetAuthorById(ctx context.Context, id primitive.ObjectID) (models.Author, error) {
	var author models.Author
	if err := m.db.Collection("authors").FindOne(ctx, bson.M{"id": id}).Decode(&author); err != nil {
		return models.Author{}, nil
	}

	return author, nil
}

type PostFilter struct {
	AuthorId *string
}

func (m *mongoRepository) GetBooksWithComments(ctx context.Context, filter PostFilter) ([]models.BookWithComments, error) {
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

	cur, err := m.db.Collection("books").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var res = make([]models.BookWithComments, 0)
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
