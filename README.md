# sitemap

[![Build Status][build-img]][build-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]

Sitemap protocol library in Go.

## Install

Go version 1.13

```
go get github.com/FoodyTech/sitemap
```

## Example

```go
mysite := `<some sitemap in XML>`

smap, err := sitemap.ParseString(mysite)
if err != nil {
    // process on error
}

for _, url := smap.URL {
    println(url.Location)
    println(url.LastModification)
    println()
}
```

## Documentation

See [these docs](https://godoc.org/github.com/FoodyTech/sitemap).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/FoodyTech/sitemap/workflows/build/badge.svg
[build-url]: https://github.com/FoodyTech/sitemap/actions
[doc-img]: https://godoc.org/github.com/FoodyTech/sitemap?status.svg
[doc-url]: https://godoc.org/github.com/FoodyTech/sitemap
[reportcard-img]: https://goreportcard.com/badge/FoodyTech/sitemap
[reportcard-url]: https://goreportcard.com/report/Foody
