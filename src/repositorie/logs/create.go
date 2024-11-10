package logs

import (
	"time"

	"github.com/mviniciusgc/onfly/src/entity"
	"github.com/mviniciusgc/onfly/src/enum"
)

func (c *ClientRepository) CreateLog(logType enum.LogType, message string) error {

	log := entity.LogEntity{
		LogType:   logType,
		Message:   message,
		CreatedAt: time.Now(),
	}

	nameTable := "logs"
	fields := []string{"log_type", "message", "created_at"}

	conn, err := c.getConnection()
	if err != nil {
		return err
	}

	// Inicia a transação
	tx, err := c.db.Begin(conn)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Realiza rollback em caso de erro

	// Prepara a instrução SQL de inserção
	stmt, err := c.db.Prepare(tx, nameTable, fields)
	if err != nil {
		return err
	}

	// Executa a instrução com os valores do log
	_, err = c.db.Exec(stmt, []interface{}{
		log.LogType.String(),
		log.Message,
		log.CreatedAt,
	})
	if err != nil {
		return err
	}

	// Fecha a instrução
	err = c.db.Close(stmt)
	if err != nil {
		return err
	}

	// Commit da transação
	err = c.db.Commit(tx)
	if err != nil {
		return err
	}

	return nil
}
