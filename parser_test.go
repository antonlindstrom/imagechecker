package imagechecker

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestBaseURL(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got:  "http://example.com",
			want: "http://example.com",
		},
		{
			got:  "http://example.com/home",
			want: "http://example.com",
		},
		{
			got:  "https://example.com",
			want: "https://example.com",
		},
		{
			got:  "http://www.example.com",
			want: "http://www.example.com",
		},
		{
			got:  "broken",
			want: "",
		},
		{
			got:  "909:/broken",
			want: "",
		},
	}

	for _, test := range tests {
		got := baseURL(test.got)
		if got != test.want {
			t.Errorf("want url = %q, got %q", test.want, got)
		}
	}
}

func TestAbsoluteImageURL(t *testing.T) {
	tests := []struct {
		url   string
		image string
		want  string
	}{
		{
			url:   "http://example.com/home",
			image: "/image.jpeg",
			want:  "http://example.com/image.jpeg",
		},
		{
			url:   "http://example.com/home",
			image: "image.jpeg",
			want:  "http://example.com/image.jpeg",
		},
		{
			url:   "http://example.com/home",
			image: "http://test.com/image.jpeg",
			want:  "http://test.com/image.jpeg",
		},
	}

	for _, test := range tests {
		got := absoluteImagePath(test.url, test.image)
		if got != test.want {
			t.Errorf("want url = %q, got %q", test.want, got)
		}
	}
}

func TestGetImageURL(t *testing.T) {
	tests := []struct {
		got  html.Token
		want string
	}{
		{
			got: html.Token{
				Attr: []html.Attribute{
					{
						Key: "src",
						Val: "/image.jpeg",
					},
				},
			},
			want: "/image.jpeg",
		},
		{
			got: html.Token{
				Attr: []html.Attribute{
					{
						Key: "data-src",
						Val: "/image.jpeg",
					},
				},
			},
			want: "/image.jpeg",
		},
		{
			got: html.Token{
				Attr: []html.Attribute{
					{
						Key: "a",
						Val: "/image1.jpeg",
					},
					{
						Key: "src",
						Val: "/image2.jpeg",
					},
				},
			},
			want: "/image2.jpeg",
		},
		{
			got: html.Token{
				Attr: []html.Attribute{
					{
						Key: "a",
						Val: "/image1.jpeg",
					},
				},
			},
			want: "",
		},
	}

	for _, test := range tests {
		got, ok := getImageURL(test.got)
		if test.want == "" && ok {
			t.Errorf("want ok = false, got true")
		}

		if got != test.want {
			t.Errorf("want url = %q, got %q", test.want, got)
		}
	}
}

func TestAbsImageURLs(t *testing.T) {
	tests := []struct {
		url  string
		body io.Reader
		want []string
	}{
		{
			url:  "http://example.com",
			body: strings.NewReader(`<html><body><img src="/image.jpeg" /></body></html>`),
			want: []string{"http://example.com/image.jpeg"},
		},
		{
			url:  "http://example.com",
			body: strings.NewReader(`<html><body><img data-src="/image.jpeg" /></body></html>`),
			want: []string{"http://example.com/image.jpeg"},
		},
	}

	for _, test := range tests {
		got := AbsImageURLs(test.url, test.body)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want urls = %+v, got %+v", test.want, got)
		}
	}
}

func TestImageURLs(t *testing.T) {
	tests := []struct {
		body io.Reader
		want []string
	}{
		{
			body: strings.NewReader(`<html><body><img src="/image.jpeg" /></body></html>`),
			want: []string{"/image.jpeg"},
		},
		{
			body: strings.NewReader(`<html><body><img data-src="/image.jpeg" /></body></html>`),
			want: []string{"/image.jpeg"},
		},
		{
			body: strings.NewReader(`<html><body><img src="/image.png"/ width="400"></body></html>`),
			want: []string{"/image.png"},
		},
	}

	for _, test := range tests {
		got := ImageURLs(test.body)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want urls = %+v, got %+v", test.want, got)
		}
	}
}
