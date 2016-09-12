// Author: ZHU HAIHUA
// Date: 9/12/16
package gap

import (
	pathpkg "path"
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
	"github.com/kimiazhu/vfs"
)

type assetsData map[string][]byte

type Filter func(path string, patterns []string) bool

func DefaultFilter(path string, ignoreList []string) bool {
	if ignoreList == nil {
		return false
	}

	p := pathpkg.Clean(path)
	for _, s := range ignoreList {
		return strings.Contains(pathpkg.Base(p), s)
	}

	return false
}

type Packager struct {
	FileFilter Filter
}

var DefaultPackager = &Packager{
	FileFilter: DefaultFilter,
}

// ReadAsset will read all data under root recursively into a assetData
func (p *Packager)ReadAsset(root string, ignoreError bool, ignoreList []string) (assetsData, error) {
	root = pathpkg.Clean(root)
	_, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	var data = make(assetsData)
	err = filepath.Walk(root, p.walkFunc(data, ignoreError, ignoreList))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Packager)Pack(root string, ignoreError bool, ignoreList []string) (vfs.FileSystem, error) {

}

func (p *Packager)walkFunc(data assetsData, ignoreError bool, ignoreList []string) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if !ignoreError && err != nil {
			return err
		}

		if !fi.IsDir() && !p.FileFilter(path, ignoreList) {
			_d, err := ioutil.ReadFile(path)
			if !ignoreError && err != nil {
				return err
			}
			data[path] = _d
		}

		return nil
	}
}
