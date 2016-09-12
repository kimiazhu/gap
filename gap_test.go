// Author: ZHU HAIHUA
// Date: 9/12/16
package gap_test

import (
	"testing"
	"github.com/kimiazhu/gap"
)

func TestReadAsset(t *testing.T) {
	p := gap.DefaultPackager
	//_pt := "/Users/zhuhaihua/Development/GoProject/src/github.com/kimiazhu/gap/testdata/"
	_pt := "testdada/"
	d, e := p.ReadAsset(_pt, true, nil)
	if e != nil {
		t.Errorf("read asset failed: %v", e)
	} else {
		t.Log(len(d))
		for k, v := range d {
			t.Logf("k is: %v, len of value is: %v", k, len(v))
		}
	}
}
