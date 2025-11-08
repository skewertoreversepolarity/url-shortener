package sqlite

import (
	"database/sql"
	"fmt"

	// Use a pure-Go SQLite driver to avoid CGO requirements on all platforms
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	//константа операции для ошибок
	const op = "storage.sqlite.New"
	// Открываем соединение с базой данных SQLite
	// driver name for modernc.org/sqlite is "sqlite"
	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	// Проверяем, что соединение действительно работоспособно
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Создаем таблицу, если она не существует
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS url (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Создаем индекс отдельно
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
