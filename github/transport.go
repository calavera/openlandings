package github

import (
	"net/http"
	"os"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func newTransportContext() context.Context {
	var cache httpcache.Cache = httpcache.NewMemoryCache()
	if cachePath := os.Getenv("GITHUB_CACHE_PATH"); cachePath != "" {
		cache = diskcache.New(cachePath)
	}
	tr := httpcache.NewTransport(cache)

	c := &http.Client{Transport: tr}
	return context.WithValue(context.Background(), oauth2.HTTPClient, c)
}
