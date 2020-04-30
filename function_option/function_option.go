package function_option

import "fmt"

//采用函数选项设计模式
type Pormfermer interface {
	//执行
	Do() error
	//开始执行前
	BeforeExecute() error
	//执行后
	AfterExecute() error
}

type Client struct {
	conn    Connection //连接
	timeout int        //超时
	retry   int        //重试
	IsDead  bool       //是否已死
}

func (c *Client) String() string {
	return fmt.Sprintf("timeout:<%d>\tretry:<%d>\tconn:<%v>\tisDead:<%v>", c.timeout, c.retry, c.conn, c.IsDead)
}

type ClientOption func(c *Client)

var defaultClient = Client{
	timeout: 5,
	retry:   3,
}

//NewClient 创建客户端
func NewClient(options ...ClientOption) *Client {
	c := defaultClient
	for _, op := range options {
		op(&c)
	}
	return &c
}

func WithIsDead(isDead bool) ClientOption {
	return func(c *Client) {
		c.IsDead = isDead
	}
}

//WithTimeout 设置超时时间
func WithTimeout(timeout int) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

//WithRetry  设置retry属性
func WithRetry(retry int) ClientOption {
	return func(c *Client) {
		c.retry = retry
	}
}

//WithConnection 设置conn
func WithConnection(conn Connection) ClientOption {
	return func(c *Client) {
		c.conn = conn
	}
}

func (c *Client) Do() error {
	c.BeforeExecute()
	fmt.Println("Do...")
	c.AfterExecute()
	return nil
}

func (c *Client) BeforeExecute() error {
	fmt.Println("BeforeExecute...")
	return nil
}

func (c *Client) AfterExecute() error {
	fmt.Println("AfterExecute...")
	return nil
}

type Connection struct {
}
