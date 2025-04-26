package sqlwrap

import (
	"context"

	"gorm.io/gorm"
)

// ISQL defines the interface for SQL operations
type ISQL interface {
	Find(ctx context.Context, result any, conds ...any) error
	First(ctx context.Context, result any, conds ...any) error
	Create(ctx context.Context, value any) error
	Update(ctx context.Context, column string, value any) error
	Updates(ctx context.Context, value any) error
	Delete(ctx context.Context, conds ...any) error
	Joins(ctx context.Context, query string, args ...any) ISQL
	Where(ctx context.Context, query any, args ...any) ISQL
	Preload(ctx context.Context, query string, args ...any) ISQL
	Order(ctx context.Context, value string) ISQL
	Limit(ctx context.Context, limit int) ISQL
	Offset(ctx context.Context, offset int) ISQL
	Raw(ctx context.Context, sql string, values ...any) ISQL
	Exec(ctx context.Context) error
	Transaction(ctx context.Context, fc func(tx ISQL) error) error
}

// SQLTable wraps GORM DB
type SQLTable struct {
	db *gorm.DB
}

// NewSQLTable creates a new SQLTable instance
func NewSQLTable(db *gorm.DB) *SQLTable {
	return &SQLTable{db: db}
}

// Find retrieves multiple records
func (s *SQLTable) Find(ctx context.Context, result any, conds ...any) error {
	return s.db.WithContext(ctx).Find(result, conds...).Error
}

// First retrieves the first record
func (s *SQLTable) First(ctx context.Context, result any, conds ...any) error {
	return s.db.WithContext(ctx).First(result, conds...).Error
}

// Create inserts a new record
func (s *SQLTable) Create(ctx context.Context, value any) error {
	return s.db.WithContext(ctx).Create(value).Error
}

// Update updates a single column
func (s *SQLTable) Update(ctx context.Context, column string, value any) error {
	return s.db.WithContext(ctx).Update(column, value).Error
}

// Updates updates multiple columns with a map or struct
func (s *SQLTable) Updates(ctx context.Context, value any) error {
	return s.db.WithContext(ctx).Updates(value).Error
}

// Delete removes records
func (s *SQLTable) Delete(ctx context.Context, conds ...any) error {
	return s.db.WithContext(ctx).Delete(nil, conds...).Error
}

// Joins adds a JOIN clause
func (s *SQLTable) Joins(ctx context.Context, query string, args ...any) ISQL {
	s.db = s.db.WithContext(ctx).Joins(query, args...)
	return s
}

// Where adds a WHERE clause
func (s *SQLTable) Where(ctx context.Context, query any, args ...any) ISQL {
	s.db = s.db.WithContext(ctx).Where(query, args...)
	return s
}

// Preload preloads associations
func (s *SQLTable) Preload(ctx context.Context, query string, args ...any) ISQL {
	s.db = s.db.WithContext(ctx).Preload(query, args...)
	return s
}

// Order specifies the order of results
func (s *SQLTable) Order(ctx context.Context, value string) ISQL {
	s.db = s.db.WithContext(ctx).Order(value)
	return s
}

// Limit sets the maximum number of records
func (s *SQLTable) Limit(ctx context.Context, limit int) ISQL {
	s.db = s.db.WithContext(ctx).Limit(limit)
	return s
}

// Offset sets the offset for records
func (s *SQLTable) Offset(ctx context.Context, offset int) ISQL {
	s.db = s.db.WithContext(ctx).Offset(offset)
	return s
}

// Raw executes a raw SQL query
func (s *SQLTable) Raw(ctx context.Context, sql string, values ...any) ISQL {
	s.db = s.db.WithContext(ctx).Raw(sql, values...)
	return s
}

// Exec executes the accumulated query
func (s *SQLTable) Exec(ctx context.Context) error {
	return s.db.WithContext(ctx).Error
}

// Transaction handles transactions
func (s *SQLTable) Transaction(ctx context.Context, fc func(tx ISQL) error) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fc(&SQLTable{db: tx})
	})
}
