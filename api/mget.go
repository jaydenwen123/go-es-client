package api

import (
	"context"
	"encoding/json"

	elastic "github.com/jaydenwen123/go-es-client"
)

type Mget struct {
	*baseCtx
	index string
}

func MgetAPI(c *elastic.Client) *Mget {
	return MgetApi(c)
}

func MgetApi(c *elastic.Client) *Mget {
	return &Mget{
		baseCtx: &baseCtx{client: c},
	}
}

func (m *Mget) Index(index string) *Mget {
	m.index = index
	return m
}


func (m *Mget) BodyJson(ctx context.Context, condJson string) ([]*DocumentInfo, error) {
	return m.mget(ctx, condJson)
}

func (m *Mget) Body(ctx context.Context, cond *MgetCondition) ([]*DocumentInfo, error) {
	return m.mget(ctx, cond)
}

func (m *Mget) mget(ctx context.Context, cond interface{}) ([]*DocumentInfo, error) {
	path := ""
	if m.index != "" {
		path = "/" + m.index
	}
	path += "/_mget"
	_, bdata, err := m.client.Get(ctx, path, cond)
	if err != nil {
		return nil, err
	}
	resp := new(MgetResp)
	err = json.Unmarshal(bdata, resp)
	if err != nil {
		return nil, err
	}
	m.path = path
	return resp.Docs, nil
}
func (m *Mget) resetPath() {
	m.path = ""
}

type DocCondition struct {
	Index string `json:"_index,omitempty"`
	//Type  string `json:"_type"`
	Id string `json:"_id,omitempty"`
}

type MgetCondition struct {
	Conds []*DocCondition `json:"docs,omitempty"`
}

type MgetResp struct {
	Docs []*DocumentInfo `json:"docs,omitempty"`
}
