package testing

import (
	"os"
	"path"
	"runtime"
)

// A workaround for running all tests on the project root dir.

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}


