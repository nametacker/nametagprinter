package nametagprinter

import (
	"encoding/json"
	"net/http"
)

// Implements the http-problem messafes
// See http://datatracker.ietf.org/doc/draft-ietf-appsawg-http-problem/

type HttpProblemModel struct {
	JsonLDTypedModel
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func NewHttpProblem() (p *HttpProblemModel) {
	p = new(HttpProblemModel)
	p.JsonLDContext = "http://ietf.org/appsawg/http-problem"
	return
}

func HttpProblem(w http.ResponseWriter, statusCode int, message string) {
	p := NewHttpProblem()
	p.Title = message
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/problem+json")
	encoder := json.NewEncoder(w)
	encoder.Encode(p)
}
