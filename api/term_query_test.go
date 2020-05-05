package api

import (
	"github.com/jaydenwen123/go-util"
	"testing"
)

func TestTerm_Term(t *testing.T) {
	term := QueryApi(client).
		Indices("all_articles").
		Match("url", "https://www.cnblogs.com/yanghj010/p/9647491.html")
	t.Logf("the term query:\n%s", term.String())
	queryInfo, err := term.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))
}

func TestTerm_Terms(t *testing.T) {
	terms := QueryApi(client).
		Indices("all_articles").
		Terms("url", "https://www.cnblogs.com/CraryPrimitiveMan/p/4657835.html",
			"https://www.cnblogs.com/CraryPrimitiveMan/p/4657055.html")
	t.Logf("the terms query:\n%s", terms.String())
	queryInfo, err := terms.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo.Hits.Hits[0].Source, "", "  "))
}


