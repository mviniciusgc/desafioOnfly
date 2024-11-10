package tables

import (
	"database/sql"
)

func CreateTableLogs(conn *sql.DB) error {

	createEnumType := `
	DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'log_type') THEN
			CREATE TYPE log_type AS ENUM ('INFO', 'ERROR', 'WARNING', 'DEBUG');
		END IF;
	END$$;
	`

	_, err := conn.Exec(createEnumType)
	if err != nil {
		return err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		log_type log_type NOT NULL,
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Data e hora da criação
	);
	`

	_, err = conn.Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}
