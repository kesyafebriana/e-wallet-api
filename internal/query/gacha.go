package query

const (
	FindAllGacha       = `SELECT id, amount FROM gacha_boxes WHERE deleted_at IS NULL LIMIT 9`
	FindOneGacha       = `SELECT id FROM gacha_attempts WHERE deleted_at IS NULL AND user_id = $1`
	CreateGachaAttempt = `INSERT INTO gacha_attempts(user_id) VALUES ($1) RETURNING id, user_id, created_at, updated_at`
	DeleteGachaAttempt = `UPDATE gacha_attempts SET deleted_at = NOW() WHERE id = $1`
)
