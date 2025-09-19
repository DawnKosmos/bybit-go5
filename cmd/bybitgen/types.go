package main

import "regexp"

// Endpoint represents a Bybit API endpoint discovered from docs
// and the parameters extracted from its Request Parameters table.
type Endpoint struct {
	Folder string
	Method string // GET or POST
	Path   string // e.g. /v5/position/list
	Name   string // Go-friendly base name, e.g. GetPositionList
	Params []Param // request params

	// Response schema parsed from "Response Parameters" table
	// Top-level fields in result object (e.g., category, nextPageCursor, list)
	RespTop []Param
	// Nested fields keyed by their parent top-level field name (e.g., "list" -> [...] )
	RespNested map[string][]Param

	// Category-dependent response schemas (for endpoints like tickers)
	// Key is category name (e.g., "linear", "spot", "option")
	CategoryResponses map[string]CategoryResponse
	HasCategoryParam  bool // true if this endpoint has a "category" parameter
}

// CategoryResponse represents response schema for a specific category
type CategoryResponse struct {
	Category   string            // e.g., "linear", "spot", "option"
	RespTop    []Param          // top-level fields for this category
	RespNested map[string][]Param // nested fields for this category
}

// Param describes a single request parameter parsed from docs.
type Param struct {
	Name     string
	Required bool
	Type     string
	Comment  string
}

// httpLine matches lines like:  GET `/v5/position/list`
var httpLine = regexp.MustCompile(`^(GET|POST)\s+` + "`" + `([^` + "`" + `]+)` + "`")

// tabItemLine matches TabItem tags like: <TabItem value="linear" label="Linear/Inverse" default>
var tabItemLine = regexp.MustCompile(`<TabItem\s+value="([^"]+)"\s+label="([^"]+)".*?>`)

// tabItemEndLine matches closing TabItem tags
var tabItemEndLine = regexp.MustCompile(`</TabItem>`)
