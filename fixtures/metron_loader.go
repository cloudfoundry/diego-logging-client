package fixtures

import (
	"os"
	"path/filepath"
	"runtime"
)

func fixturesRoot() string {
	_, file, _, _ := runtime.Caller(0)
	anchorDir := filepath.Dir(filepath.Dir(file))
	root := filepath.Join(anchorDir, "fixtures")
	if _, err := os.Stat(root); err != nil {
		panic("fixtures directory not found at " + root)
	}
	return root
}

// Path returns a path inside "fixtures/metron", defaulting the "metron" subdir.
func Path(name string) string {
	return filepath.Join(fixturesRoot(), "metron", name)
}
