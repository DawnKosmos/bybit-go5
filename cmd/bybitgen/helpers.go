package main

import "strings"

// normalizeParamName trims and removes leading nesting markers.
// Returns (name, nested) where nested is true if the row appears to describe a nested field.
func normalizeParamName(s string) (string, bool) {
	n := strings.TrimSpace(s)
	if n == "" {
		return n, false
	}
	// Rows like "> coin" or "> collateralSwitch" denote nested fields in docs
	if strings.HasPrefix(n, ">") {
		// strip leading '>' and whitespace
		n = strings.TrimSpace(strings.TrimLeft(n, ">"))
		return n, true
	}
	// Remove leading bullets or dashes
	for strings.HasPrefix(n, "-") || strings.HasPrefix(n, "•") {
		n = strings.TrimSpace(n[1:])
	}
	// Remove trailing colon
	n = strings.TrimSuffix(n, ":")
	return n, false
}

// normalizeRespName returns the cleaned parameter name and its nesting depth based on leading '>' markers.
// depth: 0 = top-level, 1 = child ('> field'), 2 = grandchild ('>> field'), etc.
func normalizeRespName(s string) (string, int) {
    t := strings.TrimSpace(stripFormatting(s))
    if t == "" {
        return "", 0
    }
    depth := 0
    for i := 0; i < len(t) && t[i] == '>'; i++ {
        depth++
    }
    // trim leading '>' and whitespace after them
    t = strings.TrimSpace(strings.TrimLeft(t, ">"))
    // Remove trailing colon
    t = strings.TrimSuffix(t, ":")
    return t, depth
}

// parseTableRow splits a Markdown table row "| a | b | c |" into cells.
func parseTableRow(line string) []string {
	if strings.HasPrefix(line, "|") {
		line = line[1:]
	}
	if strings.HasSuffix(line, "|") {
		line = line[:len(line)-1]
	}
	parts := strings.Split(line, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// stripFormatting removes simple Markdown formatting like links [text](url) and backticks.
func stripFormatting(s string) string {
	// remove backticks
	s = strings.ReplaceAll(s, "`", "")
	// remove links: [text](...)
	for {
		start := strings.Index(s, "[")
		mid := strings.Index(s, "](")
		end := strings.Index(s, ")")
		if start >= 0 && mid > start && end > mid {
			s = s[:start] + s[start+1:mid] + s[end+1:]
			continue
		}
		break
	}
	return strings.TrimSpace(s)
}

// cleanComment flattens whitespace and strips a few common HTML tags that appear in docs.
func cleanComment(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	repl := []string{"<li>", "</li>", "<ul>", "</ul>", "<br/>", "<br>", "<p>", "</p>"}
	for _, r := range repl {
		s = strings.ReplaceAll(s, r, " ")
	}
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

// mapType maps docs types to Go types for request params, following old_models style.
func mapType(t string) string {
	tt := strings.ToLower(strings.TrimSpace(t))
	switch tt {
	case "string", "text":
		return "string"
	case "integer", "int", "int32", "int64", "long", "timestamp(ms)", "timestamp":
		return "int64"
	case "boolean", "bool":
		return "bool"
	case "number", "float", "double", "decimal", "object", "array":
		// default to string for ambiguous numeric/object types in requests
		return "string"
	default:
		// fallback
		return "string"
	}
}

// toExported converts a param name like "baseCoin" or "base-coin" to an exported Go field name.
func toExported(name string) string {
	n := strings.TrimSpace(name)
	if n == "" {
		return ""
	}
	// replace separators with spaces
	seps := func(r rune) bool { return r == '-' || r == '_' || r == ' ' }
	parts := strings.FieldsFunc(n, seps)
	if len(parts) == 1 {
		// handle camelCase: uppercase the first rune
		return strings.ToUpper(parts[0][:1]) + parts[0][1:]
	}
	var sb strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		sb.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			sb.WriteString(p[1:])
		}
	}
	out := sb.String()
	// if starts with digit, prefix with X
	if len(out) > 0 && out[0] >= '0' && out[0] <= '9' {
		out = "X" + out
	}
	return out
}

// makeName returns a Go-style method/struct name based on HTTP method and path
func makeName(method, path string) string {
	// drop leading /v5/
	p := strings.TrimPrefix(path, "/v5/")
	parts := strings.Split(p, "/")
	var sb strings.Builder
	// Method-based prefix
	if method == "GET" {
		sb.WriteString("Get")
	} else {
		sb.WriteString("Post")
	}
	for _, part := range parts {
		if part == "" {
			continue
		}
		for _, seg := range strings.Split(part, "-") {
			sb.WriteString(strings.Title(seg))
		}
	}
	return sb.String()
}
