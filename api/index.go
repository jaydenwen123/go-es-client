package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	elastic "github.com/jaydenwen123/go-es-client"

	"github.com/astaxie/beego/logs"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

//todo 重新定义数据
type IndexInfo struct {
	Aliases  *AliasInfo   `json:"aliases,omitempty"`
	Mappings *MappingInfo `json:"mappings,omitempty"`
	Settings *SettingInfo `json:"settings,omitempty"`
}

//todo 补充
type AliasInfo struct {
}

type MigrationMappingPropertyHashes struct {
	MigrationVersion string `json:"migrationVersion,omitempty"`
	Task             string `json:"task,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	References       string `json:"references,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
	Type             string `json:"type,omitempty"`
	Config           string `json:"config,omitempty"`
}
type Meta struct {
	MigrationMappingPropertyHashes *MigrationMappingPropertyHashes `json:"migrationMappingPropertyHashes,omitempty"`
}
type Config struct {
	Dynamic    string `json:"dynamic,omitempty"`
	Properties struct {
		BuildNum struct {
			Type string `json:"type,omitempty"`
		} `json:"buildNum,omitempty"`
	} `json:"properties,omitempty"`
}
type Keyword struct {
	Type        string `json:"type,omitempty"`
	IgnoreAbove int    `json:"ignore_above,omitempty"`
}
type MigrationVersion struct {
	Dynamic    string `json:"dynamic,omitempty"`
	Properties struct {
		Task struct {
			Type   string `json:"type,omitempty"`
			Fields struct {
				Keyword Keyword `json:"keyword,omitempty"`
			} `json:"fields,omitempty"`
		} `json:"task,omitempty"`
	} `json:"properties,omitempty"`
}

type References struct {
	Type       string `json:"type,omitempty"`
	Properties struct {
		ID struct {
			Type string `json:"type,omitempty"`
		} `json:"id,omitempty"`
		Name struct {
			Type string `json:"type,omitempty"`
		} `json:"name,omitempty"`
		Type struct {
			Type string `json:"type,omitempty"`
		} `json:"type,omitempty"`
	} `json:"properties,omitempty"`
}
type Schedule struct {
	Properties struct {
		Interval struct {
			Type string `json:"type,omitempty"`
		} `json:"interval,omitempty"`
	} `json:"properties,omitempty"`
}
type TaskProperties struct {
	Attempts struct {
		Type string `json:"type,omitempty"`
	} `json:"attempts,omitempty"`
	OwnerID struct {
		Type string `json:"type,omitempty"`
	} `json:"ownerId,omitempty"`
	Params struct {
		Type string `json:"type,omitempty"`
	} `json:"params,omitempty"`
	RetryAt struct {
		Type string `json:"type,omitempty"`
	} `json:"retryAt,omitempty"`
	RunAt struct {
		Type string `json:"type,omitempty"`
	} `json:"runAt,omitempty"`
	Schedule    Schedule `json:"schedule,omitempty"`
	ScheduledAt struct {
		Type string `json:"type,omitempty"`
	} `json:"scheduledAt,omitempty"`
	Scope struct {
		Type string `json:"type,omitempty"`
	} `json:"scope,omitempty"`
	StartedAt struct {
		Type string `json:"type,omitempty"`
	} `json:"startedAt,omitempty"`
	State struct {
		Type string `json:"type,omitempty"`
	} `json:"state,omitempty"`
	Status struct {
		Type string `json:"type,omitempty"`
	} `json:"status,omitempty"`
	TaskType struct {
		Type string `json:"type,omitempty"`
	} `json:"taskType,omitempty"`
	User struct {
		Type string `json:"type,omitempty"`
	} `json:"user,omitempty"`
}
type Properties struct {
	Config           Config           `json:"config,omitempty"`
	MigrationVersion MigrationVersion `json:"migrationVersion,omitempty"`
	Namespace        struct {
		Type string `json:"type,omitempty"`
	} `json:"namespace,omitempty"`
	References References `json:"references,omitempty"`
	Task       struct {
		Properties TaskProperties `json:"properties,omitempty"`
	} `json:"task,omitempty"`
	Type struct {
		Type string `json:"type,omitempty"`
	} `json:"type,omitempty"`
	UpdatedAt struct {
		Type string `json:"type,omitempty"`
	} `json:"updated_at,omitempty"`
}
type MappingInfo struct {
	Dynamic    string     `json:"dynamic,omitempty"`
	Meta       Meta       `json:"_meta,omitempty"`
	Properties Properties `json:"properties,omitempty"`
}

