package sqlite

import (
	"MPBunce/SnippetBox/pkg/models"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (?, ?, CURRENT_TIMESTAMP, datetime(CURRENT_TIMESTAMP, '+' || ? || ' days'))`
	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		log.Printf("Error executing insert: %v", err)
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) GetAll() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
