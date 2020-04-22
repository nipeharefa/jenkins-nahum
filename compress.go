package main

import (
	"compress/flate"

	"github.com/mholt/archiver/v3"
	log "github.com/sirupsen/logrus"
)

func cc(filename string) {
	z := archiver.Zip{
		CompressionLevel:     flate.DefaultCompression,
		MkdirAll:             true,
		SelectiveCompression: true,
		ContinueOnError:      false,
		OverwriteExisting:    false,
	}

	err := z.Archive([]string{"node_modules", ".next/cache"}, filename)
	if err != nil {
		log.Error(err)
	}
}
