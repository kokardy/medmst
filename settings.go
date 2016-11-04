package main

import (
	"regexp"
)

var (
	Y_PAGE_URL   = `http://www.iryohoken.go.jp/shinryohoshu/downloadMenu/`
	Y_URL_REGEXP = regexp.MustCompile(`/shinryohoshu/downloadMenu/yFile;jsessionid=[0-9A-Z]+`)

	HOT_PAGE_URL      = `http://www2.medis.or.jp/hcode/`
	HOT_ADDPAGE_URL   = HOT_PAGE_URL + `tuika/index.html`
	HOT_URL_REGEXP    = regexp.MustCompile(HOT_PAGE_URL + `moto_data/h[0-9]{8}.lzh`)
	HOT_DELURL_REGEXP = regexp.MustCompile(HOT_PAGE_URL + `moto_data/h[0-9]{8}del.txt`)
	HOT_ADDURL_REGEXP = regexp.MustCompile(HOT_PAGE_URL + `tuika/data/[0-9]{4}/[0-9]{8}.txt`)

	SAVE_DIR = "save"
)
