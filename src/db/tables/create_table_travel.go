package tables

import (
	"database/sql"
)

// Criando a tabela e seus campos
func CreateTableTravel(conn *sql.DB) error {
	createEnumType := `
	DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'travel_status') THEN
			CREATE TYPE travel_status AS ENUM ('solicitado', 'aprovado', 'cancelado');
		END IF;
	END$$;
	`

	_, err := conn.Exec(createEnumType)
	if err != nil {
		return err
	}

	// Cria a tabela travel, se não existir
	createTable := `
	CREATE TABLE IF NOT EXISTS travel (
		id SERIAL PRIMARY KEY,
		requester_name VARCHAR(100) NOT NULL,
		destination VARCHAR(100) NOT NULL,
		departure_date DATE NOT NULL,
		return_date DATE NOT NULL,
		status travel_status NOT NULL DEFAULT 'solicitado',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = conn.Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}
