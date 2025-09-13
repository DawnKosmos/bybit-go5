package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// scanEndpoints walks a docs directory, extracting endpoints and their request params
func scanEndpoints(dir string) ([]Endpoint, error) {
	var eps []Endpoint
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".md") && !strings.HasSuffix(d.Name(), ".mdx") {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		s := bufio.NewScanner(f)
		var current *Endpoint
		inReqTable := false
		inRespTable := false
		// Maintain a stack of parent names for response nesting. Index = depth.
		var respPath []string
		for s.Scan() {
			line := strings.TrimSpace(s.Text())
			// Detect HTTP line
			if m := httpLine.FindStringSubmatch(line); len(m) == 3 {
				method := m[1]
				p := m[2]
				if strings.HasPrefix(p, "/v5/") {
					eps = append(eps, Endpoint{Method: method, Path: p})
					current = &eps[len(eps)-1]
				} else {
					current = nil
				}
				inReqTable = false
				inRespTable = false
				respPath = nil
				continue
			}
			// Detect start of Request Parameters section
			if strings.HasPrefix(line, "### ") && strings.Contains(strings.ToLower(line), "request parameters") {
				inReqTable = true
				inRespTable = false
				continue
			}
			// Detect start of Response Parameters section
			if strings.HasPrefix(line, "### ") && strings.Contains(strings.ToLower(line), "response parameters") {
				inRespTable = true
				inReqTable = false
				respPath = nil
				continue
			}
			if inReqTable {
				if !strings.HasPrefix(line, "|") {
					// end of table
					inReqTable = false
					continue
				}
				// skip header and delimiter rows
				if strings.Contains(line, ":-----") || strings.Contains(strings.ToLower(line), "parameter | required | type") {
					continue
				}
				if current == nil {
					continue
				}
				cells := parseTableRow(line)
				if len(cells) < 4 {
					continue
				}
				rawName := stripFormatting(cells[0])
				name, nested := normalizeParamName(rawName)
				if nested {
					// nested object fields are not modeled at request top-level for now
					continue
				}
				// Conservative: skip names containing spaces (usually formatting artifacts)
				if strings.Contains(name, " ") || name == "" || strings.EqualFold(name, "-") {
					continue
				}
				p := Param{
					Name:     name,
					Required: strings.Contains(strings.ToLower(cells[1]), "true") || strings.Contains(strings.ToLower(cells[1]), "yes"),
					Type:     strings.ToLower(stripFormatting(cells[2])),
					Comment:  strings.TrimSpace(cells[3]),
				}
				// skip empty parameter rows
				if p.Name == "" || strings.EqualFold(p.Name, "-") {
					continue
				}
				current.Params = append(current.Params, p)
			}
			if inRespTable {
				if !strings.HasPrefix(line, "|") {
					// end of table
					inRespTable = false
					respPath = nil
					continue
				}
				// skip header and delimiter rows
				low := strings.ToLower(line)
				if strings.Contains(line, ":-----") || strings.Contains(low, "parameter | type") || strings.Contains(low, "parameter|type") {
					continue
				}
				if current == nil {
					continue
				}
				if current.RespNested == nil {
					current.RespNested = make(map[string][]Param)
				}
				cells := parseTableRow(line)
				// Response table typically has 3 columns: name | type | comments
				if len(cells) < 3 {
					continue
				}
				rawName := cells[0]
				name, depth := normalizeRespName(rawName)
				typ := strings.ToLower(stripFormatting(cells[1]))
				comment := strings.TrimSpace(cells[2])
				if name == "" || strings.EqualFold(name, "-") {
					continue
				}
				p := Param{
					Name:    name,
					Type:    typ,
					Comment: comment,
				}
				if depth == 0 {
					// top-level
					current.RespTop = append(current.RespTop, p)
					// reset and set current path to this top-level name
					respPath = []string{name}
					continue
				}
				// nested: ensure we have a valid parent path of at least 'depth' elements
				if len(respPath) < depth {
					// cannot resolve parent, skip
					continue
				}
				// parent key is join of first 'depth' elements of path
				parentKey := strings.Join(respPath[:depth], ".")
				current.RespNested[parentKey] = append(current.RespNested[parentKey], p)
				// update current path to include this node at its depth level
				if len(respPath) >= depth {
					respPath = append(respPath[:depth], name)
				}
			}
		}
		return nil
	})
	return eps, err
}
