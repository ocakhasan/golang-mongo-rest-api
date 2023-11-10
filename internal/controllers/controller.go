package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ocakhasan/mongoapi/internal/models"
	"github.com/ocakhasan/mongoapi/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsController struct {
	repo repository.Repository
}

func New(repo repository.Repository) *PostsController {
	return &PostsController{repo: repo}
}

type CreateBookRequest struct {
	AuthorId string `json:"author_id"`
	BookName string `json:"book_name"`
}

func (u PostsController) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(CreateBookRequest)

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"err": err.Error(),
			})
		}

		objId, err := primitive.ObjectIDFromHex(req.AuthorId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"err": err.Error(),
			})
		}

		author, err := u.repo.GetAuthorById(c.Request().Context(), objId.Hex())
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"err": "author does not exist",
				})
			}
		}

		createdBook, err := u.repo.CreateBook(c.Request().Context(), models.Book{
			Title:  req.BookName,
			Author: *author,
			Likes:  0,
		})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"err": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"book": createdBook,
		})
	}
}

func (u PostsController) GetBooksWithComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := u.repo.GetBooksWithComments(c.Request().Context(), repository.PostFilter{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"books": posts,
		})
	}
}

func (u PostsController) GetAuthorBooksWithComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorId := c.Param("id")

		// validate the id
		_, err := primitive.ObjectIDFromHex(authorId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		posts, err := u.repo.GetBooksWithComments(c.Request().Context(), repository.PostFilter{
			AuthorId: &authorId,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"books": posts,
		})
	}
}
