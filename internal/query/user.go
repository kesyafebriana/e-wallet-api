package query

const (
	CreateUser = `INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	FindUserByEmail = `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email ILIKE $1 AND deleted_at IS NULL`
	FindUserById = `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL`
	UpdatePassword = `UPDATE users SET password = $1, updated_at = NOW() WHERE id = $2 RETURNING id, updated_at`
)