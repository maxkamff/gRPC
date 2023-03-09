package postgres

import (
	"database/sql"
	"fmt"
	pb "grpc-todo/proto"

	_ "github.com/lib/pq"
	"github.com/lib/pq"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "maxkamff"
	PostgresPassword = "12345"
	PostgresDatabase = "store"
)

func CreateStore(store *pb.Store) (*pb.Store, error) {

	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDatabase,
	)

	db, err := sql.Open("postgres", connDB)
	if err != nil {
		return nil, err
	}

	newStore := pb.Store{}

	err = db.QueryRow(`INSERT INTO stores(name, description, is_open, address)
	VALUES($1, $2, $3, $4)
	RETURNING id, name, description, is_open, address`,
		store.Name,
		store.Description,
		store.IsOpen,
		pq.Array(store.Address)).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Description,
		&newStore.IsOpen,
		pq.Array(&newStore.Address),
	)
	if err != nil {
		return nil, err
	}

	return &newStore, nil
}
