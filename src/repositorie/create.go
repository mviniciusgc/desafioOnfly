package repositorie

import (
	"fmt"

	"github.com/mviniciusgc/onfly/src/entity"
)

func (c *ClientRepository) Create(travel entity.TravelEntity) (id *int64, err error) {

	nameTable := "travel"
	fields := []string{"requester_name", "destination", "departure_date", "return_date", "status", "created_at"}

	//Iniciando a conexão com o banco
	conn, err := c.db.InitDB()
	if err != nil {
		return nil, err
	}

	//iniciano a transação
	tx, err := c.db.Begin(conn)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// Cria uma instrução dentro de uma transação para ser executada
	stmt, err := c.db.Prepare(tx, nameTable, fields)
	if err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Println("asdasd")
		return nil, err
	}
	fmt.Printf("%+v\n", travel)
	_, err = c.db.Exec(stmt, []interface{}{
		travel.RequesterName,
		travel.Destination,
		travel.DepartureDate,
		travel.ReturnDate,
		travel.Status,
		travel.CreatedAt,
	})

	// Fechando a instrução
	err = c.db.Close(stmt)
	if err != nil {
		return nil, err
	}

	// confirmando a transação
	err = c.db.Commit(tx)
	if err != nil {
		return nil, err
	}

	return nil, err
}
