package path

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// File wraps an os.FileInfo as well as the absolute path to the underlying file.
type File struct {
	FullPath string
	Info     os.FileInfo
}

// Standards lists all standard files.
var Standards = func() ([]File, error) {
	return filesFor("standards", "yml")
}

// Narratives lists all narrative files.
var Narratives = func() ([]File, error) {
	return filesFor("narratives", "md")
}

// Policies lists all policy files.
var Policies = func() ([]File, error) {
	return filesFor("policies", "md")
}

// Procedures lists all procedure files.
var Procedures = func() ([]File, error) {
	return filesFor("procedures", "md")
}

func filesFor(name, extension string) ([]File, error) {
	var filtered []File
	entries, err := os.ReadDir(filepath.Join(".", name))
	if err != nil {
		return nil, errors.Wrap(err, "unable to load files for: "+name)
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), "."+extension) || strings.HasPrefix(strings.ToUpper(entry.Name()), "README") {
			continue
		}
		abs, err := filepath.Abs(filepath.Join(".", name, entry.Name()))
		if err != nil {
			return nil, errors.Wrap(err, "unable to load file: "+entry.Name())
		}
		info, err := entry.Info()
		if err != nil {
			return nil, errors.Wrap(err, "unable to stat file: "+entry.Name())
		}
		filtered = append(filtered, File{abs, info})
	}
	return filtered, nil
}
