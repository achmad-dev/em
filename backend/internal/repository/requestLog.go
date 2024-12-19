package repository

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

// we can use this for checking logs of the request
// to check if anything is wrong
import (
	"context"

	"github.com/achmad/em/backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

// repository for request log
type RequestLogRepo interface {
	InsertRequestLog(ctx context.Context, request domain.RequestLog) error
}

// implementation of RequestLogRepo
type RequestLogRepoImpl struct {
	sqlx *sqlx.DB
}

// InsertRequestLog implements RequestLogRepo.
func (r *RequestLogRepoImpl) InsertRequestLog(ctx context.Context, request domain.RequestLog) error {
	query := `INSERT INTO request_logs (user_id, message, status) VALUES ($1, $2, $3)`
	_, err := r.sqlx.ExecContext(ctx, query, request.UserID, request.Message, request.Status)
	if err != nil {
		return err
	}
	return nil
}

func NewRequestLogRepository(sqlx *sqlx.DB) RequestLogRepo {
	return &RequestLogRepoImpl{sqlx: sqlx}
}
