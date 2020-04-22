package main

import (
	"crypto/md5" // #nosec
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// readerHasher generic md5 hash generater from io.Reader.
func readerHasher(readers ...io.Reader) (string, error) {
	// Use go1.14 new hashmap functions.
	h := md5.New() // #nosec

	for _, r := range readers {
		if _, err := io.Copy(h, r); err != nil {
			return "", fmt.Errorf("write reader as hash, %w", err)
		}
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func checksumFunc(p string) string {
	path, err := filepath.Abs(filepath.Clean(p))
	if err != nil {
		return ""
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
		// level.Error(logger).Log("cache key template/checksum could not open file")
		return ""
	}

	str, err := readerHasher(f)
	if err != nil {
		panic(err)
		// level.Error(logger).Log("cache key template/checksum could not generate hash")
		return ""
	}

	return str
}

func relative(parent string, path string) (string, error) {
	name := filepath.Base(path)

	rel, err := filepath.Rel(parent, filepath.Dir(path))
	if err != nil {
		return "", fmt.Errorf("relative path <%s>, base <%s>, %w", rel, name, err)
	}

	// NOTICE: filepath.Rel puts "../" when given path is not under parent.
	for strings.HasPrefix(rel, "../") {
		rel = strings.TrimPrefix(rel, "../")
	}

	rel = filepath.ToSlash(rel)

	return strings.TrimPrefix(filepath.Join(rel, name), "/"), nil
}
