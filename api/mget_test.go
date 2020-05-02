package api

import (
	"testing"

	"github.com/jaydenwen123/go-util"
)

func TestMget_BodyJson(t *testing.T) {
	docs, err := MgetApi(client).Index("student").BodyJson(ctx, `{
			"ids":["oH3503EBVYuGm8QUAUpw","mXts0HEBVYuGm8QUNwvI"]
		}`)
	if err != nil {
		t.Errorf("mget error:%s", err.Error())
		return
	}
	t.Logf("mget success:\n%s", util.Obj2JsonStrIndent(docs, "", "  "))
}

func TestMget_Body(t *testing.T) {
	docs, err := MgetApi(client).Index("student").Body(ctx, &MgetCondition{
		Conds: []*DocCondition{
			{
				Id: "oH3503EBVYuGm8QUAUpw",
			},
			{
				Id: "mXts0HEBVYuGm8QUNwvI",
			},
		},
	})
	if err != nil {
		t.Errorf("mget error:%s", err.Error())
		return
	}
	t.Logf("mget success:\n%s", util.Obj2JsonStrIndent(docs, "", "  "))
}


func TestMget_NoIndexBody(t *testing.T) {
	docs, err := MgetApi(client).Body(ctx, &MgetCondition{
		Conds: []*DocCondition{
			{
				Index: "student",
				Id:    "oH3503EBVYuGm8QUAUpw",
			},
			{
				Index: "student",
				Id: "mXts0HEBVYuGm8QUNwvI",
			},
		},
	})
	if err != nil {
		t.Errorf("mget error:%s", err.Error())
		return
	}
	t.Logf("mget success:\n%s", util.Obj2JsonStrIndent(docs, "", "  "))
}
