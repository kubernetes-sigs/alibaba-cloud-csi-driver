package dara

import (
	"fmt"
	"net/url"
	"strings"
)

// PortMap maps protocols to their corresponding ports.
var portMap = map[string]string{
	"ftp":    "21",
	"gopher": "70",
	"http":   "80",
	"https":  "443",
	"ws":     "80",
	"wss":    "443",
}

// URL is a wrapper around the URL type.
type URL struct {
	_url *url.URL
}

// NewURL constructs a new URL from a string.
func NewURL(str string) (*URL, error) {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	return &URL{_url: parsedURL}, nil
}

// Path returns the path and query of the URL.
func (t *URL) Path() string {
	if t._url.RawQuery == "" {
		return t._url.Path
	}
	return t._url.Path + "?" + t._url.RawQuery
}

// Pathname returns the pathname of the URL.
func (t *URL) Pathname() string {
	return t._url.Path
}

// Protocol returns the protocol of the URL.
func (t *URL) Protocol() string {
	return strings.TrimSuffix(t._url.Scheme, ":")
}

// Hostname returns the hostname of the URL.
func (t *URL) Hostname() string {
	return t._url.Hostname()
}

// Host returns the host (host:port) of the URL.
func (t *URL) Host() string {
	return t._url.Host
}

// Port returns the port of the URL, or the default for the protocol if not specified.
func (t *URL) Port() string {
	if p := t._url.Port(); p != "" {
		return p
	}
	return portMap[t.Protocol()]
}

// Hash returns the hash of the URL without the #.
func (t *URL) Hash() string {
	return strings.TrimPrefix(t._url.Fragment, "#")
}

// Search returns the search part of the URL without the ?.
func (t *URL) Search() string {
	return strings.TrimPrefix(t._url.RawQuery, "?")
}

// Href returns the full URL.
func (t *URL) Href() string {
	return t._url.String()
}

// Auth returns the username and password of the URL.
func (t *URL) Auth() string {
	password := getPassword(t._url.User)
	username := t._url.User.Username()
	if username == "" && password == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s", username, password)
}

// getPassword retrieves the password from a URL.User.
func getPassword(user *url.Userinfo) string {
	if password, ok := user.Password(); ok {
		return password
	}
	return ""
}

// Parse constructs a new URL from a string.
func ParseURL(urlStr string) (*URL, error) {
	return NewURL(urlStr)
}

// EncodeURL encodes a URL string.
func EncodeURL(urlStr string) string {
	if urlStr == "" {
		return ""
	}
	strs := strings.Split(urlStr, "/")
	for i, v := range strs {
		strs[i] = url.QueryEscape(v)
	}
	urlStr = strings.Join(strs, "/")
	urlStr = strings.Replace(urlStr, "+", "%20", -1)
	urlStr = strings.Replace(urlStr, "*", "%2A", -1)
	urlStr = strings.Replace(urlStr, "%7E", "~", -1)
	return urlStr
}

// PercentEncode encodes a string for use in URLs, replacing certain characters.
func PercentEncode(uri string) string {
	if uri == "" {
		return ""
	}
	uri = url.QueryEscape(uri)
	uri = strings.Replace(uri, "+", "%20", -1)
	uri = strings.Replace(uri, "*", "%2A", -1)
	uri = strings.Replace(uri, "%7E", "~", -1)
	return uri
}

// PathEncode encodes each segment of a path.
func PathEncode(path string) string {
	if path == "" || path == "/" {
		return path
	}
	strs := strings.Split(path, "/")
	for i, v := range strs {
		strs[i] = url.QueryEscape(v)
	}
	path = strings.Join(strs, "/")
	path = strings.Replace(path, "+", "%20", -1)
	path = strings.Replace(path, "*", "%2A", -1)
	path = strings.Replace(path, "%7E", "~", -1)
	return path
}
