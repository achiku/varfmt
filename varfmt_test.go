package varfmt_test

import (
	"testing"

	"github.com/achiku/varfmt"
)

func TestFormatFieldName(t *testing.T) {
	testData := map[string]string{
		"foo":           "Foo",
		"foo_bar":       "FooBar",
		"fooBar":        "FooBar",
		"foo_bar_buz":   "FooBarBuz",
		"foo_bar___buz": "FooBarBuz",
		"foo_id":        "FooID",
	}

	for k, v := range testData {
		f := varfmt.PublicVarName(k)
		if f != v {
			t.Errorf("expect %s got %s", v, f)
		}
	}
}
