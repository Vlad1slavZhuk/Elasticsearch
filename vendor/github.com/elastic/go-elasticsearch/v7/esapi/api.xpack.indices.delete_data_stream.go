// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.10.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newIndicesDeleteDataStreamFunc(t Transport) IndicesDeleteDataStream {
	return func(name []string, o ...func(*IndicesDeleteDataStreamRequest)) (*Response, error) {
		var r = IndicesDeleteDataStreamRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesDeleteDataStream - Deletes a data stream.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/data-streams.html.
//
type IndicesDeleteDataStream func(name []string, o ...func(*IndicesDeleteDataStreamRequest)) (*Response, error)

// IndicesDeleteDataStreamRequest configures the Indices Delete Data Stream API request.
//
type IndicesDeleteDataStreamRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndicesDeleteDataStreamRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_data_stream") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("/")
	path.WriteString("_data_stream")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Name, ","))

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f IndicesDeleteDataStream) WithContext(v context.Context) func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IndicesDeleteDataStream) WithPretty() func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IndicesDeleteDataStream) WithHuman() func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IndicesDeleteDataStream) WithErrorTrace() func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IndicesDeleteDataStream) WithFilterPath(v ...string) func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f IndicesDeleteDataStream) WithHeader(h map[string]string) func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f IndicesDeleteDataStream) WithOpaqueID(s string) func(*IndicesDeleteDataStreamRequest) {
	return func(r *IndicesDeleteDataStreamRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
