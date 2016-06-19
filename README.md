# varfmt

[![Build Status](https://travis-ci.org/achiku/varfmt.svg?branch=master)](https://travis-ci.org/achiku/varfmt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/varfmt/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/achiku/varfmt)](https://goreportcard.com/report/github.com/achiku/varfmt)

Convert string into Go variable name following Go naming convention

## Why created

Although there are a lot of needs to generate go struct from various sources (DDL, yml, toml, JSON, etc), I couldn't find any common library to format struct field names that follow go naming convention. Plus, I found really nice little piece of code that formats variable name in [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson).


## How to use

```go
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
```

## Credits

Code is almost entirely borrowed from https://github.com/ChimeraCoder/gojson. Huge thanks to [ChimeraCoder](https://github.com/ChimeraCoder) !!
