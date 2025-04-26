package sqlwrap

import (
	"context"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ISQLConnect interface {
	Close() error
	Database() IDatabase
	NewSession() (ISQLSession, error)
}

type SQLConnect struct {
	db       *gorm.DB
	database IDatabase
}

func New(dsn string) (ISQLConnect, error) {

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	return &SQLConnect{
		db:       db,
		database: &Database{db: db},
	}, nil
}

func (conn *SQLConnect) Close() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (conn *SQLConnect) Database() IDatabase {
	return conn.database
}

func (conn *SQLConnect) NewSession() (ISQLSession, error) {
	tx := conn.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &SQLSession{tx: tx}, nil
}
