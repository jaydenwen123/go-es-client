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
