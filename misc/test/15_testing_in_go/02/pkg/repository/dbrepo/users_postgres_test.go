//go:build integration

package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/pkg/data"
	"myapp/pkg/repository"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=utc connect_timeout=5"
)

var (
	resource *dockertest.Resource
	pool     *dockertest.Pool
	testDB   *sql.DB
	testRepo repository.DatabaseRepo
)

func TestMain(m *testing.M) {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Couldn't connect to docker. Is it running? %s", err)
	}

	pool = p

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Couldn't start a resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Couldn't connect to database: %s", err)
	}

	err = createTables()
	if err != nil {
		log.Fatalf("Error creating sql tables: %s", err)
	}

	testRepo = &PostgresDBRepo{DB: testDB}

	ret := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Couldn't purge resource: %s", err)
	}

	os.Exit(ret)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Errorf("Can't ping database: %s", err)
	}
}

func TestPostgresDBRepoInsertUser(t *testing.T) {
	testUser := data.User{
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("Insert user returned an error: %s", err)
	}

	if id != 1 {
		t.Errorf("Insert user returned wrong id; expected 1, but got %d", id)
	}
}

func TestPostgresDBRepoAllUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Fatalf("AllUsers returned an error: %s", err)
	}

	if len(users) != 1 {
		t.Fatalf("AllUsers returned wrong size; expected 1, got: %d", len(users))
	}

	testUser := data.User{
		FirstName: "Jack",
		LastName:  "Smith",
		Email:     "jack@smith.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Fatalf("AllUsers returned an error: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("AllUsers returned wrong size after insert; expected 2, got: %d", len(users))
	}
}

func TestPostgresDBRepoGetUser(t *testing.T) {
	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("Error getting user by id: %s", err)
	}

	if user.Email != "admin@example.com" {
		t.Errorf("Wrong email returned by GetUser; expected admin@example.com, got %v", user.Email)
	}

	_, err = testRepo.GetUser(3)
	if err == nil {
		t.Errorf("Error is missed getting user by id 3")
	}
}
func TestPostgresDBRepoGetUserByEmail(t *testing.T) {
	user, err := testRepo.GetUserByEmail("jack@smith.com")
	if err != nil {
		t.Errorf("Error getting user by email: %s", err)
	}

	if user.ID != 2 {
		t.Errorf("Wrong id returned by GetUserByEmail; expected 2, got %d", user.ID)
	}
}

func TestPostgresDBRepoUpdateUser(t *testing.T) {
	user, _ := testRepo.GetUser(2)
	if user.FirstName != "Jack" || user.Email != "jack@smith.com" {
		t.Errorf("Expected second user, but got %s %s", user.FirstName, user.Email)
	}

	user.FirstName = "Jane"
	user.Email = "jane@smith.com"

	err := testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("Error updating user 2: %s", err)
	}

	user, _ = testRepo.GetUser(2)
	if user.FirstName != "Jane" || user.Email != "jane@smith.com" {
		t.Errorf("Expected updated record, but got %s %s", user.FirstName, user.Email)
	}
}

func TestPostgresDBRepoDeleteUser(t *testing.T) {
	user, _ := testRepo.GetUser(2)
	if user.FirstName != "Jane" || user.Email != "jane@smith.com" {
		t.Errorf("Expected second user, but got %s %s", user.FirstName, user.Email)
	}

	err := testRepo.DeleteUser(2)
	if err != nil {
		t.Errorf("Error deleting user with id 2: %s", err)
	}

	_, err = testRepo.GetUser(2)
	if err == nil {
		t.Errorf("Retreived use id 2, who should have been deleted")
	}
}

func TestPostgresDBRepoResetPassword(t *testing.T) {
	err := testRepo.ResetPassword(1, "password")
	if err != nil {
		t.Errorf("Error resetting users's password: %s", err)
	}

	user, _ := testRepo.GetUser(1)
	matched, err := user.PasswordMatches("password")
	if err != nil {
		t.Errorf("Error when trying to match passwords")
	}
	if !matched {
		t.Errorf("Error - passwords don't match")
	}
}

func TestPostgresDBRepoInsertUserImage(t *testing.T) {
	var image data.UserImage
	image.UserID = 1
	image.FileName = "test.jpg"
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()

	newID, err := testRepo.InsertUserImage(image)
	if err != nil {
		t.Errorf("Error inserting an image: %s", err)
	}

	if newID != 1 {
		t.Errorf("Wrong image id; want 1, got %d", newID)
	}

	image.UserID = 100
	_, err = testRepo.InsertUserImage(image)
	if err == nil {
		t.Error("Error inserting an image with non-existing user id")
	}
}
