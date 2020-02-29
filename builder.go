package sitemap

// Builder helps to build own site map.
type Builder struct {
	s Sitemap
}

// Add a given loc to the site map.
func (b *Builder) Add(loc string) {
	b.s.URL = append(b.s.URL, URL{Location: loc})
}

// Add a given URL to the site map.
func (b *Builder) AddURL(url URL) {
	b.s.URL = append(b.s.URL, url)
}

// Build the site map.
func (b *Builder) Build() Sitemap {
	return b.s
}
