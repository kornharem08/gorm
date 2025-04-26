package sqlwrap

import (
	"context"

	"gorm.io/gorm"
)

type ISQL interface {
	Find(ctx context.Context, result any, conds ...any) error
	First(ctx context.Context, result any, conds ...any) error
	Create(ctx context.Context, value any) error
	Update(ctx context.Context, column string, value any) error
	Delete(ctx context.Context, conds ...any) error
}

type SQLTable struct {
	db *gorm.DB
}

func (s *SQLTable) Find(ctx context.Context, result any, conds ...any) error {
	return s.db.WithContext(ctx).Find(result, conds...).Error
}

func (s *SQLTable) First(ctx context.Context, result any, conds ...any) error {
	return s.db.WithContext(ctx).First(result, conds...).Error
}

func (s *SQLTable) Create(ctx context.Context, value any) error {
	return s.db.WithContext(ctx).Create(value).Error
}

func (s *SQLTable) Update(ctx context.Context, column string, value any) error {
	return s.db.WithContext(ctx).Update(column, value).Error
}

func (s *SQLTable) Delete(ctx context.Context, conds ...any) error {
	return s.db.WithContext(ctx).Delete(nil, conds...).Error
}
