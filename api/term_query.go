package api

import (
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/jaydenwen123/go-es-client"
	"github.com/tidwall/gjson"
	"strings"
)

type Term struct {
	*baseCtx
	indices string
}

type termCondition struct {
	query *termCond `json:"query"`
}

type termCond struct {
	term  map[string]interface{} `json:"term,omitempty"`
	terms map[string]interface{} `json:"terms,omitempty"`
}

type TermQuery struct {
	term      *Term
	termCond  *termCondition `json:"term_cond"`
	termsCond *termCondition `json:"terms_cond"`
}

func (tq *TermQuery) Query(ctx context.Context) (*QueryDocInfo, error) {
	var cond interface{}
	var path string
	var docInfo QueryDocInfo
	if tq.termCond != nil {
		cond = tq.termCond
	} else {
		cond = tq.termsCond
	}
	path = tq.term.path + tq.term.param.Encode()
	_, bdata, err := tq.term.client.Get(ctx, path, cond)
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(bdata, "error").Exists() {
		return nil, fmt.Errorf("error info:\n%s", bdata)
	}
	err = json.Unmarshal(bdata, &docInfo)
	return &docInfo, err
}

func TermQueryAPI(client *elastic.Client) *Term {
	return TermQueryApi(client)
}

func TermQueryApi(client *elastic.Client) *Term {
	return &Term{
		baseCtx: &baseCtx{
			client: client,
			param:  make(map[string][]string),
			path:   "/_search",
		},
	}
}

func (t *Term) Indices(indices ...string) *Term {
	t.indices = strings.Join(indices, ",")
	t.path = "/" + t.indices + t.path
	return t
}

//term
func (t *Term) Term(field string, value interface{}) *TermQuery {
	return &TermQuery{
		term: t,
		termCond: &termCondition{
			query: &termCond{
				term: map[string]interface{}{
					field: value,
				},
			},
		},
	}
}

//terms
func (t *Term) Terms(field string, values ...interface{}) *TermQuery {
	return &TermQuery{
		term: t,
		termCond: &termCondition{
			query: &termCond{
				terms: map[string]interface{}{
					field: values,
				},
			},
		},
	}
}
