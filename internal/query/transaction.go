package query

const (
	CalculateTransaction = `SELECT count(*) FROM transactions WHERE sender_wallet_id = $1 OR recipient_wallet_id = $1`
	CreateTransaction    = `INSERT INTO transactions(sender_wallet_id, recipient_wallet_id, amount, source_of_funds, description) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
)
