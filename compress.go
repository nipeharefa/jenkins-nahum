package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func compress() *bytes.Buffer {

	zipfile := new(bytes.Buffer)

	w := zip.NewWriter(zipfile)

	// var baseDir string
	err := filepath.Walk("node_modules", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		name, err := relative("", path)
		if err != nil {
			log.Error(err)
		}

		header.Name = name

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := w.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Error(err)
			return err
		}

		defer file.Close()
		_, err = io.Copy(writer, file)

		return err
	})

	if err != nil {
		log.Error(err)
		return nil
	}

	w.Close()

	return zipfile
}
