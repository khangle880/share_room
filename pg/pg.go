package pg

import (
	"context"
	"database/sql"
	"fmt"
)

type repoSvc struct {
	*Queries
	db *sql.DB
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := r.Queries.WithTx(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback %v", err, rbErr)
		}
	} else {
		err = tx.Commit()
	}
	return err
}

func NewRepository(db *sql.DB) *repoSvc {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}



// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
func Open(dataSourceName string) (*sql.DB, error) {
    return sql.Open("postgres", dataSourceName)
}

// StringPtrToNullString converts *string to sql.NullString.
func StringPtrToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
