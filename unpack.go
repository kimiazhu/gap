// Author: ZHU HAIHUA
// Date: 9/12/16
package gap

import (
	"io/ioutil"
	"github.com/kimiazhu/vfs"
)

func load() (vfs.FileSystem, error) {
	return nil, nil
}

func (p* Packager)Read(path string) ([]byte, error) {
	// try to load from internal
	fs, err := load()
	if err != nil {
		return err
	}

	return ioutil.ReadAll(fs)
}
