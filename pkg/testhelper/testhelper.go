// Package testhelper is helpful package for test routine
package testhelper

import (
	"fmt"
	"os"
	"path/filepath"
)

// RemoveContentFromRoot provides removing folder from the root
func RemoveContentFromRoot(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return os.Remove(dir)
}

// MakeRootDir create root dir in the case if this is not exist
func MakeRootDir(rootDir string) error {
	_, err := os.Stat(rootDir)
	if err == nil {
		return nil
	}
	if !os.IsExist(err) {
		if err := os.Mkdir(rootDir, os.ModePerm); err != nil {
			return fmt.Errorf("unable to create root dir: %v", err)
		}
	}
	return nil
}
