package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartGinServer() {
	db, handlers := InitServer()
	router := gin.Default()

	router.ContextWithFallback = true

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	infofile, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer infofile.Close()

	log.SetOutput(infofile)

	router.Use(middleware.RequestId)
	router.Use(middleware.Logger(log))
	router.Use(middleware.ErrorMiddleware())

	auth := router.Group("/")
	{
		auth.POST("/register", handlers.User.Register)
		auth.POST("/login", handlers.User.Login)
		auth.POST("/forgot-password", handlers.Token.Create)
		auth.POST("/change-password", handlers.Token.ChangePassword)
	}

	transaction := auth.Group("/")
	transaction.Use(middleware.Authenticate())
	{
		transaction.POST("/transfer", middleware.DBTransactionMiddleware(db), handlers.Transaction.Transfer)
		transaction.POST("/topup", handlers.Transaction.TopUp)
		transaction.GET("/transactions", handlers.Transaction.GetAll)
		transaction.GET("/profiles", handlers.User.GetProfile)
		transaction.GET("/gachas", handlers.Gacha.GetGacha)
		transaction.POST("/gachas", handlers.Gacha.SelectGacha)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
