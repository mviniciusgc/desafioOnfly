package repositorie

import (
	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) ListAll(status string) ([]entity.TravelEntity, error) {
	query := "SELECT * FROM travel"
	var args []interface{}

	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}

	if status != "" {
		query += " WHERE status = $1"
		args = append(args, status)
	}

	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var travels []entity.TravelEntity
	for rows.Next() {
		var travel entity.TravelEntity
		if err := rows.Scan(
			&travel.ID,
			&travel.RequesterName,
			&travel.Destination,
			&travel.DepartureDate,
			&travel.ReturnDate,
			&travel.Status,
			&travel.CreatedAt,
		); err != nil {
			return nil, err
		}
		travels = append(travels, travel)
	}

	return travels, nil
}
