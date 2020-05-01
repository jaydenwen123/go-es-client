package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
)

// 1.采用线程池
// 2.自动化拼接各个参数和url
// 3.心跳检测
// 4.便捷使用

var header = map[string][]string{
	"Content-Type": {"application/json"},
}

type ClientOption func(c *Client)

var defaultClient Client

//Connection 连接信息
type Connection struct {
	address  string
	username string
	password string
}

//Client elasticsearch的客户端
type Client struct {
	//连接信息,多个连接
	conn []*Connection
	//http的客户端
	c *http.Client

	prettyQsl bool
}

//WithPrettyQsl 是否美化qsl语句
func WithPrettyQsl(pretty bool) ClientOption {
	return func(c *Client) {
		c.prettyQsl = pretty
	}
}

//WithConnection 设置连接
func WithConnection(urls []string, username, password string) ClientOption {
	return func(c *Client) {
		if len(urls) > 0 {
			c.conn = make([]*Connection, 0)
			for _, url := range urls {
				c.conn = append(c.conn, &Connection{
					address:  url,
					username: username,
					password: password,
				})
			}
		}
	}
}

//NewClient  创建客户端
func NewClient(options ...ClientOption) *Client {
	client := defaultClient
	//赋值
	for _, op := range options {
		op(&client)
	}
	//todo 后续进行改造
	client.c = http.DefaultClient
	return &client
}

//BuildSearchQuery 构建查询语句
func (c *Client) BuildSearchQuery() {
}

func (c *Client) Get(ctx context.Context, path string, body interface{}) (*http.Response, []byte, error) {
	return c.Do(ctx, "GET", path, body)

}

func (c *Client) Post(ctx context.Context, path string, body interface{}) (*http.Response, []byte, error) {
	return c.Do(ctx, "POST", path, body)

}

func (c *Client) Delete(ctx context.Context, path string, body interface{}) (*http.Response, []byte, error) {
	return c.Do(ctx, "DELETE", path, body)
}

func (c *Client) Put(ctx context.Context, path string, body interface{}) (*http.Response, []byte, error) {
	return c.Do(ctx, "PUT", path, body)
}

//DoWithAddress 指定地址请求
func (c *Client) DoWithAddress(ctx context.Context, address, method, path string, body interface{}) (*http.Response, []byte, error) {
	//todo 通过ctx来控制超时
	address, err := c.validParam(method, address, path)
	if err != nil {
		return nil, nil, err
	}
	realUrl := address + path
	bodyData, err := c.parseBody(body)
	if err != nil {
		return nil, nil, err
	}
	//todo 后面改造
	logs.Debug("the execute url:[%s]", realUrl)
	req, err := http.NewRequest(method, realUrl, bodyData)
	if err != nil {
		logs.Error("http.NewRequest error:%s", err.Error())
		return nil, nil, err
	}
	//添加请求头
	//req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)
	req.Header = header
	resp, err := c.c.Do(req)
	if err != nil {
		if IsContextErr(err) {
			// Proceed, but don't mark the node as dead
			logs.Error("http Do error is IsContextErr:%s", err.Error())
			return resp, nil, err
		}
		logs.Error("http client do   error:%s", err.Error())
		return resp, nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("ioutil.ReadAll the resp data error:%s", err.Error())
		return resp, nil, err
	}
	return resp, data, nil
}

//Do 执行查询
//todo 做连接池、负载均衡
func (c *Client) Do(ctx context.Context, method, url string, body interface{}) (*http.Response, []byte, error) {
	return c.DoWithAddress(ctx, "", method, url, body)

}

// IsContextErr returns true if the error is from a context that was canceled or deadline exceeded
func IsContextErr(err error) bool {
	if err == context.Canceled || err == context.DeadlineExceeded {
		return true
	}
	// This happens e.g. on redirect errors, see https://golang.org/src/net/http/client_test.go#L329
	if ue, ok := err.(*url.Error); ok {
		if ue.Temporary() {
			return true
		}
		// Use of an AWS Signing Transport can result in a wrapped url.Error
		return IsContextErr(ue.Err)
	}
	return false
}

//parseBody parse body param to http body formation
func (c *Client) parseBody(body interface{}) (io.Reader, error) {
	var bodyData io.Reader
	var err error
	var bdata []byte
	if body == nil {
		return nil, nil
	}
	switch body.(type) {
	case string:
		bodyStr := body.(string)
		if len(strings.TrimSpace(bodyStr)) == 0 || bodyStr == "" {
			bodyData = nil
		} else {
			bodyData = strings.NewReader(bodyStr)
		}
		//json字符串
		if gjson.Valid(bodyStr) {
			if c.prettyQsl {
				bodyStr = string(pretty.Pretty([]byte(bodyStr)))
			}
		}
		logs.Debug("the body data:\n%s", bodyStr)
	default:
		if c.prettyQsl {
			bdata, err = json.MarshalIndent(body, "", "\t")
		} else {
			bdata, err = json.Marshal(body)
		}
		if err != nil {
			logs.Error("the body is decoded error:%s", err.Error())
			return nil, err
		}
		bodyData = bytes.NewReader(bdata)
		logs.Debug("the body data:\n%s", string(bdata))
	}
	return bodyData, nil
}

//validParam 验证参数
func (c *Client) validParam(method string, address string, url string) (string, error) {
	if len(strings.TrimSpace(method)) == 0 {
		return address, fmt.Errorf("the http method:<%s> is not valid", method)
	}

	if len(strings.TrimSpace(address)) == 0 {
		logs.Debug("the http address param:<%s> is empty,now is getting address from connection.", address)
		if c.conn == nil || len(c.conn) == 0 || len(c.conn[0].address) == 0 {
			return "", fmt.Errorf("the  connection  address is not valid")
		} else {
			address = c.conn[0].address
		}
	}

	if len(strings.TrimSpace(url)) == 0 {
		return address, fmt.Errorf("the http url:<%s> is not valid", url)
	}
	return address, nil
}
