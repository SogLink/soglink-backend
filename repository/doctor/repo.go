package doctor

import (
	"context"

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
	return nil, nil
}

func (r doctorRepo) Create(ctx context.Context, d *entity.Doctor) error {
	return nil
}

func (r doctorRepo) Update(ctx context.Context, d *entity.Doctor) error {
	return nil
}

func (r doctorRepo) Delete(ctx context.Context, params map[string]string) error {
	return nil
}
