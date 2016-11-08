package main

import (
	"regexp"
)

var SETTING_FILE = `settings.json`
var (
	Y_PAGE_URL   string
	Y_URL_REGEXP *regexp.Regexp

	HOT_PAGE_URL   string
	HOT_URL_REGEXP *regexp.Regexp

	HOT_ADDPAGE_URL   string
	HOT_ADDURL_REGEXP *regexp.Regexp

	HOT_DELPAGE_URL   string
	HOT_DELURL_REGEXP *regexp.Regexp

	SAVE_DIR_Y   string
	SAVE_DIR_HOT string
)

type Settings struct {
	Y_PAGE_URL   string
	Y_URL_REGEXP *regexp.Regexp

	HOT_PAGE_URL   string
	HOT_URL_REGEXP *regexp.Regexp

	HOT_ADDPAGE_URL   string
	HOT_ADDURL_REGEXP *regexp.Regexp

	HOT_DELPAGE_URL   string
	HOT_DELURL_REGEXP *regexp.Regexp

	SAVE_DIR_Y   string
	SAVE_DIR_HOT string
}

func (s *Settings) save() {

}

func (s *Settings) load() {

}

func NewSettings() Settings {

	return Settings{
		Y_PAGE_URL:   `http://www.iryohoken.go.jp/shinryohoshu/downloadMenu/`,
		Y_URL_REGEXP: regexp.MustCompile(`/shinryohoshu/downloadMenu/yFile;jsessionid=[0-9A-Z]+`),

		HOT_PAGE_URL:   `http://www2.medis.or.jp/hcode/`,
		HOT_URL_REGEXP: regexp.MustCompile(HOT_PAGE_URL + `moto_data/h[0-9]{8}.lzh`),

		HOT_ADDPAGE_URL:   HOT_PAGE_URL + `tuika/index.html`,
		HOT_ADDURL_REGEXP: regexp.MustCompile(HOT_PAGE_URL + `tuika/data/[0-9]{4}/[0-9]{8}.txt`),

		HOT_DELPAGE_URL:   HOT_PAGE_URL,
		HOT_DELURL_REGEXP: regexp.MustCompile(HOT_PAGE_URL + `moto_data/h[0-9]{8}del.txt`),

		SAVE_DIR_Y:   "save/y",
		SAVE_DIR_HOT: "save/hot",
	}
}
