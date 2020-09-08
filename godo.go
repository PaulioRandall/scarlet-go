package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	ROOT_DIR        = "."
	BUILD_DIR       = filepath.Join(ROOT_DIR, "build")
	BUILD_FILE_PERM = os.ModePerm
	COMMANDS        = map[string]string{
		"help":  "Show usage",
		"clean": "Remove build files",
		"build": "Build -> format",
		"test":  "Build -> format -> test",
		"run":   "Build -> format -> test -> run (with test scroll)",
	}
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Too few arguments")
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
		runTestScroll()

	default:
		fmt.Println("[ERROR] Unknown command: " + cmd)
		printUsage()
	}
}

// *** Commands ***

func setupBuild() {
	removeDir(BUILD_DIR)
	createDir(BUILD_DIR, BUILD_FILE_PERM)
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

func runTestScroll() {
	fmt.Println("Test scroll...")

	cd(BUILD_DIR)

	var e error
	exePath := filepath.Join(ROOT_DIR, "scarlet")
	exePath, e = filepath.Abs(exePath)
	if e != nil {
		panik("", e)
	}

	cmd := exec.Command(exePath, "run", "test.scroll")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
	if fileExists(dir) {
		if e := os.RemoveAll(dir); e != nil {
			panik("Failed to remove directory", e)
		}
	}
}

func createDir(dir string, mode os.FileMode) {
	if e := os.MkdirAll(dir, mode); e != nil {
		panik("Failed to create directory", e)
	}
}

func printUsage() {

	cmdWidth := findCmdNameWidth()

	fmt.Println("Usage:")
	for cmd, desc := range COMMANDS {
		fmt.Print("\t")
		cmd = pad(cmdWidth, string(cmd))
		fmt.Print(cmd)
		fmt.Println(desc)
	}
}

func findCmdNameWidth() int {

	const PADDING = 4
	w := 0

	for cmd := range COMMANDS {
		n := len(cmd)
		if n > w {
			w = n
		}
	}

	return w + PADDING
}

func pad(fixedWidth int, s string) string {

	n := fixedWidth - len(s)
	if n == 0 {
		return s
	}

	if n < 0 {
		msg := fmt.Sprintf(
			"Length of '%s' exceeds padding fixed width of %d", s, fixedWidth,
		)
		panik(msg, nil)
	}

	return s + strings.Repeat(" ", n)
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
	} else if msg != "" {
		e = fmt.Errorf("%s: %s", msg, e.Error())
	}

	panic(e)
}
