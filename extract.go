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

func SaveFile(url, dir, filename string, overwrite bool) {
	if !overwrite {
		_f := filepath.Join(dir, filename)
		if _, err := os.Stat(_f); err == nil {
			fmt.Printf("%s/%s already exists.\n", dir, filename)
			return
		}
	}
	fmt.Printf("Download from %s --> %s/%s \r", url, dir, filename)
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
	//f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, res.Body)
	fmt.Printf("Download from %s --> %s/%s OK!\n", url, dir, filename)
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

func Download(html_url string, url_reg *regexp.Regexp, save_dir, savename string, all bool, overwrite bool) {
	urls := Extract(html_url, url_reg)
	for i, pre_url := range urls {
		url := GetURL(html_url, pre_url)
		if !all && i > 0 {
			break
		}
		if savename != "" && all {
			format := "%d_%s"
			SaveFile(url, save_dir, fmt.Sprintf(format, i, savename), overwrite)
		} else if savename == "" {
			format := "%s"
			savename = Filename(pre_url)
			SaveFile(url, save_dir, fmt.Sprintf(format, savename), overwrite)
		} else {
			SaveFile(url, save_dir, savename, overwrite)
		}
	}
}

func GetY(save_dir string, overwrite bool) {
	savename := "y.zip"
	all := false
	Download(Y_PAGE_URL, Y_URL_REGEXP, save_dir, savename, all, overwrite)
}

func GetHOT(save_dir string, overwrite bool) {
	all := false
	default_name := ""
	Download(HOT_PAGE_URL, HOT_URL_REGEXP, save_dir, default_name, all, overwrite)
	Download(HOT_ADDPAGE_URL, HOT_ADDURL_REGEXP, save_dir, default_name, all, overwrite)
	Download(HOT_DELPAGE_URL, HOT_DELURL_REGEXP, save_dir, default_name, all, overwrite)
}
