package attachment

import (
	"context"
	"fmt"

	"github.com/SogLink/soglink-backend/entity"
	postgres "github.com/SogLink/soglink-backend/pkg/database"
)

var (
	tableAttachment = "attachment"
)

type attachmentRepo struct {
	table string
	db    *postgres.PostgresDB
}

func NewAttachmentRepo(db *postgres.PostgresDB) Repository {
	return &attachmentRepo{
		table: tableAttachment,
		db:    db,
	}
}

func (r attachmentRepo) Get(ctx context.Context, params map[string]string) (*entity.Attachment, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"emr_id",
		"file_id",
		"created_at",
	).From(r.table)

	for k, v := range params {
		switch k {
		case "emr_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "file_id":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		case "created_at":
			queryBuilder = queryBuilder.Where(r.db.Sq.Equal(k, v))
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, r.db.ErrSQLBuild(err, r.table+" Get")
	}

	var attachment entity.Attachment
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&attachment.Emr_ID,
		&attachment.File_ID,
		&attachment.CreatedAt,
	)
	if err != nil {
		return nil, r.db.Error(err)
	}

	return &attachment, nil
}

func (r attachmentRepo) List(ctx context.Context, limit, offset uint64, params map[string]string) ([]*entity.Attachment, error) {
	queryBuilder := r.db.Sq.Builder.Select(
		"emr_id",
		"file_id",
		"created_at",
	).From(r.table).OrderBy("created_at asc")

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	for k, v := range params {
		switch k {
		case "emr_id":
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

	var attachments []*entity.Attachment
	for rows.Next() {
		var attachment entity.Attachment
		if err := rows.Scan(
			&attachment.Emr_ID,
			&attachment.File_ID,
			&attachment.CreatedAt,
		); err != nil {
			return nil, r.db.Error(err)
		}

		attachments = append(attachments, &attachment)
	}

	return attachments, nil
}

func (r attachmentRepo) Create(ctx context.Context, req *entity.Attachment) error {

	queryBuilder := r.db.Sq.Builder.Insert(r.table).SetMap(
		map[string]interface{}{
			"emr_id":     req.Emr_ID,
			"file_id":    req.File_ID,
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

func (r attachmentRepo) Update(ctx context.Context, req *entity.Attachment) error {
	queryBuilder := r.db.Sq.Builder.Update(r.table).SetMap(
		map[string]interface{}{
			"file_id":    req.File_ID,
			"created_at": req.CreatedAt,
		},
	).Where(r.db.Sq.Equal("emr_id", req.Emr_ID))

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

func (r attachmentRepo) Delete(ctx context.Context, params map[string]string) error {
	queryBuilder := r.db.Sq.Builder.Delete(r.table)
	for k, v := range params {
		switch k {
		case "emr_id":
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
