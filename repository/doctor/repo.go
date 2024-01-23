package doctor

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableDoctor = "doctor"
)

type doctorRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewDoctorRepo(db *postgres.PostgresDB) Repository {
	return &doctorRepo{
		table: tableDoctor,
		db:    db,
	}
}

func (r doctorRepo) Get(ctx context.Context, params map[string]string) (*entity.Doctor, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"doctor_id",
		"clinic_id",
		"name",
		"surname",
		"birthday",
		"gender",
		"education",
		"certificates",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "doctor_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "clinic_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "name":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "surname":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "birthday":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "gender":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "education":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "certificates":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var doctor entity.Doctor
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&doctor.Doctor_ID,
		&doctor.Clinic_ID,
		&doctor.Name,
		&doctor.Surname,
		&doctor.Birthday,
		&doctor.Gender,
		&doctor.Education,
		&doctor.Certificates,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &doctor, nil
}

func (r doctorRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Doctor, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"doctor_id",
		"clinic_id",
		"name",
		"surname",
		"birthday",
		"gender",
		"education",
		"certificates",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "doctor_id":
			queryBuilder = queryBuilder.Where("name ILIKE '%'||?||'%'", v)
		case "surname":
			queryBuilder = queryBuilder.Where("surname ILIKE '%'||?||'%'", v)
		case "gender":
			queryBuilder = queryBuilder.Where("gender ILIKE '%'||?||'%'", v)
		case "education":
			queryBuilder = queryBuilder.Where("education ILIKE '%'||?||'%'", v)
		case "certificates":
			queryBuilder = queryBuilder.Where("certificates ILIKE '%'||?||'%'", v)
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

	var doctors []*entity.Doctor
	for rows.Next() {
		var doctor entity.Doctor
		if err := rows.Scan(
			&doctor.Doctor_ID,
			&doctor.Clinic_ID,
			&doctor.Name,
			&doctor.Surname,
			&doctor.Birthday,
			&doctor.Gender,
			&doctor.Education,
			&doctor.Certificates,
		); err != nil {
			return nil, r.db.Error(err)
		}

		doctors = append(doctors, &doctor)
	}

	return doctors, nil
}

func (r doctorRepo) Create(ctx context.Context, req *entity.Doctor) error {
	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"doctor_id":    req.Doctor_ID,
			"clinic_id":    req.Clinic_ID,
			"name":         req.Name,
			"surname":      req.Surname,
			"birthday":     req.Birthday,
			"gender":       req.Gender,
			"education":    req.Education,
			"certificates": req.Certificates,
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

func (r doctorRepo) Update(ctx context.Context, req *entity.Doctor) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"clinic_id":    req.Clinic_ID,
			"name":         req.Name,
			"surname":      req.Surname,
			"birthday":     req.Birthday,
			"gender":       req.Gender,
			"education":    req.Education,
			"certificates": req.Certificates,
		},
	).Where(r.db.Sq.Equal("doctor_id", req.Doctor_ID))

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

func (r doctorRepo) Delete(ctx context.Context, params map[string]string) error {
	queryBuilder := r.db.Sq.Builder.Delete(r.table)
	for k, v := range params {
		switch k {
		case "doctor_id":
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
