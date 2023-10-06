package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath_IsAbsolute(t *testing.T) {
	absolutes := []Path{
		"/",
		"/smth",
		"/smth/",
	}
	for _, absolute := range absolutes {
		t.Run(absolute.String(), func(t *testing.T) {
			assert.Truef(t, absolute.IsAbsolute(), `"%s" was expected to be absolute`, absolute)
			assert.Falsef(t, absolute.IsRelative(), `"%s" is both, relative and absolute`, absolute)
		})
	}
}

func TestPath_IsRelative(t *testing.T) {
	relatives := []Path{
		"smth",
		"./smth",
		"../smth/",
	}
	for _, relative := range relatives {
		t.Run(relative.String(), func(t *testing.T) {
			assert.Truef(t, relative.IsRelative(), `"%s" was expected to be relative`, relative)
			assert.Falsef(t, relative.IsAbsolute(), `"%s" is both, relative and absolute`, relative)
		})
	}
}

func TestPath_String(t *testing.T) {
	assert.Equal(t, (Path)("/").String(), "/")
}
