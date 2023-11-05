// Package link provides the core models of the ShortLink application.
package link

import (
	"net/url"
	"time"
)

// Link is the central model of the application.
// It represents a shortened link.
type Link struct {
	// ID is the unique identifier of the link.
	ID uint64

	// Key is the key of the link.
	// This is the value that is used to access the link.
	Key string

	// URL is the URL the link redirects to.
	URL url.URL

	// CreateTime is the time the link was created.
	CreateTime time.Time

	// UpdateTime is the time the link was last updated.
	UpdateTime time.Time
}
