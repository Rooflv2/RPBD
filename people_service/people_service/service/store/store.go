package store

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// make migration //вопрос 3

	m, err := migrate.New("file://migrations",
		connString)
	if err != nil {
		fmt.Println(err)
	}

	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("migration is done")

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	var PeopleList []People
	rows, err := s.conn.Query(context.Background(), "select id, name from people")
	if err != nil {
		return nil, fmt.Errorf("list people: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var man People

		if err = rows.Scan(&man.ID, &man.Name); err != nil {
			fmt.Println("err is ", err)
		}
		PeopleList = append(PeopleList, man)
	}

	return PeopleList, err
}

func (s *Store) GetPeopleByID(id int) (People, error) {
	var name string
	err := s.conn.QueryRow(context.Background(), "SELECT name FROM people WHERE id=$1", id).Scan(&name)
	if err != nil {
		fmt.Print(err)
	}
	return People{ID: id, Name: name}, err
}
