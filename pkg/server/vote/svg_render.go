package vote

import (
	"fmt"
	"net/http"
)

// XML contains the given interface object.
type SVG struct {
	Data interface{}
}

var svgContentType = []string{"image/svg+xml; charset=utf-8"}

// Render (XML) encodes the given interface object and writes data with custom ContentType.
func (r SVG) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	_, err := fmt.Fprint(w, r.Data)
	return err
}

// WriteContentType (XML) writes XML ContentType for response.
func (r SVG) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, svgContentType)
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
