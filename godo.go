package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"os/exec"
	"syscall"

	"github.com/pkg/errors"
	"io"
)

var (
	RootDir        = absPath(".")
	BuildDir       = filepath.Join(RootDir, "build")
	BuildPerm      = os.ModePerm
	BuildFlags     = ""   // "-gcflags -m -ldflags -s -w"
	TestTimeout    = "2s" // E.g. 1m, 5s, 250ms, etc
	MainPkgName    = "scarlet"
	MainPkg        = "github.com/PaulioRandall/scarlet-go/" + MainPkgName
	TestScrollName = "test.scroll"
	TestScrollSrc  = filepath.Join(RootDir, MainPkgName, TestScrollName)
	TestScroll     = filepath.Join(RootDir, "build", TestScrollName)
	Usage          = `Usage:
	help       Show usage
	clean      Remove build files
	build      Build -> format
	test       Build -> format -> test
	run        Build -> format -> test -> run`
)

func main() {

	code := 0
	args := os.Args[1:]

	if len(args) == 0 {
		usageErr("Missing command argument")
	}

	switch cmd := args[0]; strings.ToLower(cmd) {
	case "help":
		fmt.Println(Usage)

	case "clean":
		GoClean()

	case "build":
		GoClean()
		GoSetup()
		GoBuild()
		GoFormat()

	case "test":
		GoClean()
		GoSetup()
		GoBuild()
		GoFormat()
		GoTest()

	case "run":
		GoClean()
		GoSetup()
		GoBuild()
		GoFormat()
		GoTest()
		GoCopyTestScroll()
		code = GoRun("run", TestScroll)

	default:
		usageErr("Unknown command argument %q", cmd)
	}

	fmt.Printf("\nExit: %d\n", code)
	os.Exit(code)
}

func GoClean() {
	e := os.RemoveAll(BuildDir)
	exitIfErr(e, "Failed to remove build directory: %s", BuildDir)
}

func GoSetup() {
	e := os.MkdirAll(BuildDir, BuildPerm)
	exitIfErr(e, "Failed to make build directory: %s", BuildDir)
}

func GoBuild() {
	g, e := NewGo(RootDir)
	exitIfErr(e, "Failed to build")
	e = g.Build(BuildDir, BuildFlags, MainPkg)
	exitIfErr(e, "Failed to build")
}

func GoFormat() {
	g, e := NewGo(RootDir)
	exitIfErr(e, "Failed to format")
	e = g.FmtAll()
	exitIfErr(e, "Failed to format")
}

func GoTest() {
	g, e := NewGo(RootDir)
	exitIfErr(e, "Testing failed")
	e = g.TestAll(TestTimeout)
	exitIfErr(e, "Testing failed")
}

func GoCopyTestScroll() {
	e := CopyFile(TestScrollSrc, TestScroll, true)
	exitIfErr(e, "Failed to copy test scroll to build directory")
}

func GoRun(args ...string) int {
	var e error
	exePath := filepath.Join(BuildDir, MainPkgName)
	exePath, e = filepath.Abs(exePath)
	exitIfErr(e, "Failed to run")
	code, e := Run(exePath, BuildDir, args...)
	exitIfErr(e, "Failed to run")
	return code
}

func absPath(rel string) string {
	p, e := filepath.Abs(rel)
	exitIfErr(e, "Failed to identify path")
	return p
}

func exitIfErr(cause error, msg string, args ...interface{}) {
	if cause == nil {
		return
	}
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n", args...)
	fmt.Printf("Caused by: %+v", cause)
	os.Exit(code)
}

func usageErr(msg string, args ...interface{}) {
	const code = 1
	fmt.Printf("Exit: %d\n", code)
	fmt.Printf("Error: "+msg+"\n\n", args...)
	fmt.Println(Usage)
	os.Exit(code)
}

// Package "github.com/PaulioRandall/go-cookies/gobuild"

// Go represents a wrapper to the Go compiler. Functionality is provided for
// building, formatting, and testing.
type Go struct {
	Path    string
	WorkDir string
}

// NewGo returns a new Go struct. 'workDir' may be empty to signify the current
// working directory should be used.
func NewGo(workDir string) (Go, error) {

	var e error
	g := Go{WorkDir: workDir}

	if g.WorkDir == "" {
		if g.WorkDir, e = os.Getwd(); e != nil {
			return Go{}, Wrap(e,
				"Unable to identify current working directory")
		}
	}

	if g.Path, e = exec.LookPath("go"); e != nil {
		return Go{}, Wrap(e,
			"Can't find compiler. Is it installed? Environment variables set?")
	}
	return g, nil
}

