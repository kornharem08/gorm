package sqlwrap

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// ISQLSession extends ISQL with transaction control
type ISQLSession interface {
	ISQL
	Commit() error
	Rollback() error
}

// SQLSession wraps a GORM transaction
type SQLSession struct {
	tx *gorm.DB
}

// NewSQLSession creates a new SQLSession from a GORM transaction
func NewSQLSession(tx *gorm.DB) *SQLSession {
	return &SQLSession{tx: tx}
}

// Find retrieves multiple records
func (s *SQLSession) Find(ctx context.Context, result any, conds ...any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Find(result, conds...).Error
}

// First retrieves the first record
func (s *SQLSession) First(ctx context.Context, result any, conds ...any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).First(result, conds...).Error
}

// Create inserts a new record
func (s *SQLSession) Create(ctx context.Context, value any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Create(value).Error
}

// Update updates a single column
func (s *SQLSession) Update(ctx context.Context, column string, value any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Update(column, value).Error
}

// Updates updates multiple columns with a map or struct
func (s *SQLSession) Updates(ctx context.Context, value any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Updates(value).Error
}

// Delete removes records
func (s *SQLSession) Delete(ctx context.Context, conds ...any) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Delete(nil, conds...).Error
}

// Joins adds a JOIN clause
func (s *SQLSession) Joins(ctx context.Context, query string, args ...any) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Joins(query, args...)
	return s
}

// Where adds a WHERE clause
func (s *SQLSession) Where(ctx context.Context, query any, args ...any) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Where(query, args...)
	return s
}

// Preload preloads associations
func (s *SQLSession) Preload(ctx context.Context, query string, args ...any) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Preload(query, args...)
	return s
}

// Order specifies the order of results
func (s *SQLSession) Order(ctx context.Context, value string) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Order(value)
	return s
}

// Limit sets the maximum number of records
func (s *SQLSession) Limit(ctx context.Context, limit int) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Limit(limit)
	return s
}

// Offset sets the offset for records
func (s *SQLSession) Offset(ctx context.Context, offset int) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Offset(offset)
	return s
}

// Raw executes a raw SQL query
func (s *SQLSession) Raw(ctx context.Context, sql string, values ...any) ISQL {
	if s.tx == nil {
		return &SQLSession{tx: &gorm.DB{Error: errors.New("transaction is nil")}}
	}
	s.tx = s.tx.WithContext(ctx).Raw(sql, values...)
	return s
}

// Exec executes the accumulated query
func (s *SQLSession) Exec(ctx context.Context) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Error
}

// Transaction handles nested transactions
func (s *SQLSession) Transaction(ctx context.Context, fc func(tx ISQL) error) error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	return s.tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fc(&SQLSession{tx: tx})
	})
}

// Commit commits the transaction
func (s *SQLSession) Commit() error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	err := s.tx.Commit().Error
	s.tx = nil // Prevent further operations
	return err
}

// Rollback rolls back the transaction
func (s *SQLSession) Rollback() error {
	if s.tx == nil {
		return errors.New("transaction is nil")
	}
	err := s.tx.Rollback().Error
	s.tx = nil // Prevent further operations
	return err
}
