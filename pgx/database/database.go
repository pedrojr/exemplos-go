package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
	ctx  context.Context
}

func Open(url string) (*Database, error) {
	ctx := context.Background()
	conn, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &Database{conn: conn, ctx: ctx}, nil
}

func (d *Database) Close() {
	d.conn.Close()
}

func (d *Database) Execute(sql string, args ...interface{}) error {
	_, err := d.conn.Exec(d.ctx, sql, args...)
	return err
}

func (d *Database) getRecord(dst interface{}, sql string, args ...interface{}) error {
	return pgxscan.Get(d.ctx, d.conn, dst, sql, args...)
}

func (d *Database) Migrate() {
	_, err := d.conn.Exec(d.ctx, `DROP TABLE IF EXISTS public.testes`)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = d.conn.Exec(d.ctx,
		`CREATE TABLE public.testes (
			id serial4 NOT NULL,
			str_teste text NOT NULL,
			int_teste int8 NULL,
			float_teste numeric NULL,
			blob_teste bytea NULL,
			CONSTRAINT testes_pkey PRIMARY KEY (id)
		)`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
