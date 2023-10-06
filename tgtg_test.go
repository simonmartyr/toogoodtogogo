package toogoodtogo

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

func testClient(code int, body io.Reader, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.WriteHeader(code)
		//all responses are gzipped
		gzHttpWriter := gzip.NewWriter(w)
		defer gzHttpWriter.Close()
		_, _ = io.Copy(gzHttpWriter, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := &Client{
		http:        http.DefaultClient,
		baseURL:     server.URL + "/",
		accessToken: "mockToken",
	}
	return client, server
}

func testClientFile(code int, filename string, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testClient(code, f, validators...)
}
