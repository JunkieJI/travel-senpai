package store

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/JunkieJI/travel-senpai/config"

	_ "github.com/go-sql-driver/mysql" // for sqlx
	"github.com/jmoiron/sqlx"
)

// Store interface
type Store interface {
	Ping() bool
	GetDestinationByID(context context.Context, destinationID int64) (*Destination, error)
	GetDestinations(context context.Context) (*[]Destination, error)
	CreateDestination(ctx context.Context, dest *Destination) (*Destination, error)
	UpdateDestination(ctx context.Context, dest *Destination) (*Destination, error)
	DeleteDestination(ctx context.Context, id int64) error
}

type store struct {
	db *sqlx.DB
}

// NewStore creates a sqlx connection
func NewStore(dsn config.DSN) Store {
	if dsn.Driver == "" {
		config.Init()
		dsn = config.GetDSN()
	}
	s := &store{}
	s.connect(dsn)

	return s
}

func (s *store) connect(dsn config.DSN) error {
	log.Println("Connecting...")
	var err error
	s.db, err = sqlx.Connect(dsn.Driver, dsn.ConnectionString())
	if err != nil {
		return err
	}

	return nil
}

func (s *store) Ping() bool {
	return s.db.Ping() == nil
}
