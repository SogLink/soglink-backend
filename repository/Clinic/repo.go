package clinic

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableClinic = "clinic"
)

type clinicRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewClinicRepo(db *postgres.PostgresDB) Repository {
	return &clinicRepo{
		table: tableClinic,
		db:    db,
	}
}

func (r clinicRepo) Get(ctx context.Context, params map[string]string) (*entity.Clinic, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"location_id",
		"name",
		"created_at",
		"updated_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "location_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "name":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "created_at":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "updated_at":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var clinic entity.Clinic
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&clinic.ID,
		&clinic.GUID,
		&clinic.Location_ID,
		&clinic.Name,
		&clinic.CreatedAt,
		&clinic.UpdatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &clinic, nil
}

func (r clinicRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Clinic, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"location_id",
		"name",
		"created_at",
		"updated_at",
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

func (r clinicRepo) Create(ctx context.Context, req *entity.Clinic) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"id":          req.ID,
			"guid":        req.GUID,
			"location_id": req.Location_ID,
			"name":        req.Name,
			"created_at":  req.CreatedAt,
			"updated_at":  req.UpdatedAt,
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

func (r clinicRepo) Update(ctx context.Context, req *entity.Clinic) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"guid":        req.GUID,
			"location_id": req.Location_ID,
			"name":        req.Name,
			"created_at":  req.CreatedAt,
			"updated_at":  req.UpdatedAt,
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

func (r clinicRepo) Delete(ctx context.Context, params map[string]string) error {
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
