package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

type PGStore struct {
	ledger     string
	connString string
	conn       *pgx.Conn
}

func (s *PGStore) connect() error {
	log.Println("initiating postgres connection")

	conn, err := pgx.Connect(
		context.TODO(),
		s.connString,
	)

	if err != nil {
		return err
	}

	s.conn = conn

	return nil
}

func (s *PGStore) Conn() *pgx.Conn {
	if s.conn.IsClosed() {
		err := s.connect()

		if err != nil {
			panic(err)
		}
	}

	return s.conn
}

func NewStore(name string) (*PGStore, error) {
	store := &PGStore{
		ledger:     name,
		connString: viper.GetString("storage.postgres.conn_string"),
	}

	err := store.connect()

	if err != nil {
		return store, err
	}

	return store, nil
}

func (s *PGStore) Initialize() error {
	statements := `
		CREATE TABLE IF NOT EXISTS transactions (
			"ledger"    varchar,
			"id"        bigint,
			"timestamp" varchar,
			"reference" varchar,
			"hash"      varchar,

			UNIQUE("ledger", "id"),
			UNIQUE("ledger", "reference")
		);

		CREATE TABLE IF NOT EXISTS postings (
			"ledger"      varchar,
			"id"          smallint,
			"txid"        bigint,
			"source"      varchar,
			"destination" varchar,
			"amount"      bigint,
			"asset"       varchar,

			UNIQUE("id", "txid")
		);

		CREATE INDEX IF NOT EXISTS p_c0 ON postings (
			"ledger",
			"txid" DESC,
			"source",
			"destination"
		);

		CREATE TABLE IF NOT EXISTS metadata (
			"meta_id" bigint,
			"meta_target_type" varchar,
			"meta_target_id" varchar,
			"meta_key" varchar,
			"meta_value" varchar,
			"timestamp" varchar,
		
			UNIQUE("meta_id")
		);
		
		CREATE INDEX IF NOT EXISTS m_i0 ON metadata (
			"meta_target_type",
			"meta_target_id"
		);

		CREATE OR REPLACE VIEW addresses AS SELECT ledger, address FROM (
			SELECT ledger, source as address FROM postings GROUP BY ledger, source
			UNION
			SELECT ledger, destination as address FROM postings GROUP BY ledger, destination
		) addr_agg GROUP BY address, ledger;
	`

	_, err := s.Conn().Exec(
		context.Background(),
		statements,
	)

	if err != nil {
		panic(err)
	}

	return err
}

func (s *PGStore) Close() {
	s.conn.Close(context.TODO())
}
