package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
	"net/http"
)

type WSRequest struct {
	Request   *http.Request `json:"-"`
	Tag       string
	RequestId string
	Meta      *Meta
	Methods   []string
	URL       string
	Data      interface{}
	Ignore    bool
}

func NewWSRequest(r *http.Request) *WSRequest {
	return &WSRequest{Request: r}
}

func (s *WSRequest) Valid() bool {
	return len(s.Methods) != 0 && len(s.URL) != 0
}

func (s *WSRequest) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func BytesToWSRequest(data []byte) (*WSRequest, error) {
	var r WSRequest
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
