package appointment

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableAppointment = "appointment"
)

type appointmentRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewAppointmentRepo(db *postgres.PostgresDB) Repository {
	return &appointmentRepo{
		table: tableAppointment,
		db:    db,
	}
}

func (r appointmentRepo) Get(ctx context.Context, params map[string]string) (*entity.Appointment, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"doctor_id",
		"patient_id",
		"appointment_at",
		"appointment_reason",
		"status",
		"emr_id",
		"created_at",
		"updated_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "doctor_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "patient_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "appointment_at":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "appointment_reason":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "status":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "emr_id":
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

	var appointment entity.Appointment
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&appointment.Doctor_ID,
		&appointment.Patient_ID,
		&appointment.ID,
		&appointment.GUID,
		&appointment.AppointmentAt,
		&appointment.AppointmentReason,
		&appointment.Price,
		&appointment.Status,
		&appointment.EmrID,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &appointment, nil
}

func (r appointmentRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Appointment, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"doctor_id",
		"patient_id",
		"appointment_at",
		"appointment_reason",
		"status",
		"emr_id",
		"created_at",
		"updated_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "patient_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
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

	var appointments []*entity.Appointment
	for rows.Next() {
		var appointment entity.Appointment
		if err := rows.Scan(
			&appointment.Doctor_ID,
			&appointment.Patient_ID,
			&appointment.ID,
			&appointment.GUID,
			&appointment.AppointmentAt,
			&appointment.AppointmentReason,
			&appointment.Price,
			&appointment.Status,
			&appointment.EmrID,
			&appointment.CreatedAt,
			&appointment.UpdatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		appointments = append(appointments, &appointment)
	}

	return appointments, nil
}

func (r appointmentRepo) Create(ctx context.Context, req *entity.Appointment) error {
	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"doctor_id":          req.Doctor_ID,
			"patient_id":         req.Patient_ID,
			"id":                 req.ID,
			"guid":               req.GUID,
			"appointment_at":     req.AppointmentAt,
			"appointment_reason": req.AppointmentReason,
			"status":             req.Status,
			"emr_id":             req.EmrID,
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

func (r appointmentRepo) Update(ctx context.Context, req *entity.Appointment) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"doctor_id":          req.Doctor_ID,
			"patient_id":         req.Patient_ID,
			"id":                 req.ID,
			"guid":               req.GUID,
			"appointment_at":     req.AppointmentAt,
			"appointment_reason": req.AppointmentReason,
			"status":             req.Status,
			"emr_id":             req.EmrID,
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

func (r appointmentRepo) Delete(ctx context.Context, params map[string]string) error {
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
