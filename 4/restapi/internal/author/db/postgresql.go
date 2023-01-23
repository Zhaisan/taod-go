package author

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"restapi/internal/author"
	"restapi/pkg/client/postgresql"
	"restapi/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `INSERT INTO author (name) VALUES ($1) RETURNING id`

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))

			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []author.Author, err error) {
	q := `SELECT id, name FROM public.author`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var ath author.Author

		err = rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil

}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `SELECT id, name FROM public.author WHERE id = $1`

	var ath author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, user author.Author) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
