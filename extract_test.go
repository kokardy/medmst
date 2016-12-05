package main

import "testing"

func TestExtract(t *testing.T) {
	extracted := Extract(Y_PAGE_URL, Y_URL_REGEXP)

	for _, url := range extracted {
		t.Log(url)
	}
}

func TestGetY(t *testing.T) {
	GetY(SAVE_DIR_Y)
}
