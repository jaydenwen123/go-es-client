package api

import (
	"context"
	"strings"

	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)


type getActionOp string

const (
	op_index   getActionOp = "indices"
	op_mapping getActionOp = "_mapping"
	op_setting getActionOp = "_settings"
	op_alias   getActionOp = "_alias"
)


func doGetAction(ctx context.Context, b *baseCtx, op getActionOp, indices ...string, ) ([]byte, *IndexErrorInfo, error) {
	var
	(
		//indexInfoMap   interface{}
		indexErrorInfo IndexErrorInfo
		err            error
		data           []byte
	)
	if indices == nil || len(indices) == 0 || len(indices) == 1 && indices[0] == "" {
		b.path = "/_all"
	} else {
		b.path = "/" + strings.Join(indices, ",")
	}
	//拼接op
	if op != op_index {
		b.path += "/" + string(op)
	}
	_, data, err = b.client.Get(ctx, b.path, nil)
	if gjson.Get(string(data), "error").Exists() {
		err = jsoniter.Unmarshal(data, &indexErrorInfo)
		if err != nil {
			logs.Error("jsoniter.Unmarshal error:%s", err.Error())
			return nil, nil, err
		}
		return nil, &indexErrorInfo, nil
	}
	return data, nil, nil
}

func decodeRespData(bdata []byte, multiMap interface{}) error {
	err := jsoniter.Unmarshal(bdata, multiMap)
	if err != nil {
		logs.Error("jsoniter.Unmarshal error")
	}
	return err
}

