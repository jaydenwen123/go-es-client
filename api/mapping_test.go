package api

import (
	"github.com/jaydenwen123/go-util"
	"testing"
)

func TestMapping_GetMapping(t *testing.T) {
	mappingInfo, errorInfo, err := IndexAPI(client).Mapping(ctx, "student")
	if err != nil {
		t.Errorf("get mapping error:%s", err.Error())
		return
	}
	if errorInfo != nil {
		t.Errorf("get mapping info failed:\n%s", util.Obj2JsonStr(errorInfo))
		return
	}
	t.Log("get mapping info success:\n", util.Obj2JsonStr(mappingInfo))
}

func TestMapping_SetMappingWithJson(t *testing.T) {
	err := MappingAPI(client).SetWithJson(ctx, "student", `{"properties": {
       "name":{
         "type": "keyword"
       },
       "stu_no":{
         "type": "keyword"
       },
       "age":{
         "type": "integer"
       },
       "class_name":{
         "type": "text"
       },
		"class_no":{
		 "type":"keyword"
		}
     }}`)
	if err != nil {
		t.Errorf("set mapping error:%s", err.Error())
		return
	}
	t.Logf("setMapping success")
}
