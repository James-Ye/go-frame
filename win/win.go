package win

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	FALSE = 0
	TRUE  = 1
)

const is64bit bool = unsafe.Sizeof(uintptr(0)) == 8

type (
	BOOL    uint32
	HRESULT uint32
	// time_t  uint64 //64位
	// Char    uint8
	// UChar   uint8
	// WChar   uint16
	// Short   int16
	// UShort  uint16
	// Long    int32
	// ULong   uint32
	// WORD    uint16

	WPARAM  *uint
	LPARAM  *int32
	LRESULT *int32
)

func FAILED(hr HRESULT) bool {
	return hr < 0
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func LOWORD(dw uint32) uint16 {
	return uint16(dw)
}

func HIWORD(dw uint32) uint16 {
	return uint16(dw >> 16 & 0xffff)
}

// type WCHAR = wchar_t
// type wchar_t = uint16

// UTF16toString converts a pointer to a UTF16 string into a Go string.
func UTF16toString(p *uint16) string {
	return syscall.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}

func StrPtr(s string) uintptr {
	systemNameUTF16String, err := syscall.UTF16PtrFromString(s)
	if err == nil {
		// 这里转换的时候出错,则不继续执行 OR 赋值用本地的
		tmp := utf16.Encode([]rune("\x00"))
		systemNameUTF16String = &tmp[0]
	}

	return uintptr(unsafe.Pointer(systemNameUTF16String))
}

func CharsPtr(s string) uintptr {
	bPtr, err := syscall.BytePtrFromString(s)
	if err != nil {
		return uintptr(0) // 这么写肯定不太对 @TODO
	}
	return uintptr(unsafe.Pointer(bPtr))
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func Touint32(num int) uint32 {
	var ret uint32 = 0xFFFFFFFF
	if num < 0 {
		ret = uint32(int(ret) + num + 1)
	} else {
		ret = uint32(num)
	}
	return ret
}

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}
