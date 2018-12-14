// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/xo/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// PgExpandarray calls the stored procedure 'information_schema._pg_expandarray(anyarray) SETOF record' on db.
func PgExpandarray(db XODB, v0 pgtypes.Anyarray) ([]pgtypes.Record, error) {
	var err error

	// sql query
	const sqlstr = `SELECT information_schema._pg_expandarray($1)`

	// run query
	var ret []pgtypes.Record
	XOLog(sqlstr, v0)
	err = db.QueryRow(sqlstr, v0).Scan(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
