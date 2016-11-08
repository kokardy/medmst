package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Extract(url string, reg *regexp.Regexp) (ext []string) {
	ext = make([]string, 0, 0)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	br := bufio.NewReader(res.Body)
	for {
		line, err := br.ReadString('\n')
		//fmt.Println(line)
		if err != nil {
			break
		}
		list := reg.FindAllString(line, 1)
		ext = append(ext, list...)
	}
	return
}

func SaveFile(url, dir, filename string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	filepath := filepath.Join(dir, filename)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, res.Body)
}

func GetURL(url1, url2 string) string {
	if strings.HasPrefix(url2, "http") {
		return url2
	} else if strings.HasPrefix(url2, "/") {
		reg := regexp.MustCompile(`https?://.*?/`)
		prefix := reg.FindString(url1)
		return prefix + url2[1:]
	} else {
		if !strings.HasSuffix(url1, "/") {
			url1 = url1 + "/"
		}
		return url1 + url2
	}

}

func Filename(url string) string {
	reg := regexp.MustCompile(`[^/]*$`)
	return reg.FindString(url)
}

func Download(html_url string, url_reg *regexp.Regexp, save_dir, savename string) {
	urls := Extract(html_url, url_reg)
	for i, pre_url := range urls {
		url := GetURL(html_url, pre_url)
		if savename == "" {
			format := "%s"
			savename = Filename(pre_url)
			SaveFile(url, save_dir, fmt.Sprintf(format, savename))
		} else {
			format := "%d_%s"
			SaveFile(url, save_dir, fmt.Sprintf(format, i, savename))
		}

	}
}

func GetY(save_dir string) {
	savename := "y.zip"
	Download(Y_PAGE_URL, Y_URL_REGEXP, save_dir, savename)
}

func GetHOT(save_dir string) {
	Download(HOT_PAGE_URL, HOT_URL_REGEXP, save_dir, "")
	Download(HOT_ADDPAGE_URL, HOT_ADDURL_REGEXP, save_dir, "")
	Download(HOT_DELPAGE_URL, HOT_DELURL_REGEXP, save_dir, "")
}
