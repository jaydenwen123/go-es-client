package api

import (
	"github.com/jaydenwen123/go-util"
	"testing"
)

func TestQuery_Match(t *testing.T) {
	match := QueryApi(client).
		Indices("all_articles").
		Match("title", "学习技巧")
	t.Logf("the terms query:\n%s", match.String())
	queryInfo, err := match.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))
}

func TestQuery_MatchAll(t *testing.T) {
	match := QueryApi(client).
		Indices("all_articles").
		MatchAll("", "")
	t.Logf("the terms query:\n%s", match.String())
	queryInfo, err := match.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))
}

func TestQuery_MultiMatch(t *testing.T) {
	match := QueryApi(client).
		Indices("all_articles").
		MultiMatch([]string{"title", "category_title"}, "Golang")
	t.Logf("the terms query:\n%s", match.String())
	queryInfo, err := match.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))
}

func TestQuery_MatchPhrasePrefix(t *testing.T) {
	match := QueryApi(client).
		Indices("all_articles").
		MatchPhrasePrefix("category_title", "每日新闻GoCN每日新闻(2020")
	t.Logf("the terms query:\n%s", match.String())
	queryInfo, err := match.Query(ctx)
	if err != nil {
		t.Error("the Query Query error:", err.Error())
		return
	}
	t.Logf("Query success:\n%s", util.Obj2JsonStrIndent(queryInfo, "", "  "))
}
