package sql

// Stmt ..
type Stmt struct {
	RetrieveTasks       string
	ValidateUserMaxTodo string
	AddTask             string
	ValidateUser        string
	CreateUser          string
}

var (
	// Sqlite ..
	Sqlite Stmt
	// Psql ..
	Psql Stmt
)

func init() {
	Sqlite = Stmt{
		RetrieveTasks:       "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?",
		ValidateUserMaxTodo: "SELECT id FROM users WHERE max_todo > (SELECT COUNT(id) FROM tasks WHERE user_id = ? AND created_date = ?)",
		AddTask:             "INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)",
		ValidateUser:        "SELECT id FROM users WHERE id = ? AND password = ?",
		CreateUser:          "INSERT INTO users (id, password, max_todo) VALUES(?, ?, ?)",
	}
	Psql = Stmt{
		RetrieveTasks:       "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2",
		ValidateUserMaxTodo: "SELECT id FROM users WHERE max_todo > (SELECT COUNT(id) FROM tasks WHERE user_id = $1 AND created_date = $2)",
		AddTask:             "INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)",
		ValidateUser:        "SELECT id FROM users WHERE id = $1 AND password = $2",
		CreateUser:          "INSERT INTO users (id, password, max_todo) VALUES($1, $2, $3)",
	}
}
