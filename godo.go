package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	BUILD_DIR_NAME  = "build"
	OUTPUT_EXE_NAME = "scarlet"
	MAIN_GO_FILE    = "scarlet/scarlet.go"
	BUILD_FLAGS     = "" // -gcflags -m -ldflags "-s -w
	TEST_TIMEOUT    = "2s"
)

var (
	ROOT_DIR  string
	BUILD_DIR string
	OUTPUT    string
	MAIN_GO   string
)

func init() {

	ex, e := os.Executable()
	if e != nil {
		panic(e)
	}

	ROOT_DIR = filepath.Dir(ex)
	BUILD_DIR = filepath.Join(ROOT_DIR, BUILD_DIR_NAME)
	OUTPUT = filepath.Join(BUILD_DIR, OUTPUT_EXE_NAME)
	MAIN_GO = filepath.Join(ROOT_DIR, MAIN_GO_FILE)
}

func checkExists(f string) error {

	stat, e := os.Stat(f)
	if e != nil {
		return fmt.Errorf("Missing file: %s", f)
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("Not a regular file: %s", f)
	}

	return nil
}

func copyFile(src, dst string) error {

	e := checkExists(src)
	if e != nil {
		return e
	}

	srcFile, e := os.Open(src)
	if e != nil {
		return e
	}
	defer srcFile.Close()

	dstFile, e := os.Create(dst)
	if e != nil {
		return e
	}
	defer dstFile.Close()

	_, e = io.Copy(dstFile, srcFile)
	if e != nil {
		return e
	}

	return nil
}

func copyTestScroll() {

	src := filepath.Join(ROOT_DIR, "scarlet/test.scroll")
	dst := filepath.Join(BUILD_DIR, "test.scroll")

	e := copyFile(src, dst)
	if e != nil {
		e = fmt.Errorf("Failed to copy test scroll: %s", e.Error())
		panic(e)
	}
}
