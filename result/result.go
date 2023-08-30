package result

import (
	"encoding/json"
	"github.com/qwerty-iot/tox"
)

type Result struct {
	Context  string `json:"@odata.context"`
	NextLink string `json:"@odata.nextLink,omitempty"`
	Value    []any  `json:"value"`
}

func New(context string) *Result {
	return &Result{Context: context}
}

func (r *Result) SetNextLink(url string) {
	r.NextLink = url
}

func (r *Result) Add(obj tox.Object) {
	r.Value = append(r.Value, obj)
}

func (r *Result) Append(arr []any) {
	r.Value = append(r.Value, arr...)
}

func (r *Result) ToJson() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}
