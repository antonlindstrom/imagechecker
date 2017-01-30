package imagechecker

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCheckLink(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
	}))
	defer ts.Close()

	tests := []struct {
		want *Link
	}{
		{
			want: &Link{
				URL:          ts.URL,
				ContentType:  "image/png",
				ResponseCode: 200,
				ETag:         "",
			},
		},
	}

	for _, test := range tests {
		client := New()

		ch := make(chan *Link, 1)
		client.checkLink(ts.URL, ch)

		got := <-ch

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want Link = %+v, got %+v", test.want, got)
		}
	}
}

func TestLinksFromURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
	}))
	defer ts.Close()

	tests := []struct {
		want []*Link
	}{
		{
			want: []*Link{
				{
					URL:          ts.URL,
					ContentType:  "image/png",
					ResponseCode: 200,
					ETag:         "",
				},
			},
		},
	}

	for _, test := range tests {
		client := New()

		got := client.linksFromURL([]string{ts.URL})

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want []Link = %+v, got %+v", test.want, got)
		}
	}
}

func TestDocumentFromReader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
	}))
	defer ts.Close()

	tests := []struct {
		want *Document
	}{
		{
			want: &Document{
				Links: []*Link{
					{
						URL:          ts.URL,
						ContentType:  "image/png",
						ResponseCode: 200,
						ETag:         "",
					},
				},
			},
		},
	}

	for _, test := range tests {
		client := &Client{
			Client: http.DefaultClient,
			URLParser: func(string, io.Reader) []string {
				return []string{ts.URL}
			},
		}

		got, err := client.DocumentFromURL(ts.URL)
		if err != nil {
			t.Errorf("want err = nil, got %v", err)
			continue
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want []Link = %+v, got %+v", test.want, got)
		}
	}
}

func TestDocumentFromURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
	}))
	defer ts.Close()

	tests := []struct {
		want *Document
	}{
		{
			want: &Document{
				Links: []*Link{
					{
						URL:          ts.URL,
						ContentType:  "image/png",
						ResponseCode: 200,
						ETag:         "",
					},
				},
			},
		},
	}

	for _, test := range tests {
		client := &Client{
			Client: http.DefaultClient,
			URLParser: func(string, io.Reader) []string {
				return []string{ts.URL}
			},
		}

		got := client.DocumentFromReader(ts.URL, strings.NewReader(""))

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want []Link = %+v, got %+v", test.want, got)
		}
	}
}
