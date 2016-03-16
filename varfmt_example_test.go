package varfmt_test

import (
	"fmt"

	"github.com/achiku/varfmt"
)

func ExamplePublicVarName() {
	malformattedVars := []string{
		"foo",
		"foo_bar",
		"fooBar",
		"foo_bar_buz",
		"foo_bar___buz",
		"foo_id",
		"foo_tls",
		"foo_json",
	}

	for _, s := range malformattedVars {
		fmt.Printf("%s -> %s\n", s, varfmt.PublicVarName(s))
	}
	// Output:
	// foo -> Foo
	// foo_bar -> FooBar
	// fooBar -> FooBar
	// foo_bar_buz -> FooBarBuz
	// foo_bar___buz -> FooBarBuz
	// foo_id -> FooID
	// foo_tls -> FooTLS
	// foo_json -> FooJSON

}
