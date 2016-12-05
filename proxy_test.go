package main

import (
	"github.com/elazarl/goproxy"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {

	s := NewSettings()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxyURL, err := url.Parse("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Fatal(http.ListenAndServe(":8080", proxy))
	}()
	ProxySetting := http.ProxyURL(proxyURL)
	http.DefaultTransport = &http.Transport{
		Proxy: ProxySetting,
	}

	url := s.Y_PAGE_URL
	t.Log("Y url:", url)
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status:", res.Status)
}
