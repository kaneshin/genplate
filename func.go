package genplate

import (
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

var (
	funcMap = template.FuncMap{
		"initialCap": initialCap,
		"initialLow": initialLow,
		"methodCap":  methodCap,
		"methodLow":  methodLow,
	}
)

// FuncMap returns a collention of the template's function map.
func FuncMap() template.FuncMap {
	return funcMap
}

var (
	newlines      = regexp.MustCompile(`(?m:\s*$)`)
	camelcase     = regexp.MustCompile(`(?m)[-.$/:_{}\s]`)
	acronyms      = regexp.MustCompile(`(Url|Http|Id|Io|Uuid|Guid|Api|Uri|Ssl|Cname|Oauth|Otp)([^a-z]|$)`)
	acronymMapper = map[string]string{
		"Url":   "URL",
		"Http":  "HTTP",
		"Id":    "ID",
		"Io":    "IO",
		"Uuid":  "UUID",
		"Guid":  "GUID",
		"Api":   "API",
		"Uri":   "URI",
		"Ssl":   "SSL",
		"Cname": "CNAME",
		"Oauth": "OAuth",
		"Otp":   "OTP",
	}
)

func contains(n string, r []string) bool {
	for _, r := range r {
		if r == n {
			return true
		}
	}
	return false
}

func initialCap(ident string) string {
	if ident == "" {
		// blank identifier
		return ""
	}
	return depunct(ident, true)
}

func initialLow(ident string) string {
	if ident == "" {
		// blank identifier
		return ""
	}
	return depunct(ident, false)
}

func methodCap(ident string) string {
	return initialCap(strings.ToLower(ident))
}

func methodLow(ident string) string {
	return initialLow(strings.ToLower(ident))
}

func depunct(ident string, initialCap bool) string {
	matches := camelcase.Split(ident, -1)
	for i, m := range matches {
		if initialCap || i > 0 {
			m = capFirst(m)
		}
		matches[i] = acronyms.ReplaceAllStringFunc(m, func(c string) string {
			if v, ok := acronymMapper[c]; ok {
				return v
			}
			return c
		})
	}
	str := strings.Join(matches, "")
	if !initialCap {
		str = lowFirst(str)
	}
	return str
}

func capFirst(ident string) string {
	r, n := utf8.DecodeRuneInString(ident)
	return string(unicode.ToUpper(r)) + ident[n:]
}

func lowFirst(ident string) string {
	r, n := utf8.DecodeRuneInString(ident)
	return string(unicode.ToLower(r)) + ident[n:]
}
