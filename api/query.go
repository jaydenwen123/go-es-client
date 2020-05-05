package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	elastic "github.com/jaydenwen123/go-es-client"
	"github.com/jaydenwen123/go-util"
	"github.com/tidwall/gjson"
)

type QueryInfo struct {
	Took     int     `json:"took,omitempty"`
	TimedOut bool    `json:"timed_out,omitempty"`
	Shards   *Shards `json:"_shards,omitempty"`
	Hits     *Hits   `json:"hits,omitempty"`
}

type Shards struct {
	Total      int `json:"total,omitempty"`
	Successful int `json:"successful,omitempty"`
	Skipped    int `json:"skipped,omitempty"`
	Failed     int `json:"failed,omitempty"`
}

type Total struct {
	Value    int    `json:"value,omitempty"`
	Relation string `json:"relation,omitempty"`
}

type Hits struct {
	Total    *Total      `json:"total,omitempty"`
	MaxScore float64     `json:"max_score,omitempty"`
	Hits     []*InnerHit `json:"hits,omitempty"`
}

type InnerHit struct {
	Index     string          `json:"_index,omitempty"`
	Type      string          `json:"_type,omitempty"`
	ID        string          `json:"_id,omitempty"`
	Score     float64         `json:"_score,omitempty"`
	Source    json.RawMessage `json:"_source,omitempty"`
	Highlight *HighLight      `json:"highlight,omitempty"`
}

type HighLight struct {
	Title []string `json:"title,omitempty"`
}

type Query struct {
	*baseCtx
	indices string
}

type queryCondition struct {
	Query *queryCond `json:"query"`
}

func (c *queryCondition) String() string {
	return util.Obj2JsonStrIndent(c, "", "  ")
}

//Json 返回query 的json数据
func (c *queryCondition) Json() string {
	return c.String()
}

type RangeCond struct {
	Gt       interface{} `json:"gt,omitempty"`
	Lt       interface{} `json:"lt,omitempty"`
	Gte      interface{} `json:"gte,omitempty"`
	Lte      interface{} `json:"lte,omitempty"`
	Format   string      `json:"format,omitempty"`
	Relation string      `json:"relation,omitempty"`
	TimeZone string      `json:"time_zone,omitempty"`
	Boost    float32     `json:"boost,omitempty"`
}

type RegexpCond struct {
	Value                 string `json:"value,omitempty"`
	Flags                 string `json:"flags,omitempty"`
	MaxDeterminizedStates int    `json:"max_determinized_states,omitempty"`
	Rewrite               string `json:"rewrite,omitempty"`
}

type WildcardCond struct {
	Value   string  `json:"value,omitempty"`
	Rewrite string  `json:"rewrite,omitempty"`
	Boost   float32 `json:"boost,omitempty"`
}

type queryCond struct {
	Term  map[string]interface{} `json:"term,omitempty"`
	Terms map[string]interface{} `json:"terms,omitempty"`

	Ids map[string][]string `json:"ids,omitempty"`

	Exists   map[string]interface{}   `json:"exists,omitempty"`
	Prefix   map[string]interface{}   `json:"prefix,omitempty"`
	Range    map[string]*RangeCond    `json:"range,omitempty"`
	Regexp   map[string]*RegexpCond   `json:"regexp,omitempty"`
	Wildcard map[string]*WildcardCond `json:"wildcard,omitempty"`

	//match 相关查询
	Match             map[string]interface{} `json:"match,omitempty"`
	MatchAll          interface{}            `json:"match_all,omitempty"`
	MatchNone         map[string]interface{} `json:"match_none,omitempty"`
	MatchPhrase       map[string]interface{} `json:"match_phrase,omitempty"`
	MatchPhrasePrefix map[string]interface{} `json:"match_phrase_prefix,omitempty"`
	MultiMatch        map[string]interface{} `json:"multi_match,omitempty"`

	QueryString map[string]interface{} `json:"query_string,omitempty"`
}

type QueryProxy struct {
	term      *Query          `json:"-"`
	QueryCond *queryCondition `json:"query_cond"`
	//TermsCond *queryCondition `json:"terms_cond"`
}

func (tq *QueryProxy) String() string {
	cond := tq.QueryCond
	return cond.String()
}

//Query query
func (tq *QueryProxy) Query(ctx context.Context) (*QueryInfo, error) {
	var cond = tq.QueryCond
	var path string
	var docInfo QueryInfo
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

func QueryAPI(client *elastic.Client) *Query {
	return QueryApi(client)
}

func QueryApi(client *elastic.Client) *Query {
	return &Query{
		baseCtx: &baseCtx{
			client: client,
			param:  make(map[string][]string),
			path:   "/_search",
		},
	}
}

func (t *Query) Indices(indices ...string) *Query {
	t.indices = strings.Join(indices, ",")
	t.path = "/" + t.indices + t.path
	return t
}
