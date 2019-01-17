package service

import (
	"context"
	omegadbconfig "github.com/angadsharma1016/omega_dbconfig"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ParticipantsService) ParticipantsService

type loggingMiddleware struct {
	logger log.Logger
	next   ParticipantsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ParticipantsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ParticipantsService) ParticipantsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateP(ctx context.Context, s omegadbconfig.Participant) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateP", "s", s, "rs", rs, "err", err)
	}()
	return l.next.CreateP(ctx, s)
}
func (l loggingMiddleware) ReadP(ctx context.Context, e omegadbconfig.Query) (rs []events.Participant, err error) {
	defer func() {
		l.logger.Log("method", "ReadP", "e", e, "rs", rs, "err", err)
	}()
	return l.next.ReadP(ctx, e)
}
func (l loggingMiddleware) UpdateP(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "UpdateP", "e", e, "rs", rs, "err", err)
	}()
	return l.next.UpdateP(ctx, e)
}
func (l loggingMiddleware) DeleteP(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "DeleteP", "e", e, "rs", rs, "err", err)
	}()
	return l.next.DeleteP(ctx, e)
}
func (l loggingMiddleware) CreateFeedback(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "CreateFeedback", "e", e, "rs", rs, "err", err)
	}()
	return l.next.CreateFeedback(ctx, e)
}
func (l loggingMiddleware) ReadFeedback(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "ReadFeedback", "e", e, "rs", rs, "err", err)
	}()
	return l.next.ReadFeedback(ctx, e)
}
func (l loggingMiddleware) PostAttendance(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "PostAttendance", "e", e, "rs", rs, "err", err)
	}()
	return l.next.PostAttendance(ctx, e)
}
func (l loggingMiddleware) ReadAttendance(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "ReadAttendance", "e", e, "rs", rs, "err", err)
	}()
	return l.next.ReadAttendance(ctx, e)
}
