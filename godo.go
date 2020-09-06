package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	BUILD_DIR_NAME  = "build"
	BUILD_FILE_PERM = 0777
	OUTPUT_EXE_NAME = "scarlet"
	MAIN_GO_FILE    = "scarlet/scarlet.go"
	BUILD_FLAGS     = "" // "-gcflags -m -ldflags -s -w"
	TEST_TIMEOUT    = "2s"
)

var (
	ROOT_DIR  string
	BUILD_DIR string
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Too few arguments")
		printUsage()
		return
	}

	ROOT_DIR = pwd()
	BUILD_DIR = filepath.Join(ROOT_DIR, BUILD_DIR_NAME)

	switch cmd := os.Args[1]; cmd {
	case "clean":
		removeDir(BUILD_DIR)

	case "build":
		setupBuild()
		goBuild()
		goFmt()

	case "test":
		setupBuild()
		goBuild()
		goFmt()
		goTest()

	case "run":
		setupBuild()
		goBuild()
		goFmt()
		goTest()
		runTestScroll()

	case "help":
		printUsage()

	default:
		fmt.Println("[ERROR] Unknown command: " + cmd)
		printUsage()
	}
}

// *** Commands ***

func setupBuild() {
	removeDir(BUILD_DIR)
	createDir(BUILD_DIR)
}

func goBuild() {
	// mainFile = filepath.Join(ROOT_DIR, MAIN_GO_FILE)
	// 	output := filepath.Join(BUILD_DIR, OUTPUT_EXE_NAME)
	// TODO
}

func goFmt() {
	// TODO
}

func goTest() {
	// TODO
}

func runTestScroll() {
	// TODO
}

// *** Script utils ***

func pwd() string {
	pwd, e := os.Getwd()
	if e != nil {
		panik("Failed to identify current directory", e)
	}
	return pwd
}

func removeDir(dir string) {
	if fileExists(dir) {
		if e := os.RemoveAll(dir); e != nil {
			panik("Failed to remove directory", e)
		}
	}
}

func createDir(dir string) {
	if e := os.MkdirAll(dir, BUILD_FILE_PERM); e != nil {
		panik("Failed to create directory", e)
	}
}

func copyTestScroll() {

	src := filepath.Join(ROOT_DIR, "scarlet/test.scroll")
	dst := filepath.Join(BUILD_DIR, "test.scroll")

	if e := copyFile(src, dst); e != nil {
		panik("Failed to copy test scroll", e)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\t./godo clean    \tDelete build directory")
	fmt.Println("\t./godo build    \tBuild -> format")
	fmt.Println("\t./godo test     \tBuild -> format -> test")
	fmt.Println("\t./godo run      \tBuild -> format -> test -> run test scroll")
	fmt.Println("\t./godo help     \tShow usage")
}

// *** General utils ***

func fileExists(f string) bool {
	_, e := os.Stat(f)
	return !os.IsNotExist(e)
}

func checkRegFile(f string) error {

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

	e := checkRegFile(src)
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

func panik(msg string, e error) {

	if e == nil {
		e = fmt.Errorf(msg)
	} else {
		e = fmt.Errorf("%s: %s", msg, e.Error())
	}

	panic(e)
}
