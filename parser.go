package imagechecker

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// AbsImageURLs returns the absolute URL for all images in img tags.
func AbsImageURLs(srcURL string, r io.Reader) []string {
	var images []string
	for _, image := range ImageURLs(r) {
		images = append(images, absoluteImagePath(srcURL, image))
	}

	return images
}

// ImageURLs returns the URL for all images in img tags.
func ImageURLs(r io.Reader) []string {
	z := html.NewTokenizer(r)

	var images []string

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return images
		case tt == html.SelfClosingTagToken:
			t := z.Token()

			isImg := t.Data == "img"
			if !isImg {
				continue
			}

			image := getImageURL(t)
			if image == "" {
				continue
			}

			images = append(images, image)
		}
	}
}

// getImageURL checks if the img html.Token contains an src, if that's the
// case returns the URL.
func getImageURL(t html.Token) string {
	for _, i := range t.Attr {
		if i.Key == "src" || i.Key == "data-src" {
			return i.Val
		}
	}

	return ""
}

// absoluteImagePath checks if the url to an image already is absolute,
// otherwise returns the absolute one.
func absoluteImagePath(u, image string) string {
	if !strings.HasPrefix(image, "http") {
		image = baseURL(u) + image
	}

	return image
}

// baseURL returns the absolute URL for a link, if the url.Parse fails,
// returns an empty string.
func baseURL(fullURL string) string {
	u, err := url.Parse(fullURL)
	if err != nil {
		return ""
	}

	if u.Scheme == "" && u.Host == "" {
		return ""
	}

	return u.Scheme + "://" + u.Host
}
