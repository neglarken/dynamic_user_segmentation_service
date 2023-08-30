package repository

import (
	"time"

	"github.com/joho/sqltocsv"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository/repoerrors"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

type RecordsRepository struct {
	store *store.Store
}

func NewRecordsRepository(store *store.Store) *RecordsRepository {
	return &RecordsRepository{store}
}

func (r *RecordsRepository) Create(rec *entity.Records) error {
	err := r.store.Db.QueryRow(
		"INSERT INTO records (user_id, slug_title, operation) values ($1, $2, $3) RETURNING id",
		&rec.UserId,
		&rec.SlugTitle,
		&rec.Operation,
	).Scan(&rec.Id)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"slugs_title_key\"" {
			return repoerrors.ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (r *RecordsRepository) GetToCsv(time time.Time) error {
	// recs := make([]*entity.Records, 0)
	rows, err := r.store.Db.Query(
		"SELECT user_id, slug_title, operation, created_at FROM records WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2",
		time.Year(),
		int(time.Month()),
	)
	defer rows.Close()
	if err != nil {
		return err
	}
	if err := sqltocsv.WriteFile("records/records.csv", rows); err != nil {
		return err
	}
	return nil
	// for rows.Next() {
	// 	rec := &entity.Records{}
	// 	err := rows.Scan(
	// 		&rec.UserId,
	// 		&rec.SlugTitle,
	// 		&rec.Operation,
	// 		&rec.CreatedAt,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	recs = append(recs, rec)
	// }
	// if len(recs) == 0 {
	// 	return nil, repoerrors.ErrNotFound
	// }
	// return recs, nil
}
