package endpoint

import (
	service "angadsharma1016/omega/src/participants/pkg/service"
	"context"
	omegadbconfig "github.com/angadsharma1016/omega_dbconfig"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreatePRequest collects the request parameters for the CreateP method.
type CreatePRequest struct {
	S omegadbconfig.Participant `json:"s"`
}

// CreatePResponse collects the response parameters for the CreateP method.
type CreatePResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreatePEndpoint returns an endpoint that invokes CreateP on the service.
func MakeCreatePEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePRequest)
		rs, err := s.CreateP(ctx, req.S)
		return CreatePResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreatePResponse) Failed() error {
	return r.Err
}

// ReadPRequest collects the request parameters for the ReadP method.
type ReadPRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// ReadPResponse collects the response parameters for the ReadP method.
type ReadPResponse struct {
	Rs  []events.Participant `json:"rs"`
	Err error                `json:"err"`
}

// MakeReadPEndpoint returns an endpoint that invokes ReadP on the service.
func MakeReadPEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadPRequest)
		rs, err := s.ReadP(ctx, req.E)
		return ReadPResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadPResponse) Failed() error {
	return r.Err
}

// UpdatePRequest collects the request parameters for the UpdateP method.
type UpdatePRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// UpdatePResponse collects the response parameters for the UpdateP method.
type UpdatePResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdatePEndpoint returns an endpoint that invokes UpdateP on the service.
func MakeUpdatePEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePRequest)
		rs, err := s.UpdateP(ctx, req.E)
		return UpdatePResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdatePResponse) Failed() error {
	return r.Err
}

// DeletePRequest collects the request parameters for the DeleteP method.
type DeletePRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// DeletePResponse collects the response parameters for the DeleteP method.
type DeletePResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeletePEndpoint returns an endpoint that invokes DeleteP on the service.
func MakeDeletePEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePRequest)
		rs, err := s.DeleteP(ctx, req.E)
		return DeletePResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeletePResponse) Failed() error {
	return r.Err
}

// CreateFeedbackRequest collects the request parameters for the CreateFeedback method.
type CreateFeedbackRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// CreateFeedbackResponse collects the response parameters for the CreateFeedback method.
type CreateFeedbackResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeCreateFeedbackEndpoint returns an endpoint that invokes CreateFeedback on the service.
func MakeCreateFeedbackEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFeedbackRequest)
		rs, err := s.CreateFeedback(ctx, req.E)
		return CreateFeedbackResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateFeedbackResponse) Failed() error {
	return r.Err
}

// ReadFeedbackRequest collects the request parameters for the ReadFeedback method.
type ReadFeedbackRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// ReadFeedbackResponse collects the response parameters for the ReadFeedback method.
type ReadFeedbackResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeReadFeedbackEndpoint returns an endpoint that invokes ReadFeedback on the service.
func MakeReadFeedbackEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadFeedbackRequest)
		rs, err := s.ReadFeedback(ctx, req.E)
		return ReadFeedbackResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadFeedbackResponse) Failed() error {
	return r.Err
}

// PostAttendanceRequest collects the request parameters for the PostAttendance method.
type PostAttendanceRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// PostAttendanceResponse collects the response parameters for the PostAttendance method.
type PostAttendanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakePostAttendanceEndpoint returns an endpoint that invokes PostAttendance on the service.
func MakePostAttendanceEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostAttendanceRequest)
		rs, err := s.PostAttendance(ctx, req.E)
		return PostAttendanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r PostAttendanceResponse) Failed() error {
	return r.Err
}

// ReadAttendanceRequest collects the request parameters for the ReadAttendance method.
type ReadAttendanceRequest struct {
	E omegadbconfig.Query `json:"e"`
}

// ReadAttendanceResponse collects the response parameters for the ReadAttendance method.
type ReadAttendanceResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeReadAttendanceEndpoint returns an endpoint that invokes ReadAttendance on the service.
func MakeReadAttendanceEndpoint(s service.ParticipantsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadAttendanceRequest)
		rs, err := s.ReadAttendance(ctx, req.E)
		return ReadAttendanceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadAttendanceResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateP implements Service. Primarily useful in a client.
func (en Endpoints) CreateP(ctx context.Context, s omegadbconfig.Participant) (rs string, err error) {
	request := CreatePRequest{S: s}
	response, err := en.CreatePEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreatePResponse).Rs, response.(CreatePResponse).Err
}

// ReadP implements Service. Primarily useful in a client.
func (en Endpoints) ReadP(ctx context.Context, e omegadbconfig.Query) (rs []events.Participant, err error) {
	request := ReadPRequest{E: e}
	response, err := en.ReadPEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadPResponse).Rs, response.(ReadPResponse).Err
}

// UpdateP implements Service. Primarily useful in a client.
func (en Endpoints) UpdateP(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := UpdatePRequest{E: e}
	response, err := en.UpdatePEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdatePResponse).Rs, response.(UpdatePResponse).Err
}

// DeleteP implements Service. Primarily useful in a client.
func (en Endpoints) DeleteP(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := DeletePRequest{E: e}
	response, err := en.DeletePEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeletePResponse).Rs, response.(DeletePResponse).Err
}

// CreateFeedback implements Service. Primarily useful in a client.
func (en Endpoints) CreateFeedback(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := CreateFeedbackRequest{E: e}
	response, err := en.CreateFeedbackEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateFeedbackResponse).Rs, response.(CreateFeedbackResponse).Err
}

// ReadFeedback implements Service. Primarily useful in a client.
func (en Endpoints) ReadFeedback(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := ReadFeedbackRequest{E: e}
	response, err := en.ReadFeedbackEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadFeedbackResponse).Rs, response.(ReadFeedbackResponse).Err
}

// PostAttendance implements Service. Primarily useful in a client.
func (en Endpoints) PostAttendance(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := PostAttendanceRequest{E: e}
	response, err := en.PostAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAttendanceResponse).Rs, response.(PostAttendanceResponse).Err
}

// ReadAttendance implements Service. Primarily useful in a client.
func (en Endpoints) ReadAttendance(ctx context.Context, e omegadbconfig.Query) (rs string, err error) {
	request := ReadAttendanceRequest{E: e}
	response, err := en.ReadAttendanceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ReadAttendanceResponse).Rs, response.(ReadAttendanceResponse).Err
}
