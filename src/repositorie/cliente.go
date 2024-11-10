package repositorie

import "github.com/mviniciusgc/onfly/src/db"

type ClientRepository struct {
	db db.IClientDB
}

func InitializeClienteRepository(db db.IClientDB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}
