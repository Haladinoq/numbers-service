package utils

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SpeakDB instance myspeak Database
type SpeakDB struct {
	DB *gorm.DB
}

//InitialConection mock connection
func InitialConection() (sqlmock.Sqlmock, *gorm.DB, error) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return mock, gormDB, nil
}
