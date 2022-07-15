package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"excelreader/metadata"
)

func main() {
	url := "https://download.microsoft.com/download/1/4/E/14EDED28-6C58-4055-A65C-23B4DA81C4DE/Financial%20Sample.xlsx"
	res, err := http.Get(url)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}
	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"Core Props: %s %s - App Props: %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}
