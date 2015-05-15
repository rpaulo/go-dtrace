package dtrace

// #cgo CFLAGS: -I/usr/src/cddl/contrib/opensolaris/lib/libdtrace/common
// #cgo CFLAGS: -I/usr/src/sys/cddl/compat/opensolaris
// #cgo CFLAGS: -I/usr/src/sys/cddl/contrib/opensolaris/uts/common
// #cgo CFLAGS: -Wno-unknown-attributes
// #cgo LDFLAGS: -ldtrace -lelf -lctf -lproc -lrtld_db -lutil
// #include <dtrace.h>
import "C"

import (
	"unsafe"
)

const DTraceVersion = 3
type DTraceHandle *C.struct_dtrace_hdl

func Open(version int, flags int, error *int) DTraceHandle {
	return C.dtrace_open(C.int(version), C.int(flags),
		(*C.int)(unsafe.Pointer(error)))
}

func Close(dtp DTraceHandle) {
	C.dtrace_close(dtp)
}

func ErrMsg(dtp DTraceHandle, error int) string {
	return C.GoString(C.dtrace_errmsg(dtp, C.int(error)))
}

func Stop(dtp DTraceHandle) {
	C.dtrace_stop(dtp)
}

func Go(dtp DTraceHandle) int {
	return int(C.dtrace_go(dtp))
}

func SetOpt(dtp DTraceHandle, opt string, val string) int {
	return int(C.dtrace_setopt(dtp, C.CString(opt), C.CString(val)))
}
