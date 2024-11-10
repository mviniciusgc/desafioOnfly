package repositorie

import (
	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) Create(travel entity.TravelEntity) (id *int64, err error) {

	// Obtendo a conexão com o banco
	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}

	// Iniciando a transação
	tx, err := c.db.Begin(conn)
	if err != nil {
		return nil, err
	}

	// Definindo a query de inserção
	query := `
		INSERT INTO travel(requester_name, destination, departure_date, return_date, status, created_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	// Preparando a instrução com a query e a transação
	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Executando a inserção e pegando o ID gerado
	var newID int64
	err = stmt.QueryRow(travel.RequesterName, travel.Destination, travel.DepartureDate, travel.ReturnDate, travel.Status, travel.CreatedAt).Scan(&newID)
	if err != nil {
		return nil, err
	}

	// Confirmando a transação
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Retornando o ID gerado
	id = &newID
	return id, nil
}
