package api

import (
	"github.com/jaydenwen123/go-util"
	"testing"
)

func TestQueryProxy_Query(t *testing.T) {
	match := QueryApi(client).
		Indices("all_articles").
		Regexp("url", &RegexpCond{
		Value:   ".*implement-unary-grpc.+",
		Flags:   RegexpFlagOp_ALL,
		Rewrite: RewriteOp_Constant_Score,
	})
	t.Logf("the terms query:\n%s", match.String())
	queryInfo, err := match.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))

}
