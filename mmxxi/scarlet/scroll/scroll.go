package scroll

import (
	"path/filepath"
	"strings"
)

// Scroll represents a Scarlet scroll or source file.
type Scroll struct {
	Name string // Name of the file without extension
	Dir  string // Directory containing the file
	Path string // Path to the file including name and extension
}

// Make returns a new scroll from a file path.
func Make(path string) (s Scroll, e error) {

	s.Path, e = filepath.Abs(path)
	if e != nil {
		return s, e
	}

	s.Name = filepath.Base(path)
	s.Name = strings.TrimSuffix(s.Name, filepath.Ext(s.Name))
	s.Dir = filepath.Dir(path)

	return s, nil
}
