package port

import "context"

type TxKey struct{}

type TxManager interface {
	WithinTx(ctx context.Context, fn func(ctx context.Context) error) error
}
