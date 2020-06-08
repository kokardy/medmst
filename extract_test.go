package main

import "testing"

func TestExtract(t *testing.T) {
	extracted := Extract(CONFIG.Y.URL, CONFIG.Y.CompiledTarget())

	for _, url := range extracted {
		t.Log(url)
	}
}

func TestGetY(t *testing.T) {
	GetY(CONFIG.Y.Dirname, true)
}

func TestGetGenericMaster(t *testing.T) {
	GetGenericMaster(CONFIG.GenericMaster.Dirname, true)
}
