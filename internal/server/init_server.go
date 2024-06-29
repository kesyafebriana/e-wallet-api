package server

import (
	"database/sql"
	"log"

	"github.com/kesyafebriana/e-wallet-api/internal/handler"
	"github.com/kesyafebriana/e-wallet-api/internal/pkg/config"
	database "github.com/kesyafebriana/e-wallet-api/internal/pkg/db/connection"
	"github.com/kesyafebriana/e-wallet-api/internal/repository"
	"github.com/kesyafebriana/e-wallet-api/internal/usecase"
)

type Handlers struct {
	User        *handler.User
	Token       *handler.Token
	Transaction *handler.Transaction
	Gacha       *handler.Gacha
}

func InitServer() (*sql.DB, *Handlers) {
	err := config.ConfigInit()
	if err != nil {
		log.Fatalf("error configuration: %s", err.Error())
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err.Error())
	}

	walletRepository := repository.NewWalletRepository(db)
	gachaRepository := repository.NewGachaRepository(db)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserImplementation(userRepository, walletRepository)
	userHandler := handler.NewUser(userUsecase)

	tokenRepository := repository.NewTokenRepository(db)
	tokenUsecase := usecase.NewTokenImplementation(tokenRepository, userRepository)
	tokenHandler := handler.NewToken(tokenUsecase)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionUsecase := usecase.NewTransactionImplementation(transactionRepository, walletRepository, userRepository, gachaRepository)
	transactionHandler := handler.NewTransaction(transactionUsecase)

	gachaUsecase := usecase.NewGachaImplementation(gachaRepository)
	gachaHandler := handler.NewGacha(gachaUsecase, transactionUsecase)

	handlers := &Handlers{
		User:        userHandler,
		Token:       tokenHandler,
		Transaction: transactionHandler,
		Gacha:       gachaHandler,
	}

	return db, handlers
}
