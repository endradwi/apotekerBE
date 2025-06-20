package lib

import (
	"html"
	"strings"
)

func ToOrderList(raw string) string {

	lines := strings.Split(raw, "\n")
	var b strings.Builder
	b.WriteString("<ol>")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			b.WriteString("<li>")
			b.WriteString(html.EscapeString(line))
			b.WriteString("</li>")
		}
	}
	b.WriteString("</ol>")
	return b.String()
}