//todo 补充
type Version struct {
	Created string `json:"created,omitempty"`
}
type IndexObj struct {
	NumberOfShards     string  `json:"number_of_shards,omitempty"`
	AutoExpandReplicas string  `json:"auto_expand_replicas,omitempty"`
	ProvidedName       string  `json:"provided_name,omitempty"`
	CreationDate       string  `json:"creation_date,omitempty"`
	NumberOfReplicas   string  `json:"number_of_replicas,omitempty"`
	UUID               string  `json:"uuid,omitempty"`
	Version            Version `json:"version,omitempty"`
}
type SettingInfo struct {
	Index IndexObj `json:"op_index,omitempty"`
}

//Index 索引api
type Index struct {
	*baseCtx
	index string
}

type IndexErrorInfo struct {
	Error struct {
		RootCause []struct {
			Type         string `json:"type,omitempty"`
			Reason       string `json:"reason,omitempty"`
			ResourceType string `json:"resource.type,omitempty"`
			ResourceID   string `json:"resource.id,omitempty"`
			IndexUUID    string `json:"index_uuid,omitempty"`
			Index        string `json:"op_index,omitempty"`
		} `json:"root_cause,omitempty"`
		Type         string `json:"type,omitempty"`
		Reason       string `json:"reason,omitempty"`
		ResourceType string `json:"resource.type,omitempty"`
		ResourceID   string `json:"resource.id,omitempty"`
		IndexUUID    string `json:"index_uuid,omitempty"`
		Index        string `json:"op_index,omitempty"`
	} `json:"error,omitempty"`
	Status int `json:"status,omitempty"`
}

//Index op_index api
func IndexAPI(client *elastic.Client) *Index {
	return &Index{
		baseCtx: &baseCtx{
			client: client,
			param:  make(map[string][]string),
		},
	}
}

//Index op_index api
func IndexApi(client *elastic.Client) *Index {
	return IndexAPI(client)
}

//新增索引 put
func (i *Index) Create(ctx context.Context, index string) (*IndexErrorInfo, error) {
	return i.createIndex(ctx, index, nil)
}

func (i *Index) createIndex(ctx context.Context, index string, mapping interface{}) (*IndexErrorInfo, error) {
	var (
		err   error
		bdata []byte
		rsp   *http.Response
	)
	err = i.validIndex(index)
	if err != nil {
		return nil, err
	}
	rsp, bdata, err = i.client.Put(ctx, i.path, mapping)
	return i.wrapResp(err, rsp, bdata)
}

//新增索引 put
func (i *Index) CreateWithMapping(ctx context.Context, index string, mapping string) (*IndexErrorInfo, error) {
	return i.createIndex(ctx, index, mapping)
}

//wrapResp 包装回包
func (i *Index) wrapResp(err error, rsp *http.Response, bdata []byte) (*IndexErrorInfo, error) {
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		var indexError IndexErrorInfo
		err = jsoniter.Unmarshal(bdata, &indexError)
		return &indexError, fmt.Errorf("the error detail info into IndexErrorInf")
	}
	return nil, nil
}

//删除索引
//delete
func (i *Index) Delete(ctx context.Context, index string) (*IndexErrorInfo, error) {
	var (
		err   error
		rsp   *http.Response
		bdata []byte
	)
	err = i.validIndex(index)
	if err != nil {
		return nil, err
	}
	rsp, bdata, err = i.client.Delete(ctx, i.path, nil)
	return i.wrapResp(err, rsp, bdata)
}

//validIndex 验证索引
func (i *Index) validIndex(index string) error {
	if len(strings.TrimSpace(index)) == 0 {
		return fmt.Errorf("the op_index is empty")
	}
	i.path = "/" + index
	return nil
}

//todo
//reindex重建索引

//Alias 获取单个索引的alias信息
func (i *Index) Alias(ctx context.Context, index string) (*AliasInfo, *IndexErrorInfo, error) {
	multiAlias, info, err := i.MultiAlias(ctx, index)
	if err != nil {
		return nil, info, err
	}
	return multiAlias[index], info, err
}

