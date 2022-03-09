package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libWs2_32 *windows.LazyDLL

	inet_addr *windows.LazyProc
	htons     *windows.LazyProc
	ntohs     *windows.LazyProc
	inet_ntoa *windows.LazyProc
)

func init() {
	// Library
	libWs2_32 = windows.NewLazySystemDLL("Ws2_32.dll")

	inet_addr = libWs2_32.NewProc("inet_addr")
	htons = libWs2_32.NewProc("htons")
	inet_ntoa = libWs2_32.NewProc("inet_ntoa")
}

func Inet_addr(cp string) uint32 {
	lpcp, _ := syscall.UTF16PtrFromString(cp)
	ret, _, _ := syscall.Syscall(inet_addr.Addr(), 1,
		uintptr(unsafe.Pointer(lpcp)),
		0,
		0)

	return uint32(ret)
}

func Htons(hostshort uint16) uint16 {
	ret, _, _ := syscall.Syscall(htons.Addr(), 1,
		uintptr(hostshort),
		0,
		0)

	return uint16(ret)
}

func Inet_ntoa(in In_addr) *byte {
	ret, _, _ := syscall.Syscall(inet_ntoa.Addr(), 1,
		uintptr(unsafe.Pointer(&in)),
		0,
		0)

	return (*byte)(unsafe.Pointer(ret))
}
