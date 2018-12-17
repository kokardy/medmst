package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/BurntSushi/toml"
)

var (
	HOTURL     = `http://www2.medis.or.jp/hcode/` //HOTURL
	ConfigFile = `config.toml`
	CONFIG     = LoadConfig(ConfigFile)
)

//Config binds config.toml
type Config struct {
	Proxy   string
	SaveDir string
	Y       *Y
	HOT     *HOT
	HOTAdd  *HOTAdd
	HOTDel  *HOTDel
}

//NewConfig create a Config struct.
func NewConfig() Config {
	return Config{
		Proxy:   "",
		SaveDir: "save",
		Y:       NewY(),
		HOT:     NewHOT(),
		HOTAdd:  NewHOTAdd(),
		HOTDel:  NewHOTDel(),
	}
}

//Init initializes Config.
func (config *Config) Init() {
	sites := []Initer{
		config.Y,
		config.HOT,
		config.HOTAdd,
		config.HOTDel,
	}
	for _, s := range sites {
		s.Init()
	}
}

//Dump create config.toml
func (config Config) Dump(path string) error {
	w, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("file open error: %s", path)
		return err
	}
	enc := toml.NewEncoder(w)
	err = enc.Encode(&config)
	return err
}

//LoadConfig create Config from config.toml.
func LoadConfig(path string) Config {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Cannot read file: %s", path)
		config := NewConfig()
		err = config.Dump(path)
		if err != nil {
			fmt.Printf("Failed create config.toml: %s", path)
		}
		return config
	}
	return readConfig(path)
}

func readConfig(path string) Config {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Printf("TOML decode error: %s", path)
		fmt.Println(err)
	}
	return config
}

//Initer has Init function
type Initer interface {
	Init()
}

//Site has setting to download files.
type Site struct {
	URL          string
	TargetRegexp string
	Dirname      string
}

//CompiledTarget return *regexp.Regexp compiling TargetRegexp.
func (s Site) CompiledTarget() *regexp.Regexp {
	compiledRegexp, err := regexp.Compile(s.TargetRegexp)
	if err != nil {
		fmt.Printf("Failed regexp compile: %s", s.TargetRegexp)
		panic("regexp error")
	}
	return compiledRegexp
}

//Y site
type Y struct {
	Site
	Initer
}

//NewY create a new Y
func NewY() *Y {
	y := &Y{}
	y.Init()
	return y
}

//Init initialize Y
func (y *Y) Init() {
	if y.URL == "" {
		y.URL = `http=//www2.medis.or.jp/hcode/`
		y.TargetRegexp = `/shinryohoshu/downloadMenu/yFile;jsessionid=[0-9A-Z]+`
		y.Dirname = `y`
	}
}

//HOT site
type HOT struct {
	Site
	Initer
}

//NewHOT create a new HOT
func NewHOT() *HOT {
	h := &HOT{}
	h.Init()
	return h
}

//Init initialize HOT
func (h *HOT) Init() {
	if h.URL == "" {
		h.URL = HOTURL
		h.TargetRegexp = HOTURL + `moto_data/h[0-9]{8}.lzh`
		h.Dirname = `hot`
	}
}

//HOTAdd site
type HOTAdd struct {
	Site
	Initer
}

//NewHOTAdd create a new HOTAdd
func NewHOTAdd() *HOTAdd {
	ha := &HOTAdd{}
	ha.Init()
	return ha
}

//Init initialize HOTAdd
func (ha *HOTAdd) Init() {
	if ha.URL == "" {
		ha.URL = HOTURL + `tuika/index.html`
		ha.TargetRegexp = HOTURL + `tuika/data/[0-9]{4}/[0-9]{8}.txt`
		ha.Dirname = `hot`
	}
}

//HOTDel site
type HOTDel struct {
	Site
	Initer
}

//NewHOTDel create a new HOTDel
func NewHOTDel() *HOTDel {
	hd := &HOTDel{}
	hd.Init()
	return hd
}

//Init initialize HOTDel
func (hd *HOTDel) Init() {
	if hd.URL == "" {
		hd.URL = HOTURL
		hd.TargetRegexp = HOTURL + `moto_data/h[0-9]{8}.lzh`
		hd.Dirname = `hot`
	}
}