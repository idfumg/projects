package db

import (
	"fmt"
	"log"
	"myapp/internal/rocket"
	"os"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore() (*Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		log.Println("Failed to run migrations")
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s Store) GetRocket(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(
		`SELECT id, type, name FROM rockets WHERE id=$1;`,
		id,
	)
	err := row.Scan(&rkt.ID, &rkt.Type, &rkt.Name)
	if err != nil {
		log.Print(err)
		return rocket.Rocket{}, err
	}
	return rkt, nil
}

func (s Store) AddRocket(r rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets (id, name, type) VALUES (:id, :name, :type)`,
		r,
	)
	if err != nil {
		return rocket.Rocket{}, err
	}
	return r, nil
}

func (s Store) DelRocket(id string) error {
	_, err := s.db.Exec(
		`DELETE FROM rockets WHERE id = $1`,
		id,
	)
	return err
}
