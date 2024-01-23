package location

import (
	"context"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableLocation = "location"
)

type locationRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewLocationRepo(db *postgres.PostgresDB) Repository {
	return &locationRepo{
		table: tableLocation,
		db:    db,
	}
}

func (r locationRepo) Get(ctx context.Context, params map[string]string) (*entity.Location, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"city",
		"region",
		"latitude",
		"longitude",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "city":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "region":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "latitude":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "longitude":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var location entity.Location
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&location.ID,
		&location.City,
		&location.Region,
		&location.Latitude,
		&location.Longitude,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &location, nil
}

func (r locationRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Location, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"city",
		"region",
		"latitude",
		"longitude",
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

	var clinics []*entity.Clinic
	for rows.Next() {
		var clinic entity.Clinic
		if err := rows.Scan(
			&clinic.ID,
			&clinic.GUID,
			&clinic.Location_ID,
			&clinic.Name,
			&clinic.CreatedAt,
			&clinic.UpdatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		clinics = append(clinics, &clinic)
	}

	return clinics, nil
}
