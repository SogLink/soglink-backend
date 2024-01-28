package doctorspeciality

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableDoctorSpecialty = "doctor_specialty"
)

type doctor_specialtyRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewDoctor_specialtyRepo(db *postgres.PostgresDB) Repository {
	return &doctor_specialtyRepo{
		table: tableDoctorSpecialty,
		db:    db,
	}
}

func (r doctor_specialtyRepo) Get(ctx context.Context, params map[string]string) (*entity.Doctor_specialty, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"doctor_id",
		"specialty_id",
		"price",
		"created_at",
		"updated_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "doctor_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "specialty_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var doctor_specialty entity.Doctor_specialty
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&doctor_specialty.DoctorID,
		&doctor_specialty.SpecialtyID,
		&doctor_specialty.Price,
		&doctor_specialty.CreatedAt,
		&doctor_specialty.UpdatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &doctor_specialty, nil
}

func (r doctor_specialtyRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Doctor_specialty, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"doctor_id",
		"specialty_id",
		"price",
		"created_at",
		"updated_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "specialty_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
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

	var doctor_specialties []*entity.Doctor_specialty
	for rows.Next() {
		var doctor_specialty entity.Doctor_specialty
		if err := rows.Scan(
			&doctor_specialty.DoctorID,
			&doctor_specialty.SpecialtyID,
			&doctor_specialty.Price,
			&doctor_specialty.CreatedAt,
			&doctor_specialty.UpdatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		doctor_specialties = append(doctor_specialties, &doctor_specialty)
	}

	return doctor_specialties, nil
}

func (r doctor_specialtyRepo) Create(ctx context.Context, req *entity.Doctor_specialty) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"doctor_id":    req.DoctorID,
			"specialty_id": req.SpecialtyID,
			"price":        req.Price,
			"created_at":   req.CreatedAt,
			"updated_at":   req.UpdatedAt,
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

func (r doctor_specialtyRepo) Update(ctx context.Context, req *entity.Doctor_specialty) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"specialty_id": req.SpecialtyID,
			"price":        req.Price,
			"created_at":   req.CreatedAt,
			"updated_at":   req.UpdatedAt,
		},
	).Where(r.db.Sq.Equal("doctor_id", req.DoctorID))

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

func (r doctor_specialtyRepo) Delete(ctx context.Context, params map[string]string) error {
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
