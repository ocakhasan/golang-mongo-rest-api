package integrationtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ocakhasan/mongoapi/pkg/database"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestDatabase struct {
	DbInstance *mongo.Database
	DbAddress  string
	container  testcontainers.Container
}

func SetupTestDatabase() *TestDatabase {
	// setup db container
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	container, dbInstance, dbAddr, err := createMongoContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup test", err)
	}

	return &TestDatabase{
		container:  container,
		DbInstance: dbInstance,
		DbAddress:  dbAddr,
	}
}

func (tdb *TestDatabase) TearDown() {
	_ = tdb.container.Terminate(context.Background())
}

func createMongoContainer(ctx context.Context) (testcontainers.Container, *mongo.Database, string, error) {

	var env = map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": "root",
		"MONGO_INITDB_ROOT_PASSWORD": "pass",
		"MONGO_INITDB_DATABASE":      "testdb",
	}
	var port = "27017/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{port},
			Env:          env,
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}

	log.Println("mongo container ready and running at port: ", p.Port())

	time.Sleep(time.Second)

	uri := fmt.Sprintf("mongodb://root:pass@localhost:%s", p.Port())
	db, err := database.NewMongoDatabase(uri)
	if err != nil {
		return container, db, uri, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, uri, nil
}
