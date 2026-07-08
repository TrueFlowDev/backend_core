package database

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"gorm.io/gorm"
)

var _ port.TxManager = (*TxManager)(nil)

type TxManager struct{ db *gorm.DB }

func NewTxManager(db *gorm.DB) *TxManager {
	return &TxManager{db: db}
}

func (t *TxManager) WithinTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		ctxWithTx := context.WithValue(ctx, port.TxKey{}, tx)
		return fn(ctxWithTx)
	})
}
