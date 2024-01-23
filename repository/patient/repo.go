package patient

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tablePatient = "patient"
)

type patientRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewPatientRepo(db *postgres.PostgresDB) Repository {
	return &patientRepo{
		table: tablePatient,
		db:    db,
	}
}

func (r patientRepo) Get(ctx context.Context, params map[string]string) (*entity.Patient, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"patient_id",
		"name",
		"surname",
		"gender",
		"birthday",
		"pinfl",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "patient_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "name":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "surname":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "gender":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "birthday":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "pinfl":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var patient entity.Patient
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&patient.Patient_ID,
		&patient.Name,
		&patient.Surname,
		&patient.Gender,
		&patient.Birthday,
		&patient.Pinfl,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}
	return &patient, nil
}

func (r patientRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Patient, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"patient_id",
		"name",
		"surname",
		"gender",
		"birthday",
		"pinfl",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "name":
			queryBuilder = queryBuilder.Where("name ILIKE '%'||?||'%'", v)
		case "surname":
			queryBuilder = queryBuilder.Where("surname ILIKE '%'||?||'%'", v)
		case "gender":
			queryBuilder = queryBuilder.Where("gender ILIKE '%'||?||'%'", v)
		case "pinfl":
			queryBuilder = queryBuilder.Where("pinfl ILIKE '%'||?||'%'", v)
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

	var patients []*entity.Patient
	for rows.Next() {
		var patient entity.Patient
		if err := r.db.QueryRow(ctx, query, args...).Scan(
			&patient.Patient_ID,
			&patient.Name,
			&patient.Surname,
			&patient.Gender,
			&patient.Birthday,
			&patient.Pinfl,
		); err != nil {
			return nil, r.db.Error(err)
		}

		patients = append(patients, &patient)
	}
	return patients, nil
}

func (r patientRepo) Create(ctx context.Context, req *entity.Patient) error {
	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"patient_id": req.User.ID,
			"name":       req.Name,
			"surname":    req.Surname,
			"gender":     req.Gender,
			"birthday":   req.Birthday,
			"pinfl":      req.Pinfl,
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

func (r patientRepo) Update(ctx context.Context, req *entity.Patient) error {
	if req.User == nil || req.User.ID == 0 {
		return r.db.Error(fmt.Errorf("invalid User"))
	}

	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"name":     req.Name,
			"surname":  req.Surname,
			"gender":   req.Gender,
			"birthday": req.Birthday,
			"pinfl":    req.Pinfl,
		},
	).Where(r.db.Sq.Equal("patient_id", req.Patient_ID))

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

func (r patientRepo) Delete(ctx context.Context, params map[string]string) error {
	queryBuilder := r.db.Sq.Builder.Delete(r.table)

	for k, v := range params {
		switch k {
		case "patient_id":
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
