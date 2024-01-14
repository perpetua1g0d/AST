package pkg

import "strings"

func ReplaceQuotesWithEscape(s string) string {
	startsWithQuotes := s[0] == '"'
	endsWithQuotes := s[0] == '"'
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")

	s = strings.Replace(s, "\"", "\\\"", -1)

	if startsWithQuotes {
		s = "\"" + s
	}
	if endsWithQuotes {
		s += "\""
	}
	return s
}
