package main

//"github.com/elazarl/goproxy"

/*
func testProxy(t *testing.T) {
	Init()
	s := CONFIG

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	var proxyURL *url.URL
	var err error

	if p := os.Getenv("http_proxy"); p == "" {
		proxyURL, err = url.Parse("http://localhost:3128/")
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			t.Fatal(http.ListenAndServe(":3128", proxy))
		}()
	} else {
		proxyURL, err = url.Parse(p)
		if err != nil {
			t.Fatal(err)
		}
	}
	ProxySetting := http.ProxyURL(proxyURL)
	http.DefaultTransport = &http.Transport{
		Proxy: ProxySetting,
	}

	url := s.Y.URL
	t.Log("Y url:", url)
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status:", res.Status)
}
*/
