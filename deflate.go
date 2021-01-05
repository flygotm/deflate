package deflate

import (
	"bytes"
	"compress/flate"
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/mime"
	"log"
	"os"
	"strings"
)

//Define delfate strcut
type deflate struct {
	logger      *log.Logger
	contentType []string
	minSize     int
	level       int
}

//New
func New() *deflate {
	return &deflate{
		logger:      log.New(os.Stdout, "[DEFLATE]", log.LstdFlags),
		contentType: defaultContentType,
		minSize:     2 << 9, //1KB
		level:       BestSpeed,
	}
}

//Type
func (d *deflate) Type() *middleware.Type {
	return middleware.TypeAfter
}

//Name
func (d *deflate) Name() string {
	return "Deflate"
}

//Method
func (d *deflate) Method() middleware.Method {
	return middleware.MethodAny
}

//Pattern
func (d *deflate) Pattern() middleware.Pattern {
	return middleware.PatternAny
}

func (d *deflate) accept(c *c.Context) bool {
	acceptEncoding := c.Request.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "deflate")
}

//Handler
func (d *deflate) Handler() func(c *c.Context) {
	return func(ctx *c.Context) {
		if d.accept(ctx) && ctx.Render().Rended() {
			odata := ctx.Render().Buffer
			if nil != odata && len(odata) >= d.minSize {
				ct := ctx.Render().ContentType
				if strings.Index(ct, ";") != -1 {
					ct = strings.TrimSpace(strings.Split(ct, ";")[0])
				}
				calls.Empty(ct, func() {
					ct = mime.BINARY
				})
				ctx.Header().Set("Vary", "Content-Encoding")
				ctx.Header().Set("Content-Encoding", "deflate")
				var buffers bytes.Buffer
				fw, err := flate.NewWriter(&buffers, d.level)
				defer fw.Close()
				calls.NNil(err, func() {
					d.logger.Println(err)
				})
				_, werr := fw.Write(odata)
				calls.NNil(werr, func() {
					d.logger.Println(werr)
				})
				fw.Flush()
				ctx.Write(buffers.Bytes())
			}
		}
		ctx.Chain()
	}
}

//ContentType
func (d *deflate) ContentType(contentType ...string) *deflate {
	d.contentType = contentType
	return d
}

//MinSize
func (d *deflate) MinSize(minSize int) *deflate {
	d.minSize = minSize
	return d
}

//Level
func (d *deflate) Level(level int) *deflate {
	d.level = level
	return d
}
