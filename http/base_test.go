package http

import "testing"

func TestWithoutHandler(t *testing.T) {
	withoutHandler()
}

func TestWithHandler(t *testing.T) {
	WithHandler()
}
