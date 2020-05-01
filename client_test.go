package elastic

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_Do(t *testing.T) {
	var ctx = context.Background()
	client := NewClient(
		WithConnection([]string{"http://localhost:9200"},
			"", ""),
		WithPrettyQsl(true))
	_,data, err := client.Do(ctx, "GET", "/", "")
	fmt.Println(string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/", "")
	fmt.Println(string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/_cat", "")
	fmt.Println("/_cat:->", string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/_cat/indices", "")
	fmt.Println("/_cat/indices:->", string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/_cat/indices?v", "")
	fmt.Println("/_cat/indices?v-->", string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/_cat/health", "")
	fmt.Println("/_cat/health-->", string(data), err)
	_,data, err = client.DoWithAddress(ctx, "http://localhost:9200", "GET", "/_cat/health?v", "")
	fmt.Println("/_cat/health?v-->", string(data), err)
	//t.Log(string(data), err)
}

func Test_Search(t *testing.T) {
	client := NewClient(
		WithConnection([]string{"http://localhost:9200"}, "", ""),
		WithPrettyQsl(true))
	_,data, err := client.DoWithAddress(
		context.Background(),
		"http://localhost:9200",
		"GET",
		"/.kibana_1/_search", map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": struct{}{},
			},
			"from": 0,
			"size": 1,
		})
	fmt.Println("/.kibana_1/_search-->", string(data), err)

	fmt.Println("=======================")
	_,data, err = client.Do(
		context.Background(),
		"GET",
		"/.kibana_1/_search?pretty=true",
		`{"query":{"match_all":{}},"from":0,"size":1}`)
	fmt.Println("/.kibana_1/_search-->", string(data), err)

}

func Test(t *testing.T)  {
}
