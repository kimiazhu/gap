// Author: ZHU HAIHUA
// Date: 9/12/16
package gap

import (
	"io/ioutil"
	"github.com/kimiazhu/vfs"
	"github.com/kimiazhu/vfs/mapbytefs"
)

func load() (map[string][]byte, error) {
	return mapbytefs.New(nil), nil
}

func (p* Packer)Read(path string) ([]byte, error) {
	// try to load from internal
	ns := vfs.NameSpace{}
	fs, err := load()
	if err == nil {
		// there is an internal filesystem, mount it firstly.
		ns.Bind("/", fs, "/", vfs.BindReplace)
	}
	// mount the os filesystem
	ns.Bind("/", vfs.OS("."), "/", vfs.BindAfter)

	f, err := ns.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
