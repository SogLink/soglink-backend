package file

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableFile = "file"
)

type fileRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewFileRepo(db *postgres.PostgresDB) Repository {
	return &fileRepo{
		table: tableFile,
		db:    db,
	}
}

func (r fileRepo) Get(ctx context.Context, params map[string]string) (*entity.File, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"path",
		"created_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "guid":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "path":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "created_at":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var file entity.File
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&file.ID,
		&file.GUID,
		&file.Path,
		&file.CreatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &file, nil
}

func (r fileRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.File, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"id",
		"guid",
		"path",
		"created_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "path":
			queryBuilder = queryBuilder.Where("path ILIKE '%'||?||'%'", v)
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

	var files []*entity.File
	for rows.Next() {
		var file entity.File
		if err := rows.Scan(
			&file.ID,
			&file.GUID,
			&file.Path,
			&file.CreatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		files = append(files, &file)
	}

	return files, nil
}

func (r fileRepo) Create(ctx context.Context, req *entity.File) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"id":         req.ID,
			"guid":       req.GUID,
			"path":       req.Path,
			"created_at": req.CreatedAt,
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

func (r fileRepo) Update(ctx context.Context, req *entity.File) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"guid":       req.GUID,
			"path":       req.Path,
			"created_at": req.CreatedAt,
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

func (r fileRepo) Delete(ctx context.Context, params map[string]string) error {
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
