package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ROOT_DIR        = "."
	BUILD_DIR       = filepath.Join(ROOT_DIR, "build")
	BUILD_FILE_PERM = os.ModePerm
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Expected command argument")
		printUsage()
		return
	}

	switch cmd := os.Args[1]; cmd {
	case "help":
		printUsage()

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
		copyTestScroll()
		invokeScroll("run", "test.scroll")

	case "log":
		setupBuild()
		goBuild()
		goFmt()
		goTest()
		copyTestScroll()
		invokeScroll("run", "-log", ".", "test.scroll")

	default:
		fmt.Println("[ERROR] Unknown command: " + cmd)
		printUsage()
	}
}

// *** Commands ***

func setupBuild() {
	removeDir(BUILD_DIR)

	if e := os.MkdirAll(BUILD_DIR, BUILD_FILE_PERM); e != nil {
		panik("Failed to create directory", e)
	}
}

func goBuild() {
	fmt.Println("Building...")

	const (
		BUILD_FLAGS = "" // "-gcflags -m -ldflags -s -w"
		MAIN_PKG    = "github.com/PaulioRandall/scarlet-go/scarlet"
	)

	// GO_PATH build -o OUTPUT_DIR BUILD_FLAGS MAIN_PKG
	cmd := newGoCmd(
		"build",
		"-o", BUILD_DIR,
		BUILD_FLAGS,
		MAIN_PKG,
	)

	if e := cmd.Run(); e != nil {
		panik("", e)
	}
}

func goFmt() {
	fmt.Println("Formatting...")

	cmd := newGoCmd("fmt", "./...")
	if e := cmd.Run(); e != nil {
		panik("", e)
	}
}

func goTest() {
	fmt.Println("Testing...")

	cmd := newGoCmd("test", "./...", "-timeout", "2s")
	if e := cmd.Run(); e != nil {
		panik("", e)
	}
}

func copyTestScroll() {

	src := filepath.Join(ROOT_DIR, "scarlet", "test.scroll")
	dst := filepath.Join(BUILD_DIR, "test.scroll")

	if e := copyFile(src, dst); e != nil {
		panik("Failed to copy test scroll", e)
	}
}

func invokeScroll(args ...string) {
	fmt.Println("Invoking scroll...")

	cd(BUILD_DIR)

	var e error
	exePath := filepath.Join(ROOT_DIR, "scarlet")
	exePath, e = filepath.Abs(exePath)
	if e != nil {
		panik("", e)
	}

	cmd := exec.Command(exePath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println(cmd.String())

	if e := cmd.Run(); e != nil {
		panik("", e)
	}

	cd("..")
}

// *** Script utils ***

func cd(dir string) {
	if e := os.Chdir(dir); e != nil {
		panik("Failed to change directory", e)
	}
}

func newGoCmd(args ...string) *exec.Cmd {

	goPath, e := exec.LookPath("go")
	if e != nil {
		panik("Can't find Go. Is it installed? Environment variables set?", e)
	}

	cmd := exec.Command(goPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func removeDir(dir string) {

	if _, e := os.Stat(dir); os.IsNotExist(e) {
		return
	} else if e != nil {
		panik("Failed to analyse directory", e)
	}

	if e := os.RemoveAll(dir); e != nil {
		panik("Failed to remove directory", e)
	}
}

func printUsage() {
	s := `Usage:
	help     Show usage
	clean    Remove build files
	build    Build -> format
	test     Build -> format -> test
	run      Build -> format -> test -> exe (test scroll)
	log      Build -> format -> test -> exe (test scroll + logs)`

	fmt.Println(s)
}

// *** General utils ***

func copyFile(src, dst string) error {

	stat, e := os.Stat(src)
	if e != nil {
		return fmt.Errorf("Missing file: %s", src)
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("Not a regular file: %s", src)
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
	} else if msg != "" {
		e = fmt.Errorf("%s: %s", msg, e.Error())
	}

	panic(e)
}
