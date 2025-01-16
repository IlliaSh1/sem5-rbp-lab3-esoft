package transactor

import (
	"context"

	"gorm.io/gorm"
)

const txKey = "txKey"

func insertTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func GetTransaction(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return tx
}
