package sitemap

import (
	"encoding/xml"
	"io"
)

// Parse sitemap from a given io.Reader.
//
func Parse(r io.Reader) (*Sitemap, error) {
	var u Sitemap
	err := xml.NewDecoder(r).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Parse sitemap from the given bytes.
//
func ParseBytes(buf []byte) (*Sitemap, error) {
	var u Sitemap
	err := xml.Unmarshal(buf, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Parse sitemap from a given string.
//
func ParseString(s string) (*Sitemap, error) {
	var u Sitemap
	err := xml.Unmarshal([]byte(s), &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