func (g Go) NewCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(g.Path, args...)
	cmd.Dir = g.WorkDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (g Go) Build(dir, flags, pkg string) error {
	var cmd *exec.Cmd
	if dir == "" {
		cmd = g.NewCmd("build", flags, pkg)
	} else {
		cmd = g.NewCmd("build", "-o", dir, flags, pkg)
	}
	return run(cmd, "Build failed")
}

func (g Go) Fmt(pkg string) error {
	cmd := g.NewCmd("fmt", pkg)
	return run(cmd, "Format failed")
}

func (g Go) FmtAll() error {
	return g.Fmt("./...")
}

func (g Go) Test(pkg string, timeout string) error {
	var cmd *exec.Cmd
	if timeout == "" {
		cmd = g.NewCmd("test", pkg)
	} else {
		cmd = g.NewCmd("test", pkg, "-timeout", timeout)
	}
	return run(cmd, "Testing error")
}

func (g Go) TestAll(timeout string) error {
	return g.Test("./...", timeout)
}

func run(cmd *exec.Cmd, errMsg string) error {
	if e := cmd.Run(); e != nil {
		return Wrap(e, "Execution failed")
	}
	return nil
}

const (
	EXIT_OK  = 0 // Zero exit code
	EXIT_BAD = 1 // General error exit code
)

// Run runs the executable at 'exePath'. Setting the 'workDir' as empty will
// use the default as specified by functions that accept exec.Cmd. EXIT_OK is
// returned on successful execution otherwise EXIT_BAD or another non-zero
// exit code is returned.
func Run(exePath string, workDir string, args ...string) (int, error) {

	var e error

	cmd := exec.Command(exePath, args...)
	cmd.Dir = workDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if e := cmd.Start(); e != nil {
		return EXIT_BAD, e
	}

	if e = cmd.Wait(); e == nil {
		return EXIT_OK, nil
	}

	if exitErr, ok := e.(*exec.ExitError); ok {
		if stat, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			return stat.ExitStatus(), e
		}
	}

	return EXIT_BAD, e
}

// Package "github.com/PaulioRandall/go-cookies/cookies"

// Wrap wraps an error 'e' with a another message 'm'.
func Wrap(e error, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	return errors.Wrap(e, m)
}

// FileExists returns true if the file exists, false if not, and an error if
// file existence could not be determined.
func FileExists(f string) (bool, error) {
	_, e := os.Stat(f)
	if os.IsNotExist(e) {
		return false, nil
	}
	return true, e
}

// IsRegFile returns true if the file exists and is a regular file. An error is
// returned if this could not be determined.
func IsRegFile(f string) (bool, error) {
	stat, e := os.Stat(f)
	if os.IsNotExist(e) {
		return false, nil
	}
	if e != nil {
		return false, e
	}
	return stat.Mode().IsRegular(), nil
}

// SameFile returns true if the two files 'a' and 'b' describe the same file
// as determined by os.SameFile. An error is returned if the file info could
// not be retreived for either file.
func SameFile(a, b string) (bool, error) {
	aStat, e := os.Stat(a)
	if e != nil {
		return false, e
	}
	bStat, e := os.Stat(b)
	if e != nil {
		return false, e
	}
	return os.SameFile(aStat, bStat), nil
}

// CopyFile copies the single file 'src' to 'dst'.
func CopyFile(src, dst string, overwrite bool) error {

	if ok, e := IsRegFile(src); e != nil || !ok {
		return fmt.Errorf("Missing or not a regular file: %s", src)
	}

	if !overwrite {
		ok, e := FileExists(dst)
		if e != nil {
			return e
		}
		if ok {
			return fmt.Errorf("Destination already exists: %s", dst)
		}
	}

	same, e := SameFile(src, dst)
	if e == nil && same {
		return fmt.Errorf("Destination is the same as source: %s == %s", dst, src)
	}

	return NoCheckCopyFile(src, dst)
}

// NoCheckCopyFile copies the single file 'src' to 'dst' and doesn't make any
// attempt to check the file paths before hand.
func NoCheckCopyFile(src, dst string) error {

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
	return e
}
