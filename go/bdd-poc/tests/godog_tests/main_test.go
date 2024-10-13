package godog_tests

import (
	"bytes"

	"github.com/cucumber/godog"
)

type CtxKey struct{}

var buf bytes.Buffer

var Options = godog.Options{
	Format:   "pretty",
	Paths:    []string{"../../features"},
	NoColors: true,
	Output:   &buf,
}
