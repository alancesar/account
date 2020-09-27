package infra

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(host, database, username, password, port string, ssl bool) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host, port, username, database, password)

	if !ssl {
		dsn = dsn + " sslmode=disable"
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
