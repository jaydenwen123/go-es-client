package api

//Query term查询
func (t *Query) Term(field string, value interface{}) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Term: map[string]interface{}{
					field: value,
				},
			},
		},
	}
}

//Terms terms查询
func (t *Query) Terms(field string, values ...interface{}) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Terms: map[string]interface{}{
					field: values,
				},
			},
		},
	}
}

//Ids ids查询
func (t *Query) Ids(values ...string) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Ids: map[string][]string{
					"values": values,
				},
			},
		},
	}
}

//Exists exists查询
func (t *Query) Exists(field string) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Exists: map[string]interface{}{
					"field": field,
				},
			},
		},
	}
}

//Prefix prefix查询
func (t *Query) Prefix(field string, value interface{}) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Prefix: map[string]interface{}{
					field: map[string]interface{}{
						"value": value,
					},
				},
			},
		},
	}
}

//Range range查询
func (t *Query) Range(field string, value *RangeCond) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Range: map[string]*RangeCond{
					field: value,
				},
			},
		},
	}
}

//Regexp regexp查询
func (t *Query) Regexp(field string, value *RegexpCond) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Regexp: map[string]*RegexpCond{
					field: value,
				},
			},
		},
	}
}

//Wildcard wildcard查询
func (t *Query) Wildcard(field string, value *WildcardCond) *QueryProxy {
	return &QueryProxy{
		term: t,
		QueryCond: &queryCondition{
			Query: &queryCond{
				Wildcard: map[string]*WildcardCond{
					field: value,
				},
			},
		},
	}
}
