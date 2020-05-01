package api

import (
	elastic "github.com/jaydenwen123/go-es-client"
)

type AllApi struct {
	client *elastic.Client
	path   string
}

//func AllIndices(client *elastic.Client) (*AllApi,error) {
//	a := &AllApi{
//		client: client,
//		path:   "/_all",
//	}
//	_, bdata, err := a.client.GetIndex(context.Background(), a.path, nil)
//	return
//}
