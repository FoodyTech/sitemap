package sitemap_test

import (
	"strings"
	"testing"

	"github.com/FoodyTech/sitemap"
)

var testXML = buildSitemapXML([]string{"https://foo.bar", "http://bar.baz"})

func TestParse(t *testing.T) {
	r := strings.NewReader(testXML)

	sm, err := sitemap.Parse(r)
	if err != nil {
		t.Error(err)
	}

	if want := 2; len(sm.URL) != want {
		t.Errorf("got %#v, want %#v", len(sm.URL), want)
	}
}

func TestParseBytes(t *testing.T) {
	sm, err := sitemap.ParseBytes([]byte(testXML))
	if err != nil {
		t.Error(err)
	}

	if want := 2; len(sm.URL) != want {
		t.Errorf("got %#v, want %#v", len(sm.URL), want)
	}
}

func TestParseString(t *testing.T) {
	sm, err := sitemap.ParseString(testXML)
	if err != nil {
		t.Error(err)
	}

	if want := 2; len(sm.URL) != want {
		t.Errorf("got %#v, want %#v", len(sm.URL), want)
	}
}

func buildSitemapXML(urls []string) string {
	const header = `<?xml version="1.0" encoding="UTF-8"?>` +
		`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">`
	const footer = "</urlset>"

	sb := strings.Builder{}
	sb.WriteString(header)
	for _, u := range urls {
		sb.WriteString("<url><loc>")
		sb.WriteString(u)
		sb.WriteString("</loc></url>")
		sb.WriteString("\n")
	}
	sb.WriteString(footer)
	return sb.String()
}
