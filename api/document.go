package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	elastic "github.com/jaydenwen123/go-es-client"

	"github.com/tidwall/gjson"
)

const (
	doc    = "_doc"
	update = "_update"
)

var (
	op_DocGet     = opDocGet()
	op_DocPost    = opDocPost()
	op_DocPut     = opDocPut()
	op_UpdatePost = opUpdatePost()
	op_UpdatePut  = opUpdatePut()
)

type execOp struct {
	opType   string
	opMethod string
}

func op(opType, opMethod string) *execOp {
	return &execOp{
		opType:   opType,
		opMethod: opMethod,
	}
}

func opDocPut() *execOp {
	return op(doc, "PUT")
}

func opDocPost() *execOp {
	return op(doc, "POST")
}

func opUpdatePost() *execOp {
	return op(update, "POST")
}

func opUpdatePut() *execOp {
	return op(update, "PUT")
}

func opDocGet() *execOp {
	return op(doc, "GET")
}

type Document struct {
	*baseCtx
	index string
	//doc   interface{}
	param string
}

func DocApi(client *elastic.Client, index string) *Document {
	return &Document{
		baseCtx: &baseCtx{
			client: client,
			param:  make(map[string][]string),
			path:   "",
		},
		index: index,
	}
}

func DocAPI(client *elastic.Client, index string) *Document {
	return DocApi(client, index)
}

func (d *Document) Index() string {
	return d.index
}

func (d *Document) WithParam(param string) *Document {
	d.param = param
	return d
}

//增,put指定id、post不指定id
func (d *Document) AddOrUpdate(ctx context.Context, docBody interface{}) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_DocPost, nil, docBody)

}

func (d *Document) AddOrUpdateBodyJson(ctx context.Context, docBodyJson string) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_DocPost, nil, docBodyJson)

}

//增,put指定id、post不指定id
func (d *Document) AddOrUpdateWithId(ctx context.Context, id interface{}, docBody interface{}) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_DocPut, id, docBody)

}

func (d *Document) AddOrUpdateBodyJsonWithId(ctx context.Context, id interface{}, docBodyJson string) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_DocPut, id, docBodyJson)
}

func (d *Document) UpdateFieldWithId(ctx context.Context, id interface{}, bodyJson string) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_UpdatePut, id, bodyJson)
}

func (d *Document) UpdateField(ctx context.Context, bodyJson string) (docId string, err error) {
	return d.execIgnoreResp(ctx, op_UpdatePost, nil, bodyJson)
}

func (d *Document) execIgnoreResp(ctx context.Context, op *execOp, id interface{}, body interface{}) (docId string, err error) {
	_, docId, err = d.executeOp(ctx, op, id, body)
	return docId, err
}

func (d *Document) executeOp(ctx context.Context, op *execOp, id interface{}, body interface{}) ([]byte, string, error) {
	d.buildPath(op.opType, id)
	resp, bdata, err := d.client.Do(ctx, op.opMethod, d.path, body)
	if err != nil {
		return bdata, "", err
	}
	docId, err := d.wrapResp(resp, bdata)
	return bdata, docId, err
}

func (d *Document) buildPath(opType string, id interface{}) string {
	d.path = "/" + d.index + "/" + opType
	if id != nil {
		d.path += "/" + fmt.Sprintf("%s", id)
	}
	if d.param != "" {
		d.path = d.path + "?" + d.param
	}
	return d.path
}

func (d *Document) wrapResp(resp *http.Response, bdata []byte, ) (docId string, err error) {
	if resp.StatusCode > 300 || gjson.GetBytes(bdata, "error").Exists() {
		return "", fmt.Errorf("%s", bdata)
	}
	idNode := gjson.GetBytes(bdata, "_id")
	if !idNode.Exists() {
		return "", fmt.Errorf("the resp info does not contains id")
	}
	return idNode.String(), nil
}

//删
func (d *Document) DeleteWithId(ctx context.Context, id interface{}) (docId string, err error) {
	path := fmt.Sprintf("/%s/%s/%s", d.index, doc, id)
	resp, bdata, err := d.client.Delete(ctx, path, nil)
	if err != nil {
		return "", err
	}
	return d.wrapResp(resp, bdata)
}

//查
func (d *Document) GetWithId(ctx context.Context, id interface{}) (*DocumentInfo, error) {
	bdata, _, err := d.executeOp(ctx, op_DocGet, id, nil)
	if err != nil {
		return nil, err
	}
	document := new(DocumentInfo)
	err = json.Unmarshal(bdata, document)
	if err != nil {
		return document, err
	}
	return document, nil
}

type DocumentInfo struct {
	Index       string          `json:"_index,omitempty"`
	Type        string          `json:"_type,omitempty"`
	ID          string          `json:"_id,omitempty"`
	Version     int             `json:"_version,omitempty"`
	SeqNo       int             `json:"_seq_no,omitempty"`
	PrimaryTerm int             `json:"_primary_term,omitempty"`
	Found       bool            `json:"found,omitempty"`
	Source      json.RawMessage `json:"_source,omitempty"`
}
