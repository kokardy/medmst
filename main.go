package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
)

var (
	force       bool //force is a option to overwrite files.
	downloadHOT bool
	downloadY   bool
)

//Init is a initial function for global variables.
func Init() {
	flag.BoolVar(&force,
		"f",
		false,
		"-f: overwrite existing files")

	flag.BoolVar(&downloadHOT,
		"h",
		false,
		"-h download only HOT")
	flag.BoolVar(&downloadY,
		"y",
		false,
		"-y download only Y")
	flag.Parse()

	proxy := CONFIG.Proxy

	if proxy != "" {
		if proxyURL, err := url.Parse(proxy); err != nil {
			fmt.Println("proxy format must be 'scheme://[userinfo@]host/path[?query][#fragment]'")
			fmt.Println(err)
		} else {
			ProxySetting := http.ProxyURL(proxyURL)
			http.DefaultTransport = &http.Transport{
				Proxy: ProxySetting,
				/*
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).DialContext,
					MaxIdleConns:          100,
					IdleConnTimeout:       90 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				*/
			}
		}
	}
}

func main() {
	Init()
	overwrite := force
	if downloadY {
		GetY(filepath.Join(CONFIG.SaveDir, CONFIG.Y.Dirname), overwrite)
	}
	if downloadHOT {
		GetHOT(filepath.Join(CONFIG.SaveDir, CONFIG.HOT.Dirname), overwrite)
	}
	if !downloadHOT && !downloadY {
		GetHOT(filepath.Join(CONFIG.SaveDir, CONFIG.HOT.Dirname), overwrite)
		GetY(filepath.Join(CONFIG.SaveDir, CONFIG.Y.Dirname), overwrite)
	}
}
