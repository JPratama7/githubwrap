package githubwrap

import (
	"unsafe"
)

func unsafeGetBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
