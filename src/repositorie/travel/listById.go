package repositorie

import (
	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) ListById(id int) (*entity.TravelEntity, error) {
	query := "SELECT * FROM travel where id = $1"

	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	row := conn.QueryRow(query, id)

	var travel entity.TravelEntity
	err = row.Scan(
		&travel.ID,
		&travel.RequesterName,
		&travel.Destination,
		&travel.DepartureDate,
		&travel.ReturnDate,
		&travel.Status,
		&travel.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &travel, nil
}
