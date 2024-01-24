package emr

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableEmr = "emr"
)

type emrRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewEmrRepo(db *postgres.PostgresDB) Repository {
	return &emrRepo{
		table: tableEmr,
		db:    db,
	}
}

func (r emrRepo) Get(ctx context.Context, params map[string]string) (*entity.Emr, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"doctor_id",
		"patient_id",
		"diagnoses_text",
		"prescriptions_text",
		"created_at",
		"updated_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "doctor_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "patient_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "diagnoses_text":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "prescriptions_text":
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

	var emr entity.Emr
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&emr.ID,
		&emr.Doctor_ID,
		&emr.Patient_ID,
		&emr.Diagnoses_text,
		&emr.Prescriptions_text,
		&emr.CreatedAt,
		&emr.UpdatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &emr, nil
}

func (r emrRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Emr, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"doctor_id",
		"patient_id",
		"diagnoses_text",
		"prescriptions_text",
		"created_at",
		"updated_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "diagnoses_text":
			queryBuilder = queryBuilder.Where("diagnoses_text ILIKE '%'||?||'%'", v)
		case "prescriptions_text":
			queryBuilder = queryBuilder.Where("prescriptions_text ILIKE '%'||?||'%'", v)
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

	var emrs []*entity.Emr
	for rows.Next() {
		var emr entity.Emr
		if err := rows.Scan(
			&emr.ID,
			&emr.Doctor_ID,
			&emr.Patient_ID,
			&emr.Diagnoses_text,
			&emr.Prescriptions_text,
			&emr.CreatedAt,
			&emr.UpdatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		emrs = append(emrs, &emr)
	}

	return emrs, nil
}

func (r emrRepo) Create(ctx context.Context, req *entity.Emr) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"id":                 req.ID,
			"doctor_id":          req.Doctor_ID,
			"patient_id":         req.Patient_ID,
			"diagnoses_text":     req.Diagnoses_text,
			"prescriptions_text": req.Patient_ID,
			"created_at":         req.CreatedAt,
			"updated_at":         req.UpdatedAt,
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

func (r emrRepo) Update(ctx context.Context, req *entity.Emr) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"doctor_id":          req.Doctor_ID,
			"patient_id":         req.Patient_ID,
			"diagnoses_text":     req.Diagnoses_text,
			"prescriptions_text": req.Patient_ID,
			"created_at":         req.CreatedAt,
			"updated_at":         req.UpdatedAt,
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

func (r emrRepo) Delete(ctx context.Context, params map[string]string) error {
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
