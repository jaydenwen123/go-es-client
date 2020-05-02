package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
	elastic "github.com/jaydenwen123/go-es-client"
)

type Mapping struct {
	*baseCtx
}

//MappingApi 创建mapping api
func MappingApi(client *elastic.Client) *Mapping {
	return &Mapping{
		baseCtx: &baseCtx{
			client: client,
			param:  make(map[string][]string),
		},
	}
}

//MappingAPI 创建mapping api，该api是MappingApi的别名
func MappingAPI(client *elastic.Client) *Mapping {
	return MappingApi(client)
}

//set do real set mapping action.
func (m *Mapping) set(ctx context.Context, bodyJson interface{}) error {
	resp, bdata, err := m.client.Put(ctx, m.path, bodyJson)
	if err != nil {
		logs.Error("exec set mapping op is error:%s", err.Error())
		return err
	}
	m.path = ""
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the error info:%s", bdata)
	}
	return nil
}

//SetMappingWithJson 给index添加 or 更新mapping
func (m *Mapping) SetMappingWithJson(ctx context.Context, index string, bodyJson string) error {
	m.path = "/" + index + "/_mapping"
	return m.set(ctx, bodyJson)
}

//SetWithJson 给index添加 or 更新mapping
func (m *Mapping) SetWithJson(ctx context.Context, index string, bodyJson string) error {
	m.path = "/" + index + "/_mapping"
	return m.set(ctx, bodyJson)
}

//SetMapping 给index添加 or 更新mapping
func (m *Mapping) SetMapping(ctx context.Context, index string, mappingInfo *MappingInfo) error {
	m.path = "/" + index + "/_mapping"
	return m.set(ctx, mappingInfo)
}

//SetMapping 给index添加 or 更新mapping
func (m *Mapping) Set(ctx context.Context, index string, mappingInfo *MappingInfo) error {
	m.path = "/" + index + "/_mapping"
	return m.set(ctx, mappingInfo)
}

//AllMappings 查看所有的mappings
func (m *Mapping) AllMappings(ctx context.Context) (map[string]*MappingInfo, *IndexErrorInfo, error) {
	//方式一 /_all/_mapping
	//方式二 /_mapping
	return m.GetMappings(ctx)
}

//GetMappings 查看多个index 的mapping
func (m *Mapping) GetMappings(ctx context.Context, indices ...string) (map[string]*MappingInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, m.baseCtx, op_mapping, indices...)
	if err != nil {
		return nil, info, err
	}
	multiMap := make(map[string]*MappingInfo)
	err = decodeRespData(bdata, &multiMap)
	return multiMap, info, err
}

//查看index 的mapping
func (m *Mapping) GetMapping(ctx context.Context, index string) (*MappingInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, m.baseCtx, op_mapping, index)
	if err != nil {
		return nil, info, err
	}
	multiMap := make(map[string]*MappingInfo)
	err = decodeRespData(bdata, &multiMap)
	return multiMap[index], info, err
}

//查看index 的mapping
func (m *Mapping) Get(ctx context.Context, index string) (*MappingInfo, *IndexErrorInfo, error) {
	return m.GetMapping(ctx, index)
}
