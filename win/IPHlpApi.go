package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libiphlpapi *windows.LazyDLL

	getInterfaceInfo *windows.LazyProc
	ipReleaseAddress *windows.LazyProc
	ipRenewAddress   *windows.LazyProc
)

func init() {
	// Library
	libiphlpapi = windows.NewLazySystemDLL("Iphlpapi.dll")

	getInterfaceInfo = libiphlpapi.NewProc("GetInterfaceInfo")
	ipReleaseAddress = libiphlpapi.NewProc("IpReleaseAddress")
	ipRenewAddress = libiphlpapi.NewProc("IpRenewAddress")
}

func GetInterfaceInfo(pIfTable *IP_INTERFACE_INFO, dwOutBufLen *uint32) uint32 {
	ret, _, _ := syscall.Syscall(getInterfaceInfo.Addr(), 2,
		uintptr(unsafe.Pointer(pIfTable)),
		uintptr(unsafe.Pointer(dwOutBufLen)),
		0)

	return uint32(ret)
}

func IpReleaseAddress(AdapterInfo *IP_ADAPTER_INDEX_MAP) uint32 {
	ret, _, _ := syscall.Syscall(ipReleaseAddress.Addr(), 1,
		uintptr(unsafe.Pointer(AdapterInfo)),
		0,
		0)

	return uint32(ret)
}

func IpRenewAddress(AdapterInfo *IP_ADAPTER_INDEX_MAP) uint32 {
	ret, _, _ := syscall.Syscall(ipRenewAddress.Addr(), 1,
		uintptr(unsafe.Pointer(AdapterInfo)),
		0,
		0)

	return uint32(ret)
}
