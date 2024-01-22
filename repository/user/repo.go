package user

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableUser = "user"
)

type userRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewUserRepo(db *postgres.PostgresDB) Repository {
	return &userRepo{
		table: tableUser,
		db:    db,
	}
}

func (r userRepo) Get(ctx context.Context, params map[string]string) (*entity.User, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"username",
		"email",
		"phone",
		"password",
		"created_at",
		"updated_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "username":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "email":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "phone":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var user entity.User
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.GUID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &user, nil
}

func (r userRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.User, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"username",
		"email",
		"phone",
		"password",
		"created_at",
		"updated_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "username":
			queryBuilder = queryBuilder.Where("username ILIKE '%'||?||'%'", v)
		case "email":
			queryBuilder = queryBuilder.Where("email ILIKE '%'||?||'%'", v)
		case "phone":
			queryBuilder = queryBuilder.Where("phone ILIKE '%'||?||'%'", v)
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

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(
			&user.ID,
			&user.GUID,
			&user.Username,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r userRepo) Create(ctx context.Context, req *entity.User) error {
	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"guid":       req.GUID,
			"username":   req.Username,
			"email":      req.Email,
			"phone":      req.Phone,
			"password":   req.Password,
			"created_at": req.CreatedAt,
			"updated_at": req.UpdatedAt,
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

func (r userRepo) Update(ctx context.Context, req *entity.User) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"username":   req.Username,
			"email":      req.Email,
			"phone":      req.Phone,
			"password":   req.Password,
			"updated_at": req.UpdatedAt,
		},
	).Where(r.db.Sq.Equal("guid", req.GUID))

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

func (r userRepo) Delete(ctx context.Context, params map[string]string) error {
	queryBuilder := r.db.Sq.Builder.Delete(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
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
