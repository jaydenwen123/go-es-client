package api

type MatchQuery struct {
	term      *Query          `json:"-"`
	TermCond  *queryCondition `json:"term_cond"`
	TermsCond *queryCondition `json:"terms_cond"`
}

//Match match查询
func (t *Query) Match(field string, value interface{}) *QueryProxy {

	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Match: t.genQuery(field, value),
			},
		},
	}
}

func (t *Query) genQuery(field string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		field: value,
	}
}

//MatchPhrase match_phrase查询
func (t *Query) MatchPhrase(field string, value interface{}) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				MatchPhrase: t.genQuery(field, value),
			},
		},
	}
}

//MatchPhrasePrefix match_phrase_prefix查询
func (t *Query) MatchPhrasePrefix(field string, value interface{}) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				MatchPhrasePrefix: t.genQuery(field, value),
			},
		},
	}
}

//MultiMatch multi_match查询
func (t *Query) MultiMatch(fields []string, value interface{}, ) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				MultiMatch: map[string]interface{}{
					"query":  value,
					"fields": fields,
				},
			},
		},
	}
}

//MultiMatch multi_match查询
func (t *Query) QueryString(fields []string, value interface{}, ) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				QueryString: map[string]interface{}{
					"query":  value,
					"fields": fields,
				},
			},
		},
	}
}

//MatchAll match_all查询
func (t *Query) MatchAll(field string, value interface{}) *QueryProxy {
	var query interface{}
	if len(field) > 0 {
		query = t.genQuery(field, value)
	} else {
		query = struct{}{}
	}
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				MatchAll: query,
			},
		},
	}
}

//MatchNone match_none查询
func (t *Query) MatchNone() *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				MatchNone: struct{}{},
			},
		},
	}
}
