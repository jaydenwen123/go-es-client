package api

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	elastic "github.com/jaydenwen123/go-es-client"
)

type ShardsApi struct {
	*CatApi
}

//_cat/shards/{index}
func (s *ShardsApi) Index(index string) *ShardsApi {
	if s != nil {
		s.CatApi.path += "/" + index
	}
	return s
}

type IndicesApi struct {
	*CatApi
}

//_cat/indices/{index}
func (i *IndicesApi) Index(index string) *CatApi {
	if i != nil {
		i.CatApi.path += "/" + index
	}
	return i.CatApi
}

type SegmentsApi struct {
	*CatApi
}

//_cat/segments/{index}
func (s *SegmentsApi) Index(index string) *SegmentsApi {
	if s != nil {
		s.CatApi.path += "/" + index
	}
	return s
}

type CountApi struct {
	*CatApi
}

//_cat/count/{index}
func (c *CountApi) Index(index string) *CountApi {
	if c != nil {
		c.CatApi.path += "/" + index
	}
	return c
}

type RecoveryApi struct {
	*CatApi
}

//_cat/recovery/{index}
func (r *RecoveryApi) Index(index string) *RecoveryApi {
	if r != nil {
		r.CatApi.path += "/" + index
	}
	return r
}

type AliasesApi struct {
	*CatApi
}

//_cat/aliases/{alias}
func (a *AliasesApi) Index(index string) *AliasesApi {
	if a != nil {
		a.CatApi.path += "/" + index
	}
	return a
}

//ThreadPoolApi
type ThreadPoolApi struct {
	*CatApi
}

//_cat/thread_pool/{thread_pools}
func (a *ThreadPoolApi) ThreadPools(names ...string) *ThreadPoolApi {
	if a != nil {
		if len(names) == 0 {
			return a
		}
		a.CatApi.path += "/" + strings.Join(names, ",")
	}
	return a
}

//FielddataApi
type FielddataApi struct {
	*CatApi
}

//_cat/fielddata/{fields}
func (a *FielddataApi) Fields(fields ...string) *FielddataApi {
	if a != nil {
		if len(fields) == 0 {
			return a
		}
		a.CatApi.path += "/" + strings.Join(fields, ",")
	}
	return a
}

type TemplatesApi struct {
	*CatApi
}

//_cat/templates/{name}
func (a *FielddataApi) Name(name string) *FielddataApi {
	if a != nil {
		a.CatApi.path += "/" + name
	}
	return a
}

type CatApi struct {
	client *elastic.Client
	param  url.Values
	path   string
}

//Erase 擦处信息
func (c *CatApi) Erase() (path string, param url.Values) {
	path = c.path
	param = c.param
	c.path = "/_cat"
	c.param = make(map[string][]string)
	return
}

func (c *CatApi) Client() *elastic.Client {
	if c != nil {
		return c.client
	}
	return nil
}

func (c *CatApi) Param() string {
	if c == nil {
		return ""
	}
	return c.param.Encode()
}

func (c *CatApi) Path() string {
	if c == nil {
		return ""
	}
	return c.path
}

//Cat api构造器
func Cat(c *elastic.Client) *CatApi {
	return &CatApi{
		client: c,
		param:  make(map[string][]string),
		path:   "/_cat",
	}
}

//Do 执行查询
func (c *CatApi) Do(ctx context.Context) ([]byte, error) {
	if c.client == nil {
		return nil, fmt.Errorf("client is nil,you should call Cat(c *elastic.Client) to init client")
	}
	if len(strings.TrimSpace(c.path)) < 4 {
		return nil, fmt.Errorf("the cat api path is not valid:<%s>", c.path)
	}
	//构建请求
	path := c.path
	fmt.Println("=======Cat Api path:<", path, ">==================")
	_,bdata, err := c.client.Get(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	return bdata, nil
}

//Pretty 格式化显示信息
func (c *CatApi) Pretty() *CatApi {
	if c.path == "/_cat" {
		//跳过，否则会error
		return c
	}
	c.path += "?v"
	return c
}

//主要拼接query

// Allocation 分配
//_cat/allocation
//shards disk.indices disk.used disk.avail disk.total disk.percent host      ip        node
//    18      270.5mb   205.4gb       28gb    233.4gb           87 127.0.0.1 127.0.0.1 JAYDENWEN-MB0
func (c *CatApi) Allocation() *CatApi {
	c.path += "/allocation"
	return c
}

//_cat/shards
func (c *CatApi) Shards() *ShardsApi {
	if c != nil {
		c.path += "/shards"
	}
	return &ShardsApi{
		CatApi: c,
	}
}

//_cat/master
func (c *CatApi) Master() *CatApi {
	if c != nil {
		c.path += "/master"
	}
	return c
}

//_cat/nodes
func (c *CatApi) Nodes() *CatApi {
	if c != nil {
		c.path += "/nodes"
	}
	return c
}

//_cat/tasks
func (c *CatApi) Tasks() *CatApi {
	if c != nil {
		c.path += "/tasks"
	}
	return c
}

//_cat/indices
func (c *CatApi) Indices() *IndicesApi {
	if c != nil {
		c.path += "/indices"
	}
	return &IndicesApi{
		CatApi: c,
	}
}

//_cat/segments
func (c *CatApi) Segments() *SegmentsApi {
	if c != nil {
		c.path += "/segments"
	}
	return &SegmentsApi{
		CatApi: c,
	}
}

//_cat/count
func (c *CatApi) Count() *CountApi {
	if c != nil {
		c.path += "/count"
	}
	return &CountApi{
		CatApi: c,
	}
}

//_cat/recovery
func (c *CatApi) Recovery() *RecoveryApi {
	if c != nil {
		c.path += "/recovery"
	}
	return &RecoveryApi{
		CatApi: c,
	}
}

//_cat/health
func (c *CatApi) Health() *CatApi {
	if c != nil {
		c.path += "/health"
	}
	return c
}

//_cat/pending_tasks
func (c *CatApi) PendingTasks() *CatApi {
	if c != nil {
		c.path += "/pending_tasks"
	}
	return c
}

//_cat/aliases
func (c *CatApi) Aliases() *AliasesApi {
	if c != nil {
		c.path += "/aliases"
	}
	return &AliasesApi{
		CatApi: c,
	}
}

//_cat/thread_pool
func (c *CatApi) ThreadPool() *ThreadPoolApi {
	if c != nil {
		c.path += "/thread_pool"
	}
	return &ThreadPoolApi{
		CatApi: c,
	}
}

//_cat/plugins
func (c *CatApi) Plugins() *CatApi {
	if c != nil {
		c.path += "/plugins"
	}
	return c
}

//_cat/fielddata
func (c *CatApi) Fielddata() *FielddataApi {
	if c != nil {
		c.path += "/fielddata"
	}
	return &FielddataApi{
		CatApi: c,
	}
}

//_cat/nodeattrs
func (c *CatApi) Nodeattrs() *CatApi {
	if c != nil {
		c.path += "/nodeattrs"
	}
	return c
}

//_cat/repositories
func (c *CatApi) Repositories() *CatApi {
	if c != nil {
		c.path += "/repositories"
	}
	return c
}

//_cat/snapshots/{repository}
func (c *CatApi) Snapshots(repository string) *CatApi {
	if c != nil {
		c.path += "/snapshots/" + repository
	}
	return c
}

//_cat/templates
func (c *CatApi) Templates() *TemplatesApi {
	if c != nil {
		c.path += "/templates"
	}
	return &TemplatesApi{
		CatApi: c,
	}
}
