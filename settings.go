package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	Y_URL_REGEXP string

	HOT_PAGE_URL   string
	HOT_URL_REGEXP string

	HOT_ADDPAGE_URL   string
	HOT_ADDURL_REGEXP string

	HOT_DELPAGE_URL   string
	HOT_DELURL_REGEXP string

	SAVE_DIR_Y   string
	SAVE_DIR_HOT string
}

func (s Settings) Save(jsonfile string) error {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("An error occured in saving settings object.")
		return err
	}
	br := bytes.NewReader(b)
	bw, err := os.OpenFile(jsonfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("An error occured in saving settings object.")
		return err
	}
	defer bw.Close()
	io.Copy(bw, br)
	return nil
}

func LoadSettings(jsonfile string) (s Settings, err error) {
	br, err := os.OpenFile(jsonfile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("An error occured in loading settings object.")
		return
	}
	bw := bytes.NewBuffer([]byte(""))
	io.Copy(bw, br)
	b := bw.Bytes()

	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Println("An error occured in loading settings object.")
		return
	}

	return
}

func NewSettings() Settings {
	HOT_PAGE_URL = `http://www2.medis.or.jp/hcode/`
	return Settings{
		Y_PAGE_URL:   `http://www.iryohoken.go.jp/shinryohoshu/downloadMenu/`,
		Y_URL_REGEXP: `/shinryohoshu/downloadMenu/yFile;jsessionid=[0-9A-Z]+`,

		HOT_PAGE_URL:   HOT_PAGE_URL,
		HOT_URL_REGEXP: HOT_PAGE_URL + `moto_data/h[0-9]{8}.lzh`,

		HOT_ADDPAGE_URL:   HOT_PAGE_URL + `tuika/index.html`,
		HOT_ADDURL_REGEXP: HOT_PAGE_URL + `tuika/data/[0-9]{4}/[0-9]{8}.txt`,

		HOT_DELPAGE_URL:   HOT_PAGE_URL,
		HOT_DELURL_REGEXP: HOT_PAGE_URL + `moto_data/h[0-9]{8}del.txt`,

		SAVE_DIR_Y:   "save/y",
		SAVE_DIR_HOT: "save/hot",
	}
}
