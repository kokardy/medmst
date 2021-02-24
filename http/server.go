package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func pmdaServer() (result string) {
	result = os.Getenv("PMDA")
	if result == "" {
		result = "www.pmda.go.jp"
	}
	return
}

func pmdaPath(yjcode string) (path string, err error) {
	f := fmt.Sprintf("/asset/bin/yj/%s", yjcode)
	if _, err = os.Stat(f); err != nil {
		return
	}
	r, err := os.OpenFile(f, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, r)
	if err != nil {
		return
	}
	line := buf.String()
	path = strings.Trim(line, "\r\n ")
	return
}

func redirectToPMDA(c *gin.Context) {
	yjcode := c.Param("yjcode")
	path, err := pmdaPath(yjcode)
	if err != nil {
		message := fmt.Sprintf("NOT FOUND: %s \n error:%s", yjcode, err)
		fmt.Println(message)
		c.String(http.StatusNotFound, message)
		return
	}
	var pmda = pmdaServer()
	var url = fmt.Sprintf("//%s/PmdaSearch/iyakuDetail/%s", pmda, path)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func guessPmdaPath(yjcode string) (path_list []string, err error) {
	pattern := fmt.Sprintf("/asset/bin/html/??????_%s*_?_??.html", yjcode)
	path_list, err = filepath.Glob(pattern)
	return
}

//TODO 実装中
func guessRedirectToPMDA(c *gin.Context) {
	yjcode := c.Param("yjcode")
	path_list, err := guessPmdaPath(yjcode)
	fmt.Printf("path_list: %s\n", path_list)
	if err != nil {
		message := fmt.Sprintf(
			"pattern search error: yj:%s \nerror:%s\n", yjcode, err)
		fmt.Println(message)
		c.String(http.StatusNotFound, message)
		return
	}

	var url = ""
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Converter(YJcode to PMDA path) works.")
	})

	//redirect to PMDA
	r.GET("/redirect/:yjcode/", redirectToPMDA)
	r.GET("/redirect_guess/:yjcode/", guessRedirectToPMDA)

	fmt.Println("listen: 80")
	r.Run(":80")
}
