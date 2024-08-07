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

func (m *UserModel) Authenticate(email, password string) (int, error) {
	// Retrieve the id and hashed password associated with the given email.
	// If matching email doesn't exist, we return the ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte

	row := m.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	// Check whether the hashed password and plain-text password provided match
	// If they don't, we return the ErrInvalidCredentials error.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	// Otherwise, the password is correct. Return the user ID.
	return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	s := &models.User{}
	stmt := `SELECT id, name, email, created FROM users WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.Email, &s.Created)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
