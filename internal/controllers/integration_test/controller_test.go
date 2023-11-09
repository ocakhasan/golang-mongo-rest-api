package integrationtest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ocakhasan/mongoapi/internal/controllers"
	"github.com/ocakhasan/mongoapi/internal/repository"
	"github.com/ocakhasan/mongoapi/pkg/router"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	testDbInstance *mongo.Database
)

func TestMain(m *testing.M) {
	log.Println("setup is running")
	testDB := SetupTestDatabase()
	testDbInstance = testDB.DbInstance
	populateDB()
	exitVal := m.Run()
	log.Println("teardown is running")
	_ = testDB.container.Terminate(context.Background())
	os.Exit(exitVal)
}

func InitializeTestRouter() *echo.Echo {
	postgreRepo := repository.New(testDbInstance)

	userController := controllers.New(postgreRepo)

	return router.Initialize(userController)
}

func TestGetPostsWithComments(t *testing.T) {
	apitest.New().
		Handler(InitializeTestRouter()).
		Get("/api/posts").
		Header("content-type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len(`$.posts`, 3)).
		BodyFromFile("responses/posts.json").
		End()
}

func TestGetPostsByAuthorForDostoyevski(t *testing.T) {
	userId := 1
	apitest.New().
		Handler(InitializeTestRouter()).
		Get(fmt.Sprintf("/api/author/%d/posts", userId)).
		Header("content-type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len(`$.posts`, 2)).
		BodyFromFile("responses/posts_dostoyevski.json").
		End()
}

func TestGetPostsByAuthorForMarcusAurelius(t *testing.T) {
	userId := 2
	apitest.New().
		Handler(InitializeTestRouter()).
		Get(fmt.Sprintf("/api/author/%d/posts", userId)).
		Header("content-type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len(`$.posts`, 1)).
		BodyFromFile("responses/posts_marcus.json").
		End()
}

func TestGetPosts_NonExistentAuthor(t *testing.T) {
	userId := 10
	apitest.New().
		Handler(InitializeTestRouter()).
		Get(fmt.Sprintf("/api/author/%d/posts", userId)).
		Header("content-type", "application/json").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len(`$.posts`, 0)).
		Assert(jsonpath.Equal("$.posts", []interface{}{})).
		End()
}

func TestGetPosts_BadAuthorId(t *testing.T) {
	userId := "tesla"
	apitest.New().
		Handler(InitializeTestRouter()).
		Get(fmt.Sprintf("/api/author/%s/posts", userId)).
		Header("content-type", "application/json").
		Expect(t).
		Status(http.StatusBadRequest).
		Assert(jsonpath.Equal("$.error", fmt.Sprintf("strconv.Atoi: parsing \"%s\": invalid syntax", userId))).
		End()
}
