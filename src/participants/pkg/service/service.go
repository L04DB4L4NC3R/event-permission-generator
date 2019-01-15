package service

import (
	"context"

	events "github.com/angadsharma1016/omega_dbconfig"
)

// ParticipantsService describes the service.
type ParticipantsService interface {
	// Add your methods here
	CreateP(ctx context.Context, s events.Participant) (rs string, err error)
	ReadP(ctx context.Context, e events.Query) (rs []events.Participant, err error)
	UpdateP(ctx context.Context, e events.Query) (rs string, err error)
	DeleteP(ctx context.Context, e events.Query) (rs string, err error)
	CreateFeedback(ctx context.Context, e events.Query) (rs string, err error)
	ReadFeedback(ctx context.Context, e events.Query) (rs string, err error)
	PostAttendance(ctx context.Context, e events.Query) (rs string, err error)
	ReadAttendance(ctx context.Context, e events.Query) (rs string, err error)
}

type basicParticipantsService struct{}

func (b *basicParticipantsService) CreateP(ctx context.Context, s events.Participant) (rs string, err error) {
	// TODO implement the business logic of CreateP
	return rs, err
}
func (b *basicParticipantsService) ReadP(ctx context.Context, e events.Query) (rs []events.Participant, err error) {
	// TODO implement the business logic of ReadP
	return rs, err
}
func (b *basicParticipantsService) UpdateP(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of UpdateP
	return rs, err
}
func (b *basicParticipantsService) DeleteP(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of DeleteP
	return rs, err
}
func (b *basicParticipantsService) CreateFeedback(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of CreateFeedback
	return rs, err
}
func (b *basicParticipantsService) ReadFeedback(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of ReadFeedback
	return rs, err
}
func (b *basicParticipantsService) PostAttendance(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of PostAttendance
	return rs, err
}
func (b *basicParticipantsService) ReadAttendance(ctx context.Context, e events.Query) (rs string, err error) {
	// TODO implement the business logic of ReadAttendance
	return rs, err
}

// NewBasicParticipantsService returns a naive, stateless implementation of ParticipantsService.
func NewBasicParticipantsService() ParticipantsService {
	return &basicParticipantsService{}
}

// New returns a ParticipantsService with all of the expected middleware wired in.
func New(middleware []Middleware) ParticipantsService {
	var svc ParticipantsService = NewBasicParticipantsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
