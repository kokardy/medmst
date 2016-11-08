package main

import "regexp"

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

}

func main() {
	Init()
	GetY(SAVE_DIR_Y)
	GetHOT(SAVE_DIR_HOT)
}
