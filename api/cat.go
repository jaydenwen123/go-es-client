package api

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	elastic "github.com/jaydenwen123/go-es-client"
)

type ShardsApi struct {
	*Cat
}

//_cat/shards/{op_index}
func (s *ShardsApi) Index(index string) *ShardsApi {
	if s != nil {
		s.Cat.path += "/" + index
	}
	return s
}

type IndicesApi struct {
	*Cat
}

//_cat/indices/{op_index}
func (i *IndicesApi) Index(index string) *Cat {
	if i != nil {
		i.Cat.path += "/" + index
	}
	return i.Cat
}

type SegmentsApi struct {
	*Cat
}

//_cat/segments/{op_index}
func (s *SegmentsApi) Index(index string) *SegmentsApi {
	if s != nil {
		s.Cat.path += "/" + index
	}
	return s
}

type CountApi struct {
	*Cat
}

//_cat/count/{op_index}
func (c *CountApi) Index(index string) *CountApi {
	if c != nil {
		c.Cat.path += "/" + index
	}
	return c
}

type RecoveryApi struct {
	*Cat
}

//_cat/recovery/{op_index}
func (r *RecoveryApi) Index(index string) *RecoveryApi {
	if r != nil {
		r.Cat.path += "/" + index
	}
	return r
}

type AliasesApi struct {
	*Cat
}

//_cat/aliases/{alias}
func (a *AliasesApi) Index(index string) *AliasesApi {
	if a != nil {
		a.Cat.path += "/" + index
	}
	return a
}

//ThreadPoolApi
type ThreadPoolApi struct {
	*Cat
}

//_cat/thread_pool/{thread_pools}
func (a *ThreadPoolApi) ThreadPools(names ...string) *ThreadPoolApi {
	if a != nil {
		if len(names) == 0 {
			return a
		}
		a.Cat.path += "/" + strings.Join(names, ",")
	}
	return a
}

//FielddataApi
type FielddataApi struct {
	*Cat
}

//_cat/fielddata/{fields}
func (a *FielddataApi) Fields(fields ...string) *FielddataApi {
	if a != nil {
		if len(fields) == 0 {
			return a
		}
		a.Cat.path += "/" + strings.Join(fields, ",")
	}
	return a
}

type TemplatesApi struct {
	*Cat
}

//_cat/templates/{name}
func (a *FielddataApi) Name(name string) *FielddataApi {
	if a != nil {
		a.Cat.path += "/" + name
	}
	return a
}

type Cat struct {
	*baseCtx
}

//Erase 擦处信息
func (c *Cat) Erase() (path string, param url.Values) {
	path = c.path
	param = c.param
	c.path = "/_cat"
	c.param = make(map[string][]string)
	return
}

//CatAPI 构造器
func CatAPI(c *elastic.Client) *Cat {
	return &Cat{
		baseCtx: &baseCtx{
			path:   "/_cat",
			client: c,
			param:  make(map[string][]string),
		},
	}
}

//CatApi 构造器
func CatApi(c *elastic.Client) *Cat {
	return CatAPI(c)
}

//Do 执行查询
func (c *Cat) Do(ctx context.Context) ([]byte, error) {
	if c.client == nil {
		return nil, fmt.Errorf("client is nil,you should call Cat(c *elastic.Client) to init client")
	}
	if len(strings.TrimSpace(c.path)) < 4 {
		return nil, fmt.Errorf("the cat api path is not valid:<%s>", c.path)
	}
	//构建请求
	path := c.path
	fmt.Println("=======Cat Api path:<", path, ">==================")
	_, bdata, err := c.client.Get(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	return bdata, nil
}

//Pretty 格式化显示信息
func (c *Cat) Pretty() *Cat {
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
func (c *Cat) Allocation() *Cat {
	c.path += "/allocation"
	return c
}

//_cat/shards
func (c *Cat) Shards() *ShardsApi {
	if c != nil {
		c.path += "/shards"
	}
	return &ShardsApi{
		Cat: c,
	}
}

//_cat/master
func (c *Cat) Master() *Cat {
	if c != nil {
		c.path += "/master"
	}
	return c
}

//_cat/nodes
func (c *Cat) Nodes() *Cat {
	if c != nil {
		c.path += "/nodes"
	}
	return c
}

//_cat/tasks
func (c *Cat) Tasks() *Cat {
	if c != nil {
		c.path += "/tasks"
	}
	return c
}

//_cat/indices
func (c *Cat) Indices() *IndicesApi {
	if c != nil {
		c.path += "/indices"
	}
	return &IndicesApi{
		Cat: c,
	}
}

//_cat/segments
func (c *Cat) Segments() *SegmentsApi {
	if c != nil {
		c.path += "/segments"
	}
	return &SegmentsApi{
		Cat: c,
	}
}

//_cat/count
func (c *Cat) Count() *CountApi {
	if c != nil {
		c.path += "/count"
	}
	return &CountApi{
		Cat: c,
	}
}

//_cat/recovery
func (c *Cat) Recovery() *RecoveryApi {
	if c != nil {
		c.path += "/recovery"
	}
	return &RecoveryApi{
		Cat: c,
	}
}

//_cat/health
func (c *Cat) Health() *Cat {
	if c != nil {
		c.path += "/health"
	}
	return c
}

//_cat/pending_tasks
func (c *Cat) PendingTasks() *Cat {
	if c != nil {
		c.path += "/pending_tasks"
	}
	return c
}

//_cat/aliases
func (c *Cat) Aliases() *AliasesApi {
	if c != nil {
		c.path += "/aliases"
	}
	return &AliasesApi{
		Cat: c,
	}
}

//_cat/thread_pool
func (c *Cat) ThreadPool() *ThreadPoolApi {
	if c != nil {
		c.path += "/thread_pool"
	}
	return &ThreadPoolApi{
		Cat: c,
	}
}

//_cat/plugins
func (c *Cat) Plugins() *Cat {
	if c != nil {
		c.path += "/plugins"
	}
	return c
}

//_cat/fielddata
func (c *Cat) Fielddata() *FielddataApi {
	if c != nil {
		c.path += "/fielddata"
	}
	return &FielddataApi{
		Cat: c,
	}
}

//_cat/nodeattrs
func (c *Cat) Nodeattrs() *Cat {
	if c != nil {
		c.path += "/nodeattrs"
	}
	return c
}

//_cat/repositories
func (c *Cat) Repositories() *Cat {
	if c != nil {
		c.path += "/repositories"
	}
	return c
}

//_cat/snapshots/{repository}
func (c *Cat) Snapshots(repository string) *Cat {
	if c != nil {
		c.path += "/snapshots/" + repository
	}
	return c
}

//_cat/templates
func (c *Cat) Templates() *TemplatesApi {
	if c != nil {
		c.path += "/templates"
	}
	return &TemplatesApi{
		Cat: c,
	}
}
