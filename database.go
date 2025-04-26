package sqlwrap

import "gorm.io/gorm"

type IDatabase interface {
	Table(model any) ISQL
}

type Database struct {
	db *gorm.DB
}

func (d *Database) Table(model any) ISQL {
	return &SQLTable{db: d.db.Model(model)}
}
