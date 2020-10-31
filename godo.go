package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/go-cookies/cookies"
	"github.com/PaulioRandall/go-cookies/go/quick"
)

var (
	ROOT            = quick.AbsPath(".")
	BUILD           = filepath.Join(ROOT, "build")
	PROJ_PATH       = "github.com/PaulioRandall/scarlet-go"
	MAIN_PKG        = "cmd"
	TEST_SCROLL     = "test.scroll"
	TEST_SCROLL_SRC = filepath.Join(ROOT, MAIN_PKG, TEST_SCROLL)
	TEST_SCROLL_DST = filepath.Join(ROOT, "build", TEST_SCROLL)
	USAGE           = `Usage:
	help       Show usage
	clean      Remove build files
	build      Build -> format -> vet
	test       Build -> format -> test -> vet
	run        Build -> format -> test -> vet -> run`
)

var (
	BUILD_ARGS = []string{
		"-o", BUILD,
		"", // "-gcflags -m -ldflags -s -w"
		PROJ_PATH + "/" + MAIN_PKG,
	}
	FMT_ARGS  = []string{"./..."}
	TEST_ARGS = []string{"-timeout", "2s", "./..."}
	VET_ARGS  = []string{"./..."}
	RUN_ARGS  = []string{"run", TEST_SCROLL}
)

func main() {

	code := 0
	args := os.Args[1:]

	if len(args) == 0 {
		quick.UsageErr(USAGE, "Missing command argument")
	}

	switch cmd := args[0]; strings.ToLower(cmd) {
	case "help":
		fmt.Println(USAGE)

	case "clean":
		quick.Clean(BUILD)

	case "build":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)

	case "test":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Test(ROOT, TEST_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)

	case "run":
		quick.Clean(BUILD)
		quick.Setup(BUILD, os.ModePerm)
		quick.Build(ROOT, BUILD_ARGS...)
		quick.Fmt(ROOT, FMT_ARGS...)
		quick.Test(ROOT, TEST_ARGS...)
		quick.Vet(ROOT, VET_ARGS...)

		e := cookies.CopyFile(TEST_SCROLL_SRC, TEST_SCROLL_DST, true)
		quick.ExitIfErr(e, "Failed to copy test scroll to build directory")
		code = quick.Run(BUILD, MAIN_PKG, RUN_ARGS...)

	default:
		quick.UsageErr(USAGE, "Unknown command argument %q", cmd)
	}

	if code != 0 {
		fmt.Printf("\nExit: %d\n", code)
	}
	os.Exit(code)
}
