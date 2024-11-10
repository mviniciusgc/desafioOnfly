package repositorie

import (
	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) Create(travel entity.TravelEntity) (id *int64, err error) {

	nameTable := "travel"
	fields := []string{"requester_name", "destination", "departure_date", "return_date", "status", "created_at"}

	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}

	tx, err := c.db.Begin(conn)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := c.db.Prepare(tx, nameTable, fields)
	if err != nil {
		return nil, err
	}

	_, err = c.db.Exec(stmt, []interface{}{
		travel.RequesterName,
		travel.Destination,
		travel.DepartureDate,
		travel.ReturnDate,
		travel.Status,
		travel.CreatedAt,
	})

	err = c.db.Close(stmt)
	if err != nil {
		return nil, err
	}

	err = c.db.Commit(tx)
	if err != nil {
		return nil, err
	}

	return nil, err
}
