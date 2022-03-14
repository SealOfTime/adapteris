package pgx

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

// nSqlParams returns a string of N SQL Parameters.
func nSQLParams(n int) string {
	params := make([]string, n)
	for i := 1; i <= n; i++ {
		params[i-1] = fmt.Sprintf("$%d", i)
	}
	return strings.Join(params, ", ")
}

// colName is column's name in the scema.
type colName = string

// colNameAlias is column's name in the code.
type colNameAlias = string

// columnNamesAliases is a mapping between columns' pseudonym in code and actual name in the schema.
type columnNamesAliases map[colNameAlias]colName

// columns is a slice of columns.
type columns []colName

func (cs columns) sqlParams() string {
	return nSQLParams(len(cs))
}

// sqlString returns SQL Attributes in a string useful for inclusion as a part of expression.
func (cs columns) sqlString() string {
	return strings.Join(cs, ", ")
}

// sqlStringWithRoot append root as a suffix to the columns' names and then returns as a string
// useful for inclusion as a part of epxression.
func (cs columns) sqlStringWithRoot(root string) string {
	attrs := make([]string, len(cs))
	for i, c := range cs {
		attrs[i] = fmt.Sprintf("%s.%s", root, c)
	}
	return columns(attrs).sqlString()
}

type sqlValue = interface{}

// sqlValues is a mapping of values to the corresponding columns.
type sqlValues map[colName]sqlValue

// split returns a slice of column's names and values for the current query in a matching order.
func (v sqlValues) split() (columns, []sqlValue) {
	cols, vals := make([]colName, 0, len(v)), make([]sqlValue, 0, len(v))
	for col, val := range v {
		cols = append(cols, col)
		vals = append(vals, val)
	}
	return cols, vals
}

// Processes batch and then closes it
func processBatch(br pgx.BatchResults, processFunc func(br pgx.BatchResults) error) (err error) {
	err = processFunc(br)

	if brErr := br.Close(); brErr != nil {
		if err != nil {
			return fmt.Errorf("close batch results failed %+v after: %w", brErr, err)
		}
		return brErr
	}

	return nil
}
