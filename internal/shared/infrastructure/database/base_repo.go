package database

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/shared/domain/port"
	"gorm.io/gorm"
)

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db: db}
}

func (b *BaseRepo) Executor(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(port.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return b.db
}
