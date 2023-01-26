package db

import (
	"context"
	author2 "restapi/internal/author"
	"restapi/internal/book"
	"restapi/pkg/client/postgresql"
	"restapi/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) FindAll(ctx context.Context) (u []book.Book, err error) {
	q := `SELECT id, name FROM public.book`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var bk book.Book

		err = rows.Scan(&bk.ID, &bk.Name)
		if err != nil {
			return nil, err
		}

		sq := `SELECT a.id, a.name FROM book_authors ba
    			JOIN public.author a on a.id = ba.author_id
        			WHERE ba.book_id = $1`

		authorsRows, err := r.client.Query(ctx, sq, bk.ID)

		if err != nil {
			return nil, err
		}

		authors := make([]author2.Author, 0)

		for authorsRows.Next() {
			var ath author2.Author

			err = authorsRows.Scan(&ath.ID, &ath.Name)
			if err != nil {
				return nil, err
				authors = append(authors, ath)
			}
		}
		bk.Authors = authors


		books = append(books, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil

}
