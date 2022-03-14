package pgx

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/sealoftime/adapteris/domain/user"
	"github.com/sealoftime/adapteris/log"
)

var (
	account = columnNamesAliases{
		"id":           "id",
		"registeredAt": "registered_at",
		"role":         "role",
		"shortName":    "short_name",
		"fullName":     "full_name",
		"email":        "email",
		"telegram":     "telegram",
		"vk":           "vk",
		"phoneNumber":  "phone_number",
	}
	accountAllCols = columns{
		account["id"],
		account["registeredAt"],
		account["role"],
		account["shortName"],
		account["fullName"],
		account["email"],
		account["telegram"],
		account["vk"],
		account["phoneNumber"],
	}

	extAccount = columnNamesAliases{
		"id":         "id",
		"service":    "service",
		"externalId": "external_id",
		"accountId":  "account_id",
	}
	extAccountAllCols = columns{
		extAccount["id"],
		extAccount["service"],
		extAccount["externalId"],
		extAccount["accountId"],
	}
)

type AccountStorage struct {
	log log.Logger
}

var _ user.Repository = (*AccountStorage)(nil)

func NewAccountStorage(log log.Logger) *AccountStorage {
	return &AccountStorage{log: log}
}

const (
	upsertExtAccountSqlTempl = `
INSERT INTO external_account (%s) VALUES (%s) 
ON CONFLICT (id) DO UPDATE SET
	external_id=external_account.external_id
RETURNING %s`
)

var (
	selectAccountByIdSql = fmt.Sprintf(`
SELECT %s FROM account WHERE %s=$1`,
		accountAllCols.sqlString(), account["id"],
	)
	selectAccountByExtAccSql = fmt.Sprintf(`
SELECT %s FROM account WHERE %s=(
	SELECT %s FROM external_account WHERE %s=$1
)`,
		accountAllCols.sqlString(), account["id"], extAccount["accountId"], extAccount["externalId"],
	)
	selectAccountByEmailSql = fmt.Sprintf(`
SELECT %s FROM account WHERE %s=$1`,
		accountAllCols.sqlString(), account["email"],
	)
	selectExtAccountsByAccountSql = fmt.Sprintf(`
SELECT %s FROM external_account WHERE %s=$1`,
		extAccountAllCols.sqlString(), extAccount["accountId"],
	)
	upsertAccountSqlTempl = fmt.Sprintf(`
INSERT INTO account (%%s) VALUES (%%s) 
ON CONFLICT (id) DO UPDATE 
SET %s=account.%s,
	%s=account.%s,
	%s=account.%s,
	%s=account.%s,
	%s=account.%s,
	%s=account.%s,
	%s=account.%s
RETURNING %%s`,
		account["role"], account["role"],
		account["shortName"], account["shortName"],
		account["fullName"], account["fullName"],
		account["email"], account["email"],
		account["telegram"], account["telegram"],
		account["vk"], account["vk"],
		account["phoneNumber"], account["phoneNumber"],
	)
)

func (st *AccountStorage) FindById(
	ctx context.Context,
	id int64,
) (*user.Account, error) {
	tx := TxFromCtx(ctx)

	acc, err := mapAccount(tx.QueryRow(ctx, selectAccountByIdSql, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrAccountNotFoundById{Id: id}
		}
		return nil, fmt.Errorf("error querying account by id: %w", err)
	}

	if err := st.queryExtAccountsForAccount(ctx, &acc); err != nil {
		return nil, fmt.Errorf("error querying account by id: %w", err)
	}
	return &acc, nil
}

func (st *AccountStorage) FindByExternalAccount(
	ctx context.Context,
	eac user.ExternalAccount,
) (*user.Account, error) {
	tx := TxFromCtx(ctx)

	acc, err := mapAccount(tx.QueryRow(ctx, selectAccountByExtAccSql, eac.ExternalId))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrAcountNotFoundByExternalAccount{ExternalAccount: eac}
		}
		return nil, fmt.Errorf("error querying account by external_account: %w", err)
	}

	if err := st.queryExtAccountsForAccount(ctx, &acc); err != nil {
		return nil, fmt.Errorf("error querying account by external_account: %w", err)
	}

	return &acc, nil
}

func (st *AccountStorage) FindByEmail(
	ctx context.Context,
	email string,
) (*user.Account, error) {
	tx := TxFromCtx(ctx)

	acc, err := mapAccount(tx.QueryRow(ctx, selectAccountByEmailSql, email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrAccountNotFoundByEmail{Email: email}
		}
		return nil, fmt.Errorf("error querying account by email: %w", err)
	}

	if err := st.queryExtAccountsForAccount(ctx, &acc); err != nil {
		return nil, fmt.Errorf("error querying account by email: %w", err)
	}

	return &acc, nil
}

