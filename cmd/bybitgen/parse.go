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
				continue
			}
			// Detect start of Request Parameters section
			if strings.HasPrefix(line, "### ") && strings.Contains(strings.ToLower(line), "request parameters") {
				inReqTable = true
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
		}
		return nil
	})
	return eps, err
}
