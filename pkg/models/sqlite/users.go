package sqlite

import (
	"MPBunce/SnippetBox/pkg/models"
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/mattn/go-sqlite3"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES (?, ?, ?, CURRENT_TIMESTAMP)`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if strings.Contains(sqliteErr.Error(), "UNIQUE constraint failed") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
