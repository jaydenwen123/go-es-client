package api

import (
	"context"
	elastic "github.com/jaydenwen123/go-es-client"
	"net/url"
)

//Do do接口
type Do interface {
	Do(ctx context.Context) error
}

type Paramer interface {
	Param() string
	Path() string
	Client() *elastic.Client
}

//Eraser 察除
type Eraser interface {
	Erase() (path string, param url.Values)
}

//baseCtx 基础参数
type baseCtx struct {
	client *elastic.Client
	param  url.Values
	path   string
}

func (b *baseCtx) Param() string {
	if b == nil {
		return ""
	}
	return b.param.Encode()
}

func (b *baseCtx) Path() string {
	if b == nil {
		return ""
	}
	return b.path
}

func (b *baseCtx) Client() *elastic.Client {
	if b != nil {
		return b.client
	}
	return nil
}
