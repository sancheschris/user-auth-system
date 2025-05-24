package integration

import (
	"context"
	"testing"

	"github.com/sancheschris/user-auth-system/internal/auth"
	"github.com/sancheschris/user-auth-system/internal/database"
	"github.com/sancheschris/user-auth-system/internal/model"
	"github.com/stretchr/testify/assert"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
)
func setupTestMongo(t *testing.T) (*auth.MongoUserRepository, *mongo.Client, func()) {
 ctx := context.Background()

    req := testcontainers.ContainerRequest{
        Image:        "mongo:latest",
        ExposedPorts: []string{"27017/tcp"},
        WaitingFor:   wait.ForListeningPort("27017/tcp"),
    }
    mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        t.Fatalf("Failed to start MongoDB container: %v", err)
    }

    host, err := mongoC.Host(ctx)
    if err != nil {
        t.Fatalf("Failed to get MongoDB host: %v", err)
    }

    port, err := mongoC.MappedPort(ctx, "27017")
    if err != nil {
        t.Fatalf("Failed to get MongoDB port: %v", err)
    }

    // Connect to the test MongoDB instance
    db, err := database.ConnectMongoDB("mongodb://"+host+":"+port.Port())
    if err != nil {
        t.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    repo := auth.NewMongoUserRepository(db)

    cleanUp := func() {
        _ = db.Collection("users").Drop(context.Background())
        _ = mongoC.Terminate(ctx)
    }

    return repo, nil, cleanUp
}

func Test_FindUser(t *testing.T) {
	repo, _, cleanUp := setupTestMongo(t)
	defer cleanUp()

	// Create a test user
	testUser := &model.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := repo.SaveUser(testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Test finding the user
	user, err := repo.FindByUsername("testuser")
	if err != nil {
		t.Fatalf("Failed to find user: %v", err)
	}

	if user.Username != testUser.Username {
		t.Errorf("Expected username %s, got %s", testUser.Username, user.Username)
	}

	assert.Equal(t, testUser.Username, user.Username)
	assert.Equal(t, testUser, user)
}