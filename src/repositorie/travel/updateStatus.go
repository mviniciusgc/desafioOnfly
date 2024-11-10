package repositorie

import (
	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) UpdateStatus(id int, status string) (*entity.TravelEntity, error) {
	updateQuery := "UPDATE travel SET status = $1 WHERE id = $2 RETURNING *"

	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var travel entity.TravelEntity
	err = conn.QueryRow(updateQuery, status, id).Scan(
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
