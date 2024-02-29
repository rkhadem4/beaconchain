package types

import (
	"time"

	"github.com/shopspring/decimal"
)

// frontend can ignore ApiResponse type, it's just for the backend

type Paging struct {
	PrevCursor string `json:"prev_cursor,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
	TotalCount uint64 `json:"total_count,omitempty"`
}
type ApiResponse struct {
	Paging *Paging     `json:"paging,omitempty"`
	Data   interface{} `json:"data"`
}

type ApiErrorResponse struct {
	Error string `json:"error"`
}

type ApiDataResponse[T any] struct {
	Data T `json:"data"`
}
type ApiPagingResponse[T any] struct {
	Paging Paging `json:"paging"`
	Data   []T    `json:"data"`
}

type PubKey string
type Hash string // blocks, txs etc.

type Address struct {
	Hash Hash   `json:"hash"`
	Ens  string `json:"ens,omitempty"`
}
type LuckItem struct {
	Percent  float64       `json:"percent"`
	Expected time.Time     `json:"expected"`
	Average  time.Duration `json:"average"`
}

type Luck struct {
	Proposal LuckItem `json:"proposal"`
	Sync     LuckItem `json:"sync"`
}

type StatusCount struct {
	Success uint64 `json:"success"`
	Failed  uint64 `json:"failed"`
}

type ClElUnion interface {
	float64 | decimal.Decimal
}

type ClElValue[T ClElUnion] struct {
	El T `json:"el"`
	Cl T `json:"cl"`
}

type PeriodicClElValues[T ClElUnion] struct {
	Total ClElValue[T] `json:"total"`
	Day   ClElValue[T] `json:"day"`
	Week  ClElValue[T] `json:"week"`
	Month ClElValue[T] `json:"month"`
	Year  ClElValue[T] `json:"year"`
}

type ChartSeries[T int | string] struct {
	Id    T         `json:"id"`              // id may be a string or an int
	Stack string    `json:"stack,omitempty"` // for stacking bar charts
	Data  []float64 `json:"data"`            // y-axis values
}

type ChartData[T int | string] struct {
	Categories []uint64         `json:"categories"` // x-axis
	Series     []ChartSeries[T] `json:"series"`
}

type SearchResult struct {
	Type      string `json:"type"`
	ChainId   uint64 `json:"chain_id"`
	HashValue string `json:"hash_value,omitempty"`
	NumValue  uint64 `json:"num_value,omitempty"`
	StrValue  string `json:"str_value,omitempty"`
}

type SearchResponse struct {
	Data []SearchResult `json:"data"`
}

type Dashboard struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type DashboardData struct {
	ValidatorDashboards []Dashboard `json:"validator_dashboards"`
	AccountDashboards   []Dashboard `json:"account_dashboards"`
}