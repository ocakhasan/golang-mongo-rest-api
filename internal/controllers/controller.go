package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ocakhasan/mongoapi/internal/repository"
)

type PostsController struct {
	repo repository.Repository
}

func New(repo repository.Repository) *PostsController {
	return &PostsController{repo: repo}
}

func (u PostsController) GetPostsWithComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := u.repo.GetPostsWithComments(c.Request().Context(), repository.PostFilter{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"posts": posts,
		})
	}
}

func (u PostsController) GetAuthorPostsWithComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorId := c.Param("id")

		intAuthorId, err := strconv.Atoi(authorId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		posts, err := u.repo.GetPostsWithComments(c.Request().Context(), repository.PostFilter{
			AuthorId: &intAuthorId,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"posts": posts,
		})
	}
}
