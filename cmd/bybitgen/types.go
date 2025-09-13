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
