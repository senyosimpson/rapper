package spotify

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func newTestServer(code int, body io.Reader) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		io.Copy(w, body)
		r.Body.Close()
		// Possible I may need to do error checking here in the future
		body.(io.Closer).Close()
	}))
	return server
}

func newTestClient(url string) *Client {
	client := &Client{
		client: http.DefaultClient,
		baseURL: url,
	}
	return client
}