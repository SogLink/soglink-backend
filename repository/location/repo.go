package location

import (
	"context"
	"fmt"

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

	var locations []*entity.Location
	for rows.Next() {
		var location entity.Location
		if err := rows.Scan(
			&location.ID,
			&location.City,
			&location.Region,
			&location.Latitude,
			&location.Longitude,
		); err != nil {
			return nil, r.db.Error(err)
		}

		locations = append(locations, &location)
	}

	return locations, nil
}

func (r locationRepo) Create(ctx context.Context, req *entity.Location) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"id":        req.ID,
			"city":      req.City,
			"region":    req.Region,
			"latitude":  req.Latitude,
			"longitude": req.Longitude,
		},
	)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return r.db.ErrSQLBuild(err, r.table+" Create")
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return r.db.Error(err)
	}

	return nil
}

func (r locationRepo) Update(ctx context.Context, req *entity.Location) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"city":      req.City,
			"region":    req.Region,
			"latitude":  req.Latitude,
			"longitude": req.Longitude,
		},
	).Where(r.db.Sq.Equal("id", req.ID))

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return r.db.ErrSQLBuild(err, r.table+" Update")
	}

	commandTag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return r.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return r.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (r locationRepo) Delete(ctx context.Context, params map[string]string) error {
	queryBuilder := r.db.Sq.Builder.Delete(r.table)
	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return r.db.ErrSQLBuild(err, r.table+" Delete")
	}

	commandTag, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return r.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return r.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}
