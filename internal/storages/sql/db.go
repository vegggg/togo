package sql

import (
	"context"
	"database/sql"
	"log"

	"github.com/manabie-com/togo/internal/storages"
)

// Helper for working with sqllite
type Helper struct {
	DB   *sql.DB
	Stmt Stmt
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *Helper) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := l.Stmt.RetrieveTasks
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// ValidateUserMaxTodo returns true if user is allowed to create new task
func (l *Helper) ValidateUserMaxTodo(ctx context.Context, userID, createdDate sql.NullString) bool {
	stmt := l.Stmt.ValidateUserMaxTodo
	row := l.DB.QueryRowContext(ctx, stmt, userID, createdDate)
	u := &storages.User{}

	err := row.Scan(&u.ID)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// AddTask adds a new task to DB
func (l *Helper) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := l.Stmt.AddTask
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *Helper) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := l.Stmt.ValidateUser
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// CreateUser create new user
func (l *Helper) CreateUser(ctx context.Context, u *storages.User) bool {
	stmt := l.Stmt.CreateUser
	log.Println("create user", u.ID, u.Password)
	_, err := l.DB.ExecContext(ctx, stmt, &u.ID, &u.Password, &u.MaxTodo)
	if err != nil {
		return false
	}

	return true
}
