package concurrency

import (
	"archive/zip"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//download downloads url and returns the contents and error

func Download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)

	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return filename, err

}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

func WriteZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename) // 파일 생성
	if err != nil {
		return err
	}

	zw := zip.NewWriter(outf)
	for _, filename := range filenames {

		w, err := zw.Create(filename) // 파일 생성

		if err != nil {
			return err
		}

		f, err := os.Open(filename) // 파일 오픈

		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(w, f) //카피

		if err != nil {
			return err
		}

	}
	return zw.Close()
}
