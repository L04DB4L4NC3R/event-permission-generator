package http

import (
	endpoint "angadsharma1016/omega/src/participants/pkg/endpoint"
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeCreatePHandler creates the handler logic
func makeCreatePHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-p", http1.NewServer(endpoints.CreatePEndpoint, decodeCreatePRequest, encodeCreatePResponse, options...))
}

// decodeCreatePResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreatePRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreatePRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreatePResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreatePResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadPHandler creates the handler logic
func makeReadPHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/read-p", http1.NewServer(endpoints.ReadPEndpoint, decodeReadPRequest, encodeReadPResponse, options...))
}

// decodeReadPResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadPRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadPResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdatePHandler creates the handler logic
func makeUpdatePHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-p", http1.NewServer(endpoints.UpdatePEndpoint, decodeUpdatePRequest, encodeUpdatePResponse, options...))
}

// decodeUpdatePResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdatePRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdatePRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdatePResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdatePResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeletePHandler creates the handler logic
func makeDeletePHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-p", http1.NewServer(endpoints.DeletePEndpoint, decodeDeletePRequest, encodeDeletePResponse, options...))
}

// decodeDeletePResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeletePRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeletePRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeletePResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeletePResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateFeedbackHandler creates the handler logic
func makeCreateFeedbackHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-feedback", http1.NewServer(endpoints.CreateFeedbackEndpoint, decodeCreateFeedbackRequest, encodeCreateFeedbackResponse, options...))
}

// decodeCreateFeedbackResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateFeedbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateFeedbackRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateFeedbackResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateFeedbackResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadFeedbackHandler creates the handler logic
func makeReadFeedbackHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/read-feedback", http1.NewServer(endpoints.ReadFeedbackEndpoint, decodeReadFeedbackRequest, encodeReadFeedbackResponse, options...))
}

// decodeReadFeedbackResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadFeedbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadFeedbackRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadFeedbackResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadFeedbackResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostAttendanceHandler creates the handler logic
func makePostAttendanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/post-attendance", http1.NewServer(endpoints.PostAttendanceEndpoint, decodePostAttendanceRequest, encodePostAttendanceResponse, options...))
}

// decodePostAttendanceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostAttendanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PostAttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostAttendanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostAttendanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeReadAttendanceHandler creates the handler logic
func makeReadAttendanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/read-attendance", http1.NewServer(endpoints.ReadAttendanceEndpoint, decodeReadAttendanceRequest, encodeReadAttendanceResponse, options...))
}

// decodeReadAttendanceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeReadAttendanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ReadAttendanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeReadAttendanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeReadAttendanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
