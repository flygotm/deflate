package deflate

import "compress/flate"

const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

var defaultContentType = []string{
	"application/javascript",
	"application/json",
	"application/xml",
	"text/javascript",
	"text/json",
	"text/xml",
	"text/plain",
	"text/xml",
	"html/css",
}
