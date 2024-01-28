package specialty

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableName = "specialty"
)

type specialtyRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewSpecialtyRepo(db *postgres.PostgresDB) Repository {
	return &specialtyRepo{
		table: tableName,
		db:    db,
	}
}

func (r specialtyRepo) Get(ctx context.Context, params map[string]string) (*entity.Specialty, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"name",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "name":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var specialty entity.Specialty
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&specialty.ID,
		&specialty.Name,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}
	return &specialty, nil
}

func (r specialtyRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Specialty, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"name",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "name":
			queryBuilder = queryBuilder.Where("name ILIKE '%'||?||'%'", v)
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" List")
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, r.db.Error(err)
	}

	var specialties []*entity.Specialty
	for rows.Next() {
		var specialty entity.Specialty
		if err := r.db.QueryRow(ctx, query, args...).Scan(
			&specialty.ID,
			&specialty.Name,
		); err != nil {
			return nil, r.db.Error(err)
		}

		specialties = append(specialties, &specialty)
	}
	return specialties, nil
}
