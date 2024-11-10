package logs

import (
	"database/sql"

	"github.com/mviniciusgc/onfly/src/db"
)

type ClientRepository struct {
	db db.IClientDB
}

func InitializeClienteRepository(db db.IClientDB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (c *ClientRepository) getConnection() (*sql.DB, error) {
	conn, err := c.db.InitDB()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
