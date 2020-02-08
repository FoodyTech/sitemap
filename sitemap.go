package sitemap

import (
	"encoding/xml"
)

// Sitemap represents a sitemap protocol.
// See: https://www.sitemaps.org/protocol.html
//
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"     json:"-"`
	Text    string   `xml:",chardata"  json:"text"`
	Xmlns   string   `xml:"xmlns,attr" json:"-"`
	URL     []URL    `xml:"url"        json:"url"`
}

// URL represents Sitemap URL entry.
//
type URL struct {
	Text             string `xml:",chardata"  json:"chardata"`
	Location         string `xml:"loc"        json:"loc"`
	LastModification string `xml:"lastmod"    json:"lastmod"`
	ChangeFrequency  string `xml:"changefreq" json:"changefreq"`
	Priority         string `xml:"priority"   json:"priority"`
}
