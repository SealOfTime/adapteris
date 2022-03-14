package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sealoftime/adapteris/log"
	"github.com/sealoftime/adapteris/pgx"
)

func PgxTransactional(pool *pgxpool.Pool, log log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		tx, err := pool.Begin(ctx)
		if err != nil {
			log.Log("couldn't init a transaction: %+v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.SetUserContext(pgx.CtxWithTx(ctx, tx))

		if err := c.Next(); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				log.Log("couldn't commit a tx: %+v", err)
			}

			return err //error was already handled and respective status and body set
		}

		if err := tx.Commit(ctx); err != nil {
			log.Log("couldn't commit a transaction: %+v", err)
			return c.
				Status(fiber.StatusInternalServerError).
				SendString("Couldn't commit transaction")
		}

		return nil
	}
}
