package api

import (
	"testing"
)

var docApi = DocAPI(client, "student")

func TestDocument_GetWithId(t *testing.T) {
	//oH3503EBVYuGm8QUAUpw
	//mXts0HEBVYuGm8QUNwvI
	data, err := docApi.GetWithId(ctx, "oH3503EBVYuGm8QUAUpw")
	if err != nil {
		t.Errorf("DocAPI get id error:%s", err.Error())
		return
	}
	t.Logf("docapi get document id success:%s", data)
}

func TestDocument_AddOrUpdateBodyJson(t *testing.T) {
	docId,err := docApi.AddOrUpdateBodyJson(ctx, `{
 "name":"jaydenwen123444",
 "stu_no":"444",
 "age":30,
 "class_name":"4-class",
 "class_no":"1-class-no4"
}`)
	if err != nil {
		t.Errorf("add or update doc failed:%s", err.Error())
	} else {
		t.Logf("add or update doc success.docId:%s",docId)
	}
}


