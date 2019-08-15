package values

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var v Values
	var b bytes.Buffer
	err := v.Render(&b)
	if err != nil {
		t.Error(err)
	}
	if b.Len() == 0 {
		t.Error(err)
	}
}
