package api

import (
	"testing"

	"github.com/jaydenwen123/go-util"
)

//var (
//	client = elastic.NewClient(
//		elastic.WithConnection([]string{"http://localhost:9200"},
//			"", ""),
//		elastic.WithPrettyQsl(true))
//	ctx = context.Background()
//)

func TestIndexApi_Get(t *testing.T) {
	//student
	sucInfo, errorInfo, err := IndexAPI(client).GetIndex(ctx, ".kibana_task_manager_1")
	if err != nil {
		t.Errorf("the op_index get error:\n%s", err.Error())
		return
	}
	if errorInfo != nil {
		t.Errorf("the op_index get failed:\n%s", util.Obj2JsonStrIndent(errorInfo,"","  "))
		return
	}
	t.Logf("the op_index get success:\n%s", util.Obj2JsonStrIndent(sucInfo,"","  "))
}

func TestIndexApi_Exist(t *testing.T) {
	existed, err := IndexAPI(client).Exist(ctx, "student")
	if err != nil {
		t.Errorf("check op_index existed error:%s",err.Error())
		return
	}
	if existed {
		t.Logf("the op_index:<%s> is existed.","student")
	}else{
		t.Logf("the op_index:<%s> is not existed.","student")
	}
}

func TestIndexApi_Create(t *testing.T) {
	errInfo,err := IndexAPI(client).Create(ctx, "hello")
	if err != nil {
		t.Errorf("create op_index error:\n%s\n%s",err.Error(),
			util.Obj2JsonStrIndent(errInfo,"","  "))
		return
	}
	t.Logf("create the op_index:<%s> success","student")
}

func TestIndexApi_Get2(t *testing.T) {
	//student
	//sucInfo, errorInfo, err := Index(client).GetIndex(ctx, ".kibana_task_manager_1","student")
	sucInfo, errorInfo, err := IndexAPI(client).GetAllIndices(ctx)
	//sucInfo, errorInfo, err := Index(client).GetIndex(ctx, ".kibana_task_manager_1","student")
	if err != nil {
		t.Errorf("the op_index get error:\n%s", err.Error())
		return
	}
	if errorInfo != nil {
		t.Errorf("the op_index get failed:\n%s", util.Obj2JsonStrIndent(errorInfo,"","  "))
		return
	}
	t.Logf("the get op_index count:<%d>",len(sucInfo))
	t.Logf("the op_index get success:\n%s", util.Obj2JsonStrIndent(sucInfo,"","  "))
}
