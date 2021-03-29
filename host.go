package kkhost

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Host struct {
	DocumentSite string
	Production   string
	Staging      string
	Development  string
	Testing      string
}

func (h *Host) Now() string {
	switch strings.ToUpper(os.Getenv("KKAPP_ENVIRONMENT")) {
	case "PRODUCTION":
		return h.Production
	case "STAGING":
		return h.Staging
	case "DEVELOPMENT":
		return h.Development
	case "TESTING":
		return h.Testing
	default:
		return h.Production
	}
}

func (h *Host) NowHttps() string {
	return fmt.Sprintf("https://%s", h.Now())
}

func (h *Host) NowURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   h.Now(),
	}
}
