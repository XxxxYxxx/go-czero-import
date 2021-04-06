package c_superzk

import _ "github.com/sero-cash/go-czero-import/czero"

/*

#cgo CFLAGS: -I ../czero/szk_include

#cgo LDFLAGS: -L ../czero/lib_LINUX_AMD64_v4 -lsuperzk

*/
import "C"

func Is_czero_debug() bool {
	return false
}
