package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

func Init() {

	var SETTINGS Settings
	if s, err := LoadSettings(SETTING_FILE); err != nil {
		SETTINGS = NewSettings()
		if err = SETTINGS.Save(SETTING_FILE); err != nil {
			panic(err)
		}
	} else {
		SETTINGS = s
	}

	Y_PAGE_URL = SETTINGS.Y_PAGE_URL
	Y_URL_REGEXP = regexp.MustCompile(SETTINGS.Y_URL_REGEXP)

	HOT_PAGE_URL = SETTINGS.HOT_PAGE_URL
	HOT_URL_REGEXP = regexp.MustCompile(SETTINGS.HOT_URL_REGEXP)

	HOT_ADDPAGE_URL = SETTINGS.HOT_ADDPAGE_URL
	HOT_ADDURL_REGEXP = regexp.MustCompile(SETTINGS.HOT_ADDURL_REGEXP)

	HOT_DELPAGE_URL = SETTINGS.HOT_DELPAGE_URL
	HOT_DELURL_REGEXP = regexp.MustCompile(SETTINGS.HOT_DELURL_REGEXP)

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
	GetY(SAVE_DIR_Y)
	GetHOT(SAVE_DIR_HOT)
}
