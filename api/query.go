package api

import "encoding/json"

type QueryInfo struct {
	Took     int     `json:"took,omitempty"`
	TimedOut bool    `json:"timed_out,omitempty"`
	Shards   *Shards `json:"_shards,omitempty"`
	Hits     *Hits   `json:"hits,omitempty"`
}

type Shards struct {
	Total      int `json:"total,omitempty"`
	Successful int `json:"successful,omitempty"`
	Skipped    int `json:"skipped,omitempty"`
	Failed     int `json:"failed,omitempty"`
}

type Total struct {
	Value    int    `json:"value,omitempty"`
	Relation string `json:"relation,omitempty"`
}

type Hits struct {
	Total    *Total      `json:"total,omitempty"`
	MaxScore float64     `json:"max_score,omitempty"`
	Hits     []*InnerHit `json:"hits,omitempty"`
}

type InnerHit struct {
	Index     string          `json:"_index,omitempty"`
	Type      string          `json:"_type,omitempty"`
	ID        string          `json:"_id,omitempty"`
	Score     float64         `json:"_score,omitempty"`
	Source    json.RawMessage `json:"_source,omitempty"`
	Highlight *HighLight      `json:"highlight,omitempty"`
}

type HighLight struct {
	Title []string `json:"title,omitempty"`
}
