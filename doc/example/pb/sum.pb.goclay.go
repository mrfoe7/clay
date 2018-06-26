// Code generated by protoc-gen-goclay
// source: sum.proto
// DO NOT EDIT!

/*
Package sumpb is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay Service interface.
*/
package sumpb

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/utrack/clay/transport"
	"github.com/utrack/clay/transport/httpruntime"
	"github.com/utrack/clay/transport/httpruntime/httpmw"
	"github.com/utrack/clay/transport/swagger"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ = ioutil.Discard
var _ chi.Router
var _ runtime.Marshaler
var _ bytes.Buffer
var _ context.Context
var _ fmt.Formatter
var _ strings.Reader
var _ errors.Frame
var _ httpruntime.Marshaler
var _ http.Handler

// SummatorDesc is a descriptor/registrator for the SummatorServer.
type SummatorDesc struct {
	svc SummatorServer
}

// NewSummatorServiceDesc creates new registrator for the SummatorServer.
func NewSummatorServiceDesc(svc SummatorServer) *SummatorDesc {
	return &SummatorDesc{svc: svc}
}

// RegisterGRPC implements service registrator interface.
func (d *SummatorDesc) RegisterGRPC(s *grpc.Server) {
	RegisterSummatorServer(s, d.svc)
}

// SwaggerDef returns this file's Swagger definition.
func (d *SummatorDesc) SwaggerDef(options ...swagger.Option) (result []byte) {
	if len(options) > 0 {
		var err error
		var s = &spec.Swagger{}
		if err = s.UnmarshalJSON(_swaggerDef_sum_proto); err != nil {
			panic("Bad swagger definition: " + err.Error())
		}
		for _, o := range options {
			o(s)
		}
		if result, err = s.MarshalJSON(); err != nil {
			panic("Failed marshal spec.Swagger definition: " + err.Error())
		}
	} else {
		result = _swaggerDef_sum_proto
	}
	return result
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *SummatorDesc) RegisterHTTP(mux transport.Router) {
	chiMux, isChi := mux.(chi.Router)
	var h http.HandlerFunc

	// Handler for Sum, binding: POST /v1/example/sum/{a}
	h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req, err := unmarshaler_goclay_Summator_Sum_0(r)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't parse request"))
			return
		}

		ret, err := d.svc.Sum(r.Context(), req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "returned from handler"))
			return
		}

		_, outbound := httpruntime.MarshalerForRequest(r)
		w.Header().Set("Content-Type", outbound.ContentType())
		err = outbound.Marshal(w, ret)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
			return
		}
	})

	h = httpmw.DefaultChain(h)

	if isChi {
		chiMux.Method("POST", pattern_goclay_Summator_Sum_0, h)
	} else {
		panic("query URI params supported only for chi.Router")
	}

}

type Summator_httpClient struct {
	c    *http.Client
	host string
}

// NewSummatorHTTPClient creates new HTTP client for SummatorServer.
// Pass addr in format "http://host[:port]".
func NewSummatorHTTPClient(c *http.Client, addr string) SummatorClient {
	if strings.HasSuffix(addr, "/") {
		addr = addr[:len(addr)-1]
	}
	return &Summator_httpClient{c: c, host: addr}
}

func (c *Summator_httpClient) Sum(ctx context.Context, in *SumRequest, _ ...grpc.CallOption) (*SumResponse, error) {
	path := pattern_goclay_Summator_Sum_0_builder(in.A)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err := m.Marshal(buf, in.B); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}

	req.Header.Add("Accept", m.ContentType())

	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := &SumResponse{}
	err = m.Unmarshal(rsp.Body, ret)
	return ret, errors.Wrap(err, "can't unmarshal response")
}

// patterns for Summator
var (
	pattern_goclay_Summator_Sum_0 = "/v1/example/sum/{a}"

	pattern_goclay_Summator_Sum_0_builder = func(
		a int64,
	) string {
		return fmt.Sprintf("/v1/example/sum/%v", a)
	}

	unmarshaler_goclay_Summator_Sum_0_boundParams = map[string]struct{}{
		"b": struct{}{},
		"a": struct{}{},
	}
)

// marshalers for Summator
var (
	unmarshaler_goclay_Summator_Sum_0 = func(r *http.Request) (*SumRequest, error) {
		var req SumRequest

		for k, v := range r.URL.Query() {
			if _, ok := unmarshaler_goclay_Summator_Sum_0_boundParams[strings.ToLower(k)]; ok {
				continue
			}
			if err := errors.Wrap(runtime.PopulateFieldFromPath(&req, k, v[0]), "couldn't populate field from Path"); err != nil {
				return nil, err
			}
		}

		inbound, _ := httpruntime.MarshalerForRequest(r)
		if err := errors.Wrap(inbound.Unmarshal(r.Body, &req.B), "couldn't read request JSON"); err != nil {
			return nil, err
		}

		rctx := chi.RouteContext(r.Context())
		if rctx == nil {
			panic("Only chi router is supported for GETs atm")
		}
		for pos, k := range rctx.URLParams.Keys {
			if err := errors.Wrap(runtime.PopulateFieldFromPath(&req, k, rctx.URLParams.Values[pos]), "couldn't populate field from Path"); err != nil {
				return nil, err
			}
		}

		return &req, nil
	}
)

var _swaggerDef_sum_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "sum.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/sum/{a}": {
      "post": {
        "operationId": "Sum",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/sumpbSumResponse"
            }
          },
          "default": {
            "description": "Error object is returned on error.",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string",
                  "description": "Error string."
                },
                "data": {
                  "type": "object",
                  "description": "Freeform auxilliary data set of string-string."
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "a",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/sumpbNestedB"
            }
          }
        ],
        "tags": [
          "Summator"
        ]
      }
    }
  },
  "definitions": {
    "sumpbNestedB": {
      "type": "object",
      "properties": {
        "b": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "sumpbSumRequest": {
      "type": "object",
      "properties": {
        "a": {
          "type": "string",
          "format": "int64",
          "description": "A is the number we're adding to. Can't be zero for the sake of example."
        },
        "b": {
          "$ref": "#/definitions/sumpbNestedB",
          "description": "B is the number we're adding."
        }
      },
      "description": "SumRequest is a request for Summator service."
    },
    "sumpbSumResponse": {
      "type": "object",
      "properties": {
        "sum": {
          "type": "string",
          "format": "int64"
        },
        "error": {
          "type": "string"
        }
      }
    }
  }
}

`)
