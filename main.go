package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
)

var (
	FORCE bool
	HOT   bool
	Y     bool
)

func Init() {
	flag.StringVar(&SAVE_DIR, "d", "save",
		"-d save_dir: set save direactory")
	flag.StringVar(&PROXY, "p", "",
		"-p http://proxy_server:port: set proxy server")
	flag.BoolVar(&FORCE, "f", false,
		"-f: overwrite existing files")

	flag.BoolVar(&HOT,
		"h",
		true,
		"-h download only HOT")
	flag.BoolVar(&Y,
		"y",
		true,
		"-y download only Y")
	flag.Parse()

	var SETTINGS Settings
	if s, err := LoadSettings(SETTING_FILE); err != nil {
		SETTINGS = NewSettings()
		if err = SETTINGS.Save(SETTING_FILE); err != nil {
			panic(err)
		}
	} else {
		SETTINGS = s
	}
	if SAVE_DIR != "" {
		SETTINGS.SAVE_DIR = SAVE_DIR
	}
	if PROXY != "" {
		SETTINGS.PROXY = PROXY
	}

	Y_PAGE_URL = SETTINGS.Y_PAGE_URL
	Y_URL_REGEXP = regexp.MustCompile(SETTINGS.Y_URL_REGEXP)

	HOT_PAGE_URL = SETTINGS.HOT_PAGE_URL
	HOT_URL_REGEXP = regexp.MustCompile(SETTINGS.HOT_URL_REGEXP)

	HOT_ADDPAGE_URL = SETTINGS.HOT_ADDPAGE_URL
	HOT_ADDURL_REGEXP = regexp.MustCompile(SETTINGS.HOT_ADDURL_REGEXP)

	HOT_DELPAGE_URL = SETTINGS.HOT_DELPAGE_URL
	HOT_DELURL_REGEXP = regexp.MustCompile(SETTINGS.HOT_DELURL_REGEXP)

	SAVE_DIR = SETTINGS.SAVE_DIR
	SAVE_DIR_Y = SETTINGS.SAVE_DIR_Y
	SAVE_DIR_HOT = SETTINGS.SAVE_DIR_HOT

	PROXY = SETTINGS.PROXY

	if PROXY != "" {
		if proxyURL, err := url.Parse(PROXY); err != nil {
			fmt.Println("PROXY format must be 'scheme://[userinfo@]host/path[?query][#fragment]'")
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
	overwrite := FORCE
	if Y {
		GetY(filepath.Join(SAVE_DIR, SAVE_DIR_Y), overwrite)
	}
	if HOT {
		GetHOT(filepath.Join(SAVE_DIR, SAVE_DIR_HOT), overwrite)
	}
}
