package util

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

var lists = []string{
	"https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt",
}

func DownloadLists() {
	exists, err := Exists("./lists")
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		err := os.Mkdir("./lists", 0750)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, list := range lists {
		myUrl, err := url.Parse(list)
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(path.Clean("./lists/" + path.Base(myUrl.Path)))
		defer out.Close()

		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Get(list)
		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, resp.Body)
	}

	return
}
