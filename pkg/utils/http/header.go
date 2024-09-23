package http

import (
	"bufio"
	"bytes"
	"maps"
	"net/http"
	"net/textproto"
	"os"

	"k8s.io/klog/v2"
)

func ParseHeaderConfig(hstr string) (http.Header, error) {
	reader := bytes.NewBufferString(hstr)
	reader.Write([]byte("\n\n"))
	h, err := textproto.NewReader(bufio.NewReader(reader)).ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	return http.Header(h), nil
}

func mustParseHeaderTo(envName string, header http.Header) {
	if hstr := os.Getenv(envName); hstr != "" {
		h, err := ParseHeaderConfig(hstr)
		if err != nil {
			klog.Fatalf("Invaild %s: %v", envName, err)
		}
		maps.Copy(header, h)
	}
}

func MustParseHeaderEnv(envName string) http.Header {
	header := http.Header{}
	mustParseHeaderTo("ALIBABA_CLOUD_HTTP_HEADERS", header)
	mustParseHeaderTo(envName, header)
	if ascmContext := os.Getenv("X-ACSPROXY-ASCM-CONTEXT"); ascmContext != "" {
		// for private cloud, capability reason only
		header.Set("X-Acsproxy-Ascm-Context", ascmContext)
	}
	return header
}

type InsertHeaderRoundTripper struct {
	http.RoundTripper
	Header http.Header
}

func (rt *InsertHeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	maps.Copy(req.Header, rt.Header)
	return rt.RoundTripper.RoundTrip(req)
}

func RoundTripperWithHeader(rt http.RoundTripper, header http.Header) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &InsertHeaderRoundTripper{
		RoundTripper: rt,
		Header:       header,
	}
}

func MustToV2SDKHeaders(headers http.Header) map[string]*string {
	headersV2 := make(map[string]*string, len(headers))
	for k, v := range headers {
		if len(v) == 0 {
			continue
		}
		if len(v) > 1 {
			klog.Fatalf("Multi-value header %s is not supported", k)
		}
		headersV2[k] = &v[0]
	}
	return headersV2
}
