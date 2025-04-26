package sqlwrap

import (
	"context"

	"gorm.io/gorm"
)

type ISQLSession interface {
	ISQL
	Commit() error
	Rollback() error
}

type SQLSession struct {
	tx *gorm.DB
}

func (s *SQLSession) Find(ctx context.Context, result any, conds ...any) error {
	return s.tx.WithContext(ctx).Find(result, conds...).Error
}

func (s *SQLSession) First(ctx context.Context, result any, conds ...any) error {
	return s.tx.WithContext(ctx).First(result, conds...).Error
}

func (s *SQLSession) Create(ctx context.Context, value any) error {
	return s.tx.WithContext(ctx).Create(value).Error
}

func (s *SQLSession) Update(ctx context.Context, column string, value any) error {
	return s.tx.WithContext(ctx).Update(column, value).Error
}

func (s *SQLSession) Delete(ctx context.Context, conds ...any) error {
	return s.tx.WithContext(ctx).Delete(nil, conds...).Error
}

func (s *SQLSession) Commit() error {
	return s.tx.Commit().Error
}

func (s *SQLSession) Rollback() error {
	return s.tx.Rollback().Error
}
