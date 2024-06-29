package query

const (
	CreateToken = `INSERT INTO password_tokens(user_id, token) VALUES ($1, $2) RETURNING id, expired_at, created_at, updated_at`
	FindToken = `SELECT id, user_id, token, expired_at, deleted_at FROM password_tokens WHERE token = $1`
	DeleteResetToken = `UPDATE password_tokens SET deleted_at = NOW(), updated_at = NOW() WHERE id = $1 RETURNING id, token, updated_at, deleted_at`
)