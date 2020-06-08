// Go Workshop - Practise - Gube - HTTPX

package httpx

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Header fields and values.
const (
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"

	ContentTypePlain      = "text/plain"
	ContentTypeHTML       = "text/html"
	ContentTypeJSON       = "application/json"
	ContentTypeXML        = "application/xml"
	ContentTypeURLEncoded = "application/x-www-form-urlencoded"

	contentTypesText = "text/"
)

// AcceptsContentType checks a request header accepts a given content type.
func AcceptsContentType(h http.Header, contentType string) bool {
	return strings.Contains(h.Get(HeaderAccept), contentType)
}

// ContainsContentType checks if the header contains the given content type.
func ContainsContentType(h http.Header, contentType string) bool {
	return strings.Contains(h.Get(HeaderContentType), contentType)
}

// PathParts splits the request path into its parts.
func PathParts(r *http.Request) []string {
	rawParts := strings.Split(r.URL.Path, "/")
	parts := []string{}
	for _, part := range rawParts {
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

// PathAt returns the nth part of the request path and true
// if it exists. Otherwise an empty string and false.
func PathAt(r *http.Request, n int) (string, bool) {
	if n < 0 {
		panic("illegal path index")
	}
	parts := PathParts(r)
	if len(parts) < n+1 {
		return "", false
	}
	return parts[n], true
}

// ReadBody retrieves the whole body out of a HTTP request or response
// and returns it as byte slice.
func ReadBody(r io.ReadCloser) ([]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %v", err)
	}
	r.Close()
	return data, nil
}

// WriteBody writes the whole body into a HTTP request or response.
func WriteBody(w io.Writer, data []byte) error {
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("cannot write body: %v", err)
	}
	return nil
}

// UnmarshalBody parses the body data of a request or response based on the
// content type header stores the result in the value pointed by v. Conten types
// JSON and XML expect the according types as arguments, all text types
// expect a string, and all others too, but the data is encoded in BASE64.
func UnmarshalBody(r io.ReadCloser, h http.Header, v interface{}) error {
	// First read.
	data, err := ReadBody(r)
	if err != nil {
		return err
	}
	// Then unmarshal based on content type.
	switch {
	case ContainsContentType(h, contentTypesText):
		switch tv := v.(type) {
		case *string:
			*tv = string(data)
		case *interface{}:
			*tv = string(data)
		default:
			return errors.New("invalid value argument for text; need string or empty interface")
		}
		return nil
	case ContainsContentType(h, ContentTypeJSON):
		if err = json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("cannot unmarshal JSON: %v", err)
		}
		return nil
	case ContainsContentType(h, ContentTypeXML):
		if err = xml.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("cannot unmarshal XML: %v", err)
		}
		return nil
	default:
		sd := base64.StdEncoding.EncodeToString(data)
		switch tv := v.(type) {
		case *string:
			*tv = sd
		case *interface{}:
			*tv = sd
		default:
			return errors.New("invalid value argument for base64; need string or empty interface")
		}
		return nil
	}
}

// MarshalBody allows to directly marshal a value into a writer depending on
// the content type.
func MarshalBody(w io.Writer, h http.Header, v interface{}) error {
	// First marshal based on content type.
	var data []byte
	var err error
	switch {
	case ContainsContentType(h, contentTypesText):
		data = []byte(fmt.Sprintf("%v", v))
	case ContainsContentType(h, ContentTypeJSON):
		data, err = json.Marshal(v)
		if err != nil {
			return fmt.Errorf("cannot marshal to JSON: %v", err)
		}
	case ContainsContentType(h, ContentTypeXML):
		data, err = xml.Marshal(v)
		if err != nil {
			return fmt.Errorf("cannot marshal to XML: %v", err)
		}
	default:
		vbs, ok := v.([]byte)
		if !ok {
			return errors.New("invalid value argument")
		}
		data = vbs
	}
	// Then write the body.
	return WriteBody(w, data)
}
