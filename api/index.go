package api

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"strings"

	elastic "github.com/jaydenwen123/go-es-client"
)

type IndexInfo struct {
	Aliases  Aliases  `json:"aliases"`
	Mappings Mappings `json:"mappings"`
	Settings Settings `json:"settings"`
}

//todo 补充
type Aliases struct {
}

type MigrationMappingPropertyHashes struct {
	MigrationVersion string `json:"migrationVersion"`
	Task             string `json:"task"`
	UpdatedAt        string `json:"updated_at"`
	References       string `json:"references"`
	Namespace        string `json:"namespace"`
	Type             string `json:"type"`
	Config           string `json:"config"`
}
type Meta struct {
	MigrationMappingPropertyHashes MigrationMappingPropertyHashes `json:"migrationMappingPropertyHashes"`
}
type Config struct {
	Dynamic    string `json:"dynamic"`
	Properties struct {
		BuildNum struct {
			Type string `json:"type"`
		} `json:"buildNum"`
	} `json:"properties"`
}
type Keyword struct {
	Type        string `json:"type"`
	IgnoreAbove int    `json:"ignore_above"`
}
type MigrationVersion struct {
	Dynamic    string `json:"dynamic"`
	Properties struct {
		Task struct {
			Type   string `json:"type"`
			Fields struct {
				Keyword Keyword `json:"keyword"`
			} `json:"fields"`
		} `json:"task"`
	} `json:"properties"`
}

type References struct {
	Type       string `json:"type"`
	Properties struct {
		ID struct {
			Type string `json:"type"`
		} `json:"id"`
		Name struct {
			Type string `json:"type"`
		} `json:"name"`
		Type struct {
			Type string `json:"type"`
		} `json:"type"`
	} `json:"properties"`
}
type Schedule struct {
	Properties struct {
		Interval struct {
			Type string `json:"type"`
		} `json:"interval"`
	} `json:"properties"`
}
type TaskProperties struct {
	Attempts struct {
		Type string `json:"type"`
	} `json:"attempts"`
	OwnerID struct {
		Type string `json:"type"`
	} `json:"ownerId"`
	Params struct {
		Type string `json:"type"`
	} `json:"params"`
	RetryAt struct {
		Type string `json:"type"`
	} `json:"retryAt"`
	RunAt struct {
		Type string `json:"type"`
	} `json:"runAt"`
	Schedule    Schedule `json:"schedule"`
	ScheduledAt struct {
		Type string `json:"type"`
	} `json:"scheduledAt"`
	Scope struct {
		Type string `json:"type"`
	} `json:"scope"`
	StartedAt struct {
		Type string `json:"type"`
	} `json:"startedAt"`
	State struct {
		Type string `json:"type"`
	} `json:"state"`
	Status struct {
		Type string `json:"type"`
	} `json:"status"`
	TaskType struct {
		Type string `json:"type"`
	} `json:"taskType"`
	User struct {
		Type string `json:"type"`
	} `json:"user"`
}
type Properties struct {
	Config           Config           `json:"config"`
	MigrationVersion MigrationVersion `json:"migrationVersion"`
	Namespace        struct {
		Type string `json:"type"`
	} `json:"namespace"`
	References References `json:"references"`
	Task       struct {
		Properties TaskProperties `json:"properties"`
	} `json:"task"`
	Type struct {
		Type string `json:"type"`
	} `json:"type"`
	UpdatedAt struct {
		Type string `json:"type"`
	} `json:"updated_at"`
}
type Mappings struct {
	Dynamic    string     `json:"dynamic"`
	Meta       Meta       `json:"_meta"`
	Properties Properties `json:"properties"`
}

//todo 补充
type Version struct {
	Created string `json:"created"`
}
type IndexObj struct {
	NumberOfShards     string  `json:"number_of_shards"`
	AutoExpandReplicas string  `json:"auto_expand_replicas"`
	ProvidedName       string  `json:"provided_name"`
	CreationDate       string  `json:"creation_date"`
	NumberOfReplicas   string  `json:"number_of_replicas"`
	UUID               string  `json:"uuid"`
	Version            Version `json:"version"`
}
type Settings struct {
	Index IndexObj `json:"index"`
}

//Index 索引api
type Index struct {
	index  string
	client *elastic.Client
	param  url.Values
	path   string
}

type IndexErrorInfo struct {
	Error struct {
		RootCause []struct {
			Type         string `json:"type"`
			Reason       string `json:"reason"`
			ResourceType string `json:"resource.type"`
			ResourceID   string `json:"resource.id"`
			IndexUUID    string `json:"index_uuid"`
			Index        string `json:"index"`
		} `json:"root_cause"`
		Type         string `json:"type"`
		Reason       string `json:"reason"`
		ResourceType string `json:"resource.type"`
		ResourceID   string `json:"resource.id"`
		IndexUUID    string `json:"index_uuid"`
		Index        string `json:"index"`
	} `json:"error"`
	Status int `json:"status"`
}

func (i *Index) Param() string {
	if i == nil {
		return ""
	}
	return i.param.Encode()
}

func (i *Index) Path() string {
	if i == nil {
		return ""
	}
	return i.path
}

func (i *Index) Client() *elastic.Client {
	if i != nil {
		return i.client
	}
	return nil
}

//Index index api
func IndexAPI(client *elastic.Client) *Index {
	return &Index{
		client: client,
		param:  make(map[string][]string),
	}
}

//Index index api
func IndexApi(client *elastic.Client) *Index {
	return &Index{
		client: client,
		param:  make(map[string][]string),
	}
}

//新增索引 put
func (i *Index) Create(ctx context.Context, index string) (*IndexErrorInfo, error) {
	var (
		err   error
		bdata []byte
		rsp   *http.Response
	)
	err = i.validIndex(index)
	if err != nil {
		return nil, err
	}
	rsp, bdata, err = i.client.Put(ctx, i.path, nil)
	return i.wrapResp(err, rsp, bdata)
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
		return fmt.Errorf("the index is empty")
	}
	i.path = "/" + index
	return nil
}

//todo
//reindex重建索引

//All is the GetAll alias function.
func (i *Index) All(ctx context.Context) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	return i.Get(ctx)
}

//GetAll get all indices
func (i *Index) GetAll(ctx context.Context) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	return i.Get(ctx, )
}

//Get 查看索引
// if indices is nil or len(indices)==0  path use /_all
func (i *Index) Get(ctx context.Context, indices ...string) (map[string]*IndexInfo, *IndexErrorInfo, error) {
	var
	(
		indexInfoMap   map[string]*IndexInfo
		indexErrorInfo IndexErrorInfo
		err            error
		data           []byte
	)
	if indices == nil || len(indices) == 0 || len(indices) == 1 && indices[0] == "" {
		i.path = "/_all"
	} else {
		i.path = "/" + strings.Join(indices, ",")
	}
	_, data, err = i.client.Get(ctx, i.path, nil)

	if gjson.Get(string(data), "error").Exists() {
		err = jsoniter.Unmarshal(data, &indexErrorInfo)
		if err != nil {
			logs.Error("jsoniter.Unmarshal error:%s", err.Error())
			return nil, nil, err
		}
		return nil, &indexErrorInfo, nil
	} else {
		indexInfoMap = make(map[string]*IndexInfo)
		err = jsoniter.Unmarshal(data, &indexInfoMap)
		if err != nil {
			logs.Error("jsoniter.Unmarshal error:%s", err.Error())
			return nil, nil, err
		}
		return indexInfoMap, nil, nil
	}
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
