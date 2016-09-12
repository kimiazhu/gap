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
	"bytes"
	"io"
	"compress/gzip"
	"fmt"
	"github.com/golang/protobuf/proto"
)

const (
	LINE_BREAKER = "\n"
)

type data []byte
type assetsData map[string]data

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

type Packer struct {
	FileFilter Filter
}

var DefaultPackager = &Packer{
	FileFilter: DefaultFilter,
}

// ReadAsset will read all data under root recursively into a assetData
func (p *Packer)ReadAsset(root string, ignoreError bool, ignoreList []string) (assetsData, error) {
	root = pathpkg.Clean(root)
	_, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	var dat = make(assetsData)
	err = filepath.Walk(root, p.walkFunc(dat, ignoreError, ignoreList))
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func (p *Packer)Pack(root string, ignoreError bool, ignoreList []string) (vfs.FileSystem, error) {

}

func (p *Packer)walkFunc(dat assetsData, ignoreError bool, ignoreList []string) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if !ignoreError && err != nil {
			return err
		}

		if !fi.IsDir() && !p.FileFilter(path, ignoreList) {
			_d, err := ioutil.ReadFile(path)
			if !ignoreError && err != nil {
				return err
			}
			dat[path] = _d
		}

		return nil
	}
}

// Marshal return text representation of assetsData
func (d *assetsData) Marshal() string {
	proto.MarshalTextString()
}

func (d *data)compress(level int) []byte {
	if level <= 0 {
		return d
	} else {
		compressed := new(bytes.Buffer)
		gz, _ := gzip.NewWriterLevel(compressed, level)
		defer gz.Close()
		io.Copy(gz, bytes.NewBuffer(d))

		return compressed.Bytes()
	}
}

func (d* data)compressToString(data []byte, level int) string {
	compressed := d.compress(level)

	str := ""
	i := 0
	for _, v := range compressed {
		if i%12 == 0 {
			str = str + LINE_BREAKER
			i = 0
		}

		str = str + fmt.Sprintf("0x%02x,", v)
		i++
	}

	return str
}