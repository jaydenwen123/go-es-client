package api

import (
	elastic "github.com/jaydenwen123/go-es-client"
)

type AllApi struct {
	client *elastic.Client
	path   string
}

//func All(client *elastic.Client) (*AllApi,error) {
//	a := &AllApi{
//		client: client,
//		path:   "/_all",
//	}
//	_, bdata, err := a.client.Get(context.Background(), a.path, nil)
//	return
//}
