package query

const (
	CreateWallet       = `INSERT INTO wallets(user_id) VALUES ($1) RETURNING id, wallet_number, balance, created_at, updated_at`
	FindByWalletNumber = `SELECT id, user_id, wallet_number, balance, created_at, updated_at FROM wallets WHERE wallet_number = $1`
	FindByUserId       = `SELECT id, user_id, wallet_number, balance, created_at, updated_at FROM wallets WHERE user_id = $1`
	IncreaseBalance    = `UPDATE wallets SET balance = balance + $1 WHERE wallet_number = $2 RETURNING id, user_id, wallet_number, balance, created_at, updated_at`
	DecreaseBalance    = `UPDATE wallets SET balance = balance - $1 WHERE wallet_number = $2 RETURNING id, user_id, wallet_number, balance, created_at, updated_at`
)
