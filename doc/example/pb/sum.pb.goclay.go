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
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/utrack/clay/transport"
	"github.com/utrack/clay/transport/httpruntime"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ chi.Router
var _ runtime.Marshaler

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
func (d *SummatorDesc) SwaggerDef() []byte {
	return _swaggerDef_sum_proto
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *SummatorDesc) RegisterHTTP(mux transport.Router) {

	// Handlers for Sum

	mux.MethodFunc("GET", pattern_goclay_Summator_Sum_0, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req SumRequest
		err := unmarshaler_goclay_Summator_Sum_0(r, &req)
		if err != nil {
			httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't parse request"))
			return
		}

		ret, err := d.svc.Sum(r.Context(), &req)
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

}

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
    "/v1/example/sum/{b}": {
      "get": {
        "operationId": "Sum",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/sumpbSumResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "b",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "a",
            "description": "A is the number we're adding to. Can't be zero for the sake of example.",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Summator"
        ]
      }
    }
  },
  "definitions": {
    "sumpbSumRequest": {
      "type": "object",
      "properties": {
        "a": {
          "type": "string",
          "format": "int64",
          "description": "A is the number we're adding to. Can't be zero for the sake of example."
        },
        "b": {
          "type": "string",
          "format": "int64",
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

	//TODO path params aren't supported atm
	path := pattern_goclay_Summator_Sum_0_builder(in.B)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)
	err := m.Marshal(buf, in)
	if err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("GET", c.host+path, buf)
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

var (
	pattern_goclay_Summator_Sum_0         = "/v1/example/sum/{b}"
	pattern_goclay_Summator_Sum_0_builder = func(
		b int64,
	) string {
		return fmt.Sprintf("/v1/example/sum/%v", b)
	}
	unmarshaler_goclay_Summator_Sum_0 = func(r *http.Request, req *SumRequest) error {

		var err error

		rctx := chi.RouteContext(r.Context())
		if rctx == nil {
			panic("Only chi router is supported for GETs atm")
		}
		for pos, k := range rctx.URLParams.Keys {
			runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos])
		}

		return err
	}
)
