package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type ctxKey string

const (
	txCtxKey ctxKey = "TRANSACTION"
)

func CtxWithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey, tx)
}

func TxFromCtx(ctx context.Context) pgx.Tx {
	tx, ok := ctx.Value(txCtxKey).(pgx.Tx)
	if !ok {
		panic(fmt.Sprintf("%T instead of transaction in a context under transaction key", ctx.Value(txCtxKey)))
	}
	return tx
}