//queryExtAccountsForAccount queries the user.ExternalAccount linked to the provided user.Account
//and appends them to the acc.
func (st *AccountStorage) queryExtAccountsForAccount(ctx context.Context, acc *user.Account) error {
	tx := TxFromCtx(ctx)

	rows, err := tx.Query(ctx, selectExtAccountsByAccountSql, acc.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			//It is a possible case for account to not have any associated external_accounts
			//For not it's not used anywhered, but implementation of alternative auth is quite likely
			return nil
		}
		return fmt.Errorf("error querying external_accounts for account: %w", err)
	}

	for rows.Next() {
		eac, err := mapExternalAccount(rows)
		if err != nil {
			return fmt.Errorf("error mapping external_account for account: %w", err)
		}
		acc.ExternalAccounts = append(acc.ExternalAccounts, eac)
	}

	return nil
}

func (st *AccountStorage) Save(
	ctx context.Context,
	acc user.Account,
) (saved *user.Account, err error) {
	tx := TxFromCtx(ctx)

	upsertAccountSql, vals := st.buildUpsertAccountQuery(acc)
	dbAcc, err := mapAccount(tx.QueryRow(ctx, upsertAccountSql, vals...))
	if err != nil {
		return nil, fmt.Errorf("error saving account: %w", err)
	}

	b := &pgx.Batch{}
	for _, eac := range acc.ExternalAccounts {
		upsertExtAccountSql, vals := st.buildUpsertExtAccountQuery(dbAcc.Id, eac)
		b.Queue(upsertExtAccountSql, vals...)
	}
	br := tx.SendBatch(ctx, b)
	defer func() {
		if e := br.Close(); e != nil {
			saved = nil
			err = e
		}
	}()

	dbAcc.ExternalAccounts = make([]user.ExternalAccount, len(acc.ExternalAccounts))
	for i := range dbAcc.ExternalAccounts {
		eac, err := mapExternalAccount(br.QueryRow())
		if err != nil {
			return nil, err
		}

		dbAcc.ExternalAccounts[i] = eac
	}

	return &dbAcc, nil
}

func (st *AccountStorage) buildUpsertAccountQuery(acc user.Account) (query string, vals []interface{}) {
	cols, vals := accToVals(acc).split()
	upsertAccountSql := fmt.Sprintf(
		upsertAccountSqlTempl,
		cols.sqlString(),
		cols.sqlParams(),
		accountAllCols.sqlString(),
	)
	st.log.Log(upsertAccountSql)
	return upsertAccountSql, vals
}

func (st *AccountStorage) buildUpsertExtAccountQuery(accId int64, eac user.ExternalAccount) (
	query string,
	vals []interface{},
) {
	cols, vals := eacToVals(accId, eac).split()
	upsertExternalAccountSql := fmt.Sprintf(
		upsertExtAccountSqlTempl,
		cols.sqlString(),
		cols.sqlParams(),
		extAccountAllCols.sqlString(),
	)
	st.log.Log(upsertExternalAccountSql)
	return upsertExternalAccountSql, vals
}

func accToVals(acc user.Account) sqlValues {
	vs := sqlValues{}
	if acc.Id != 0 {
		vs[account["id"]] = acc.Id
	}

	if !acc.RegisteredAt.IsZero() {
		vs[account["registeredAt"]] = acc.RegisteredAt
	}

	if acc.Role != "" {
		vs[account["role"]] = acc.Role
	}

	vs[account["fullName"]] = acc.FullName
	vs[account["shortName"]] = acc.ShortName
	vs[account["email"]] = acc.Email
	vs[account["telegram"]] = acc.Telegram
	vs[account["vk"]] = acc.Vk
	vs[account["phoneNumber"]] = acc.PhoneNumber
	return vs
}

func eacToVals(accId int64, eac user.ExternalAccount) sqlValues {
	vs := sqlValues{}
	if eac.Id != 0 {
		vs[extAccount["id"]] = eac.Id
	}

	if eac.Service != "" {
		vs[extAccount["service"]] = eac.Service
	}

	if eac.ExternalId != "" {
		vs[extAccount["externalId"]] = eac.ExternalId
	}

	vs[extAccount["accountId"]] = accId

	return vs
}

func mapAccount(row pgx.Row) (acc user.Account, err error) {
	err = row.Scan(
		&acc.Id,
		&acc.RegisteredAt,
		&acc.Role,

		&acc.ShortName,
		&acc.FullName,

		&acc.Email,
		&acc.Telegram,
		&acc.Vk,
		&acc.PhoneNumber,
	)
	if err != nil {
		return
	}

	return
}

func mapExternalAccount(row pgx.Row) (eac user.ExternalAccount, err error) {
	var accIdDump int64
	err = row.Scan(
		&eac.Id,
		&eac.Service,
		&eac.ExternalId,
		&accIdDump,
	)
	if err != nil {
		return
	}

	return
}
