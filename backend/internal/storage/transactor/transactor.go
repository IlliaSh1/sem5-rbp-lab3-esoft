package transactor

import (
	"context"

	"gorm.io/gorm"
)

type Transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) *Transactor {
	return &Transactor{
		db: db,
	}
}

func (transactor *Transactor) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	return fn(insertTransaction(ctx, transactor.db))
}
