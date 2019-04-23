package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
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
	f := fmt.Sprintf("/asset/bin/data/yj/%s", yjcode)
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
		c.String(http.StatusNotFound, message)
		return
	}
	var pmda = pmdaServer()
	var url = fmt.Sprintf("//%s/PmdaSearch/iyakuSearch/%s", pmda, path)

	c.Redirect(http.StatusMovedPermanently, url)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Converter(YJcode to PMDA path) works.")
	})

	//redirect to PMDA
	r.GET("/redirect/:yjcode/", redirectToPMDA)

	fmt.Println("listen: 80")
	r.Run(":80")
}