//MultiAlias 获取多个索引的alias
func (i *Index) MultiAlias(ctx context.Context, indices ...string) (map[string]*AliasInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, i.baseCtx, op_alias, indices...)
	multiMap := make(map[string]*AliasInfo)
	err = decodeRespData(bdata, &multiMap)
	return multiMap, info, err
}

//SettingInfo 获取单个索引的setting信息
func (i *Index) Settings(ctx context.Context, index string) (*SettingInfo, *IndexErrorInfo, error) {
	multiSettings, info, err := i.MultiSettings(ctx, index)
	if err != nil {
		return nil, info, err
	}
	return multiSettings[index], info, err
}

//MultiSettings 获取多个索引的settings信息
func (i *Index) MultiSettings(ctx context.Context, indices ...string) (map[string]*SettingInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, i.baseCtx, op_setting, indices...)
	multiMap := make(map[string]*SettingInfo)
	err = decodeRespData(bdata, &multiMap)
	return multiMap, info, err
}

//Mapping 获取单个索引的mapping信息
func (i *Index) Mapping(ctx context.Context, index string) (*MappingInfo, *IndexErrorInfo, error) {
	multiMappings, info, err := i.MultiMapping(ctx, index)
	if err != nil {
		return nil, info, err
	}
	return multiMappings[index], info, err
}

//MultiMapping 获取多个索引的mappings
func (i *Index) MultiMapping(ctx context.Context, indices ...string) (map[string]*MappingInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, i.baseCtx, op_mapping, indices...)
	multiMap := make(map[string]*IndexInfo)
	err = decodeRespData(bdata, &multiMap)
	if err != nil {
		return nil, info, err
	}
	mappings := make(map[string]*MappingInfo)
	for key, indexInfo := range multiMap {
		mappings[key] = indexInfo.Mappings
	}
	return mappings, info, err
}

//AllIndices is the GetAllIndices alias function.
func (i *Index) AllIndices(ctx context.Context) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	return i.GetMultiIndex(ctx)
}

//GetAllIndices get all indices
func (i *Index) GetAllIndices(ctx context.Context) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	return i.GetMultiIndex(ctx)
}

//GetIndex 获取单个索引的详细信息
func (i *Index) GetIndex(ctx context.Context, index string) (*IndexInfo, *IndexErrorInfo, error) {
	multiInfo, info, err := i.GetMultiIndex(ctx, index)
	return multiInfo[index], info, err
}

//GetMultiIndex 查看多个索引信息
// if indices is nil or len(indices)==0  path use /_all
func (i *Index) GetMultiIndex(ctx context.Context, indices ...string) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	bdata, info, err := doGetAction(ctx, i.baseCtx, op_index, indices...)
	if err != nil {
		return nil, info, err
	}
	multiMap := make(map[string]*IndexInfo)
	err = decodeRespData(bdata, &multiMap)
	return multiMap, info, err
}

//Erase 擦处信息
func (i *Index) Erase() (path string, param url.Values) {
	path = i.path
	param = i.param
	i.path = ""
	i.param = make(map[string][]string)
	return
}

//Exist 判断索引是否存在
func (i *Index) Exist(ctx context.Context, index string) (bool, error) {
	var resp *http.Response
	var err error
	err = i.validIndex(index)
	if err != nil {
		return false, err
	}
	resp, _, err = i.client.Do(ctx, "HEAD", i.path, nil)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}

//Close 关闭索引
func (i *Index) Close(ctx context.Context, index string) (bool, error) {
	return i.open_close(ctx, index, "/_close")
}

//Open 打开索引
func (i *Index) Open(ctx context.Context, index string) (bool, error) {
	return i.open_close(ctx, index, "/_open")
}

func (i *Index) open_close(ctx context.Context, index string, op string) (bool, error) {
	var err error
	err = i.validIndex(index)
	if err != nil {
		return false, err
	}
	i.path += op
	resp, bdata, err := i.client.Post(ctx, i.path, nil)
	if err != nil {
		logs.Error("do <%s>  action error:%s", i.path, err.Error())
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("the error info:%s", bdata)
	}
	if gjson.Get(string(bdata), "error").Exists() {
		return false, fmt.Errorf("the error info:%s", bdata)
	}
	return true, nil
}
