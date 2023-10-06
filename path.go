package filesystem

import "path"

// Path is a path to a file-system object.
type Path string

// IsAbsolute returns true if the path is absolute.
func (p Path) IsAbsolute() bool {
	return path.IsAbs((string)(p))
}

// IsRelative is oppasite to IsAbsolute, it returns
// true if the path is relative, and false if it's
// absolute.
func (p Path) IsRelative() bool {
	return !p.IsAbsolute()
}

// String returns this path as a string.
func (p Path) String() string {
	return (string)(p)
}
