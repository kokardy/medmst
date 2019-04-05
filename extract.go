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

//Extract extract URLs in URL with RegExp.
func Extract(url string, reg *regexp.Regexp) (ext []string) {
	ext = make([]string, 0, 0)
	fmt.Printf("url %s\n", url)
	fmt.Printf("reg %s\n", reg)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error occured in opening http response.")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	br := bufio.NewReader(res.Body)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		list := reg.FindAllString(line, 1)
		if Len(list) > 0 {
			ext = append(ext, list...)
			fmt.Printf("append: %s\n", list)
		}
	}
	return
}

//SaveFile store a file from URL.
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

//getURL create new URL from two URLs.
func getURL(url1, url2 string) string {
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

//filename get filename from URL's postfix.
func filename(url string) string {
	reg := regexp.MustCompile(`[^/]*$`)
	return reg.FindString(url)
}

//Download download files using Extract function.
func Download(htmlURL string, urlReg *regexp.Regexp, saveDir, savename string, all bool, overwrite bool) {
	urls := Extract(htmlURL, urlReg)
	fmt.Println(urls)
	for i, preURL := range urls {
		url := getURL(htmlURL, preURL)
		if !all && i > 0 {
			break
		}
		if savename != "" && all {
			format := "%d_%s"
			SaveFile(url, saveDir, fmt.Sprintf(format, i, savename), overwrite)
		} else if savename == "" {
			format := "%s"
			savename = filename(preURL)
			SaveFile(url, saveDir, fmt.Sprintf(format, savename), overwrite)
		} else {
			SaveFile(url, saveDir, savename, overwrite)
		}
	}
}

//GetY download Y file.
func GetY(saveDir string, overwrite bool) {
	savename := "y.zip"
	all := false
	Download(
		CONFIG.Y.URL,
		CONFIG.Y.CompiledTarget(),
		saveDir,
		savename,
		all,
		overwrite,
	)
}

//GetHOT download HOT, HOTAdd and HOTDel files.
func GetHOT(saveDir string, overwrite bool) {
	all := false
	defaultName := ""
	Download(
		CONFIG.HOT.URL,
		CONFIG.HOT.CompiledTarget(),
		saveDir,
		defaultName,
		all,
		overwrite,
	)
	Download(
		CONFIG.HOTAdd.URL,
		CONFIG.HOTAdd.CompiledTarget(),
		saveDir,
		defaultName,
		all,
		overwrite,
	)
	Download(
		CONFIG.HOTDel.URL,
		CONFIG.HOTDel.CompiledTarget(),
		saveDir,
		defaultName,
		all,
		overwrite,
	)
}
