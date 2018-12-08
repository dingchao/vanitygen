package main

import (
        "bytes"
        "context"
        "encoding/json"
        "io/ioutil"
        "net/http"
	"fmt"

        "github.com/go-kit/kit/endpoint"
)

type getvanityaddressRequest struct {
	Coin string `json:"coin"`
	Prefix string `json:"prefix"`
}

type getvanityaddressResponse struct {
//	Elapsed string `json:"elapsed"`
	Addr string `json:"addr"`
	Key string `json:"key"`
	Err string `json:"err,omitempty"`
}

func makeGetvanityaddressEndpoint(sv Vanitygen) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
                req := request.(getvanityaddressRequest)
                v, key, err := sv.Getvanityaddress(req.Coin, req.Prefix)
                if err != nil {  fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
                        return getvanityaddressResponse{Err:err.Error()}, nil
                }
                return getvanityaddressResponse{v, key, ""}, nil
        }
}

func decodeGetvanityaddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
        var request getvanityaddressRequest
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
                return nil, err
        }
        return request, nil
}

func decodeGetvanityaddressResponse(_ context.Context, r *http.Response) (interface{}, error) {
        var response getvanityaddressResponse
        if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
                return nil, err
        }
        return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
        return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
        var buf bytes.Buffer
        if err := json.NewEncoder(&buf).Encode(request); err != nil {
                return err
        }
        r.Body = ioutil.NopCloser(&buf)
        return nil
}


