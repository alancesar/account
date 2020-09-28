package main

import (
	"github.com/alancesar/account/account"
	"github.com/alancesar/account/handler"
	"github.com/alancesar/account/infra"
	"github.com/alancesar/account/middleware"
	"github.com/alancesar/account/transaction"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	dbHostEnv     = "DB_HOST"
	dbNameEnv     = "DB_NAME"
	dbUsernameEnv = "DB_USERNAME"
	dbPasswordEnv = "DB_PASSWORD"
	dbPortEnv     = "DB_PORT"
	dbSslEnv      = "DB_SSL"
)

var (
	host     = os.Getenv(dbHostEnv)
	database = os.Getenv(dbNameEnv)
	username = os.Getenv(dbUsernameEnv)
	password = os.Getenv(dbPasswordEnv)
	port     = os.Getenv(dbPortEnv)
	ssl      = os.Getenv(dbSslEnv)
)

func main() {
	// Try to connect to the database
	db, err := infra.NewPostgresConnection(host, database, username, password, port, ssl == "true")
	if err != nil {
		panic(err)
	}

	// Configure logger
	infra.ConfigureLogger(&logrus.JSONFormatter{})

	// Start repositories
	ar := account.NewRepository(db)
	tr := transaction.NewRepository(db)

	// Start services
	as := account.NewAccountService(ar)
	ts := transaction.NewTransactionService(tr)

	// Create server
	engine := infra.CreateServer(gin.Recovery(), middleware.RequestLogger())

	// Start metrics
	root := engine.Group("/")
	handler.StartMetrics(root)

	// Start APIs
	api := root.Group("/api", middleware.ResponseLogger())
	handler.StartApi(api, as, ts)

	// Start server
	infra.StartServer(engine)
}
