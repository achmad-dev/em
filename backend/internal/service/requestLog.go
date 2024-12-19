package service

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"context"

	"github.com/achmad/em/backend/internal/domain"
	"github.com/achmad/em/backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type RequestLogService interface {
	InsertRequestLog(ctx context.Context, userID, message, status string) error
}

type requestLogServiceImpl struct {
	requestLogRepo repository.RequestLogRepo
	log            *logrus.Logger
}

// InsertRequestLog implements RequestLogService.
func (r *requestLogServiceImpl) InsertRequestLog(ctx context.Context, userID, message, status string) error {
	requestLog := domain.RequestLog{
		UserID:  userID,
		Message: message,
		Status:  status,
	}
	err := r.requestLogRepo.InsertRequestLog(ctx, requestLog)
	if err != nil {
		r.log.Error(err)
		return err
	}

	return nil
}

func NewRequestLogService(requestLogRepo repository.RequestLogRepo, log *logrus.Logger) RequestLogService {
	return &requestLogServiceImpl{
		requestLogRepo: requestLogRepo,
		log:            log,
	}
}
