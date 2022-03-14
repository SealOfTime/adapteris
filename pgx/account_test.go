package pgx

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/user"
)

func TestSave(t *testing.T) {
	conn := mustConnect(t, os.Getenv("PGX_TEST_DATABASE"))
	mustCreateSchema(t, conn)

	tx := mustBeginTx(t, conn)
	ctx := CtxWithTx(context.Background(), tx)
	given, _ := user.NewAccount("Петька", "email@example.com", "id0000001")
	given.LinkExternalAccount(user.ExternalAccount{Service: "vk", ExternalId: "sealoftime"})

	as := AccountStorage{}
	actual, err := as.Save(ctx, given)
	if err != nil {
		t.Fatalf("couldn't save: %+v", err)
	}
	if err := tx.Commit(context.TODO()); err != nil {
		tx.Rollback(context.TODO())
		t.Fatalf("couldn't commit tx: %+v. Rolling back\n", err)
	}

	if actual.Id == 0 {
		t.Errorf("expected Account Id not to be nil after saving to a database, found: %v", actual.Id)
	}
	if actual.ShortName != given.ShortName {
		t.Errorf("expected shortname %s, found %s", given.ShortName, actual.ShortName)
	}
	if actual.Email != given.Email {
		t.Errorf("expected email %s, found %s", given.Email, actual.Email)
	}
}

func mustCreateSchema(t testing.TB, conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS account (
			id	bigserial PRIMARY KEY,
			registered_at timestamp DEFAULT now(),
			role	role NOT NULL,

			short_name text NOT NULL,
			full_name text,

			email text NOT NULL,
			vk text,
			tg text,
			phone text
		);
		CREATE TABLE IF NOT EXISTS external_account (
			id bigserial PRIMARY KEY,
			service varchar(32),
			external_id text,
			account_id bigint REFERENCES account(id)
		);`)
	if err != nil {
		t.Fatalf("couldn't create schema: %+v", err)
	}
}

func mustBeginTx(t testing.TB, conn *pgx.Conn) pgx.Tx {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		t.Fatalf("error starting transaction: %+v", err)
	}
	return tx
}
