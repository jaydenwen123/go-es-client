package api

import (
	"context"
	"testing"

	elastic "github.com/jaydenwen123/go-es-client"
)

var client = elastic.NewClient(
	elastic.WithConnection([]string{"http://localhost:9200"},
		"", ""),
	elastic.WithPrettyQsl(true))
var ctx = context.Background()

func TestCatApi_Allocation(t *testing.T) {

	data, err := Cat(client).Allocation().Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
}

func TestCatApi_Shards(t *testing.T) {
	data, err := Cat(client).Shards().Pretty().Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
}

func TestIndicesApi(t *testing.T) {
	data, err := Cat(client).Indices().Pretty().Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
}

func TestIndicesApi_Index(t *testing.T) {
	data, err := Cat(client).Indices().
		Index(".monitoring-es-7-2020.04.29").
		//Health().
		Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
	t.Log("======================")
	data, err = Cat(client).Indices().
		Index(".monitoring-es-7-2020.04.29").
		Pretty().
		Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
}

func TestCatApi_Health(t *testing.T) {
	data, err := (&CatApi{}).Health().Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
	data, err = Cat(client).Health().
		Do(ctx)
	if err != nil {
		t.Error("TestCat error:", err.Error())
		return
	}
	t.Log("the data:\n", string(data))
}
