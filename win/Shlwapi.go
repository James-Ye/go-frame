package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	URL_UNESCAPE_INPLACE = 0x00100000
)

var (
	// Library
	libShlwapi *windows.LazyDLL

	pathQuoteSpaces     *windows.LazyProc
	pathRemoveFileSpecA *windows.LazyProc
	pathRemoveFileSpecW *windows.LazyProc
	sHGetValue          *windows.LazyProc
	urlUnescapeA        *windows.LazyProc
	urlUnescapeW        *windows.LazyProc
)

func init() {
	// Library
	libShlwapi = windows.NewLazySystemDLL("Shlwapi.dll")

	pathQuoteSpaces = libShlwapi.NewProc("PathQuoteSpacesW")
	pathRemoveFileSpecA = libShlwapi.NewProc("PathRemoveFileSpecA")
	pathRemoveFileSpecW = libShlwapi.NewProc("PathRemoveFileSpecW")
	sHGetValue = libShlwapi.NewProc("SHGetValueA")
	urlUnescapeA = libShlwapi.NewProc("UrlUnescapeA")
	urlUnescapeW = libShlwapi.NewProc("UrlUnescapeW")
}

func SHGetValueA(hkey HKEY, subKey string, value string, pdwType *uint32, pvData unsafe.Pointer, pcbData *uint32) syscall.Errno {
	lpSubkey, _ := syscall.UTF16PtrFromString(subKey)
	lpValue, _ := syscall.UTF16PtrFromString(value)
	_, _, err := syscall.Syscall6(sHGetValue.Addr(), 6,
		uintptr(hkey),
		uintptr(unsafe.Pointer(lpSubkey)),
		uintptr(unsafe.Pointer(lpValue)),
		uintptr(unsafe.Pointer(pdwType)),
		uintptr(pvData),
		uintptr(unsafe.Pointer(pcbData)))

	return err
}

func PathQuoteSpaces(szPath *string) bool {
	sz, _ := syscall.UTF16FromString(*szPath)
	r, _, _ := syscall.Syscall(pathQuoteSpaces.Addr(), 1,
		uintptr(unsafe.Pointer(&sz[0])),
		0,
		0)

	if r != 0 {
		*szPath = syscall.UTF16ToString(sz)
	}

	return r != 0
}

func PathRemoveFileSpecA(szPath *byte) bool {
	r, _, _ := syscall.Syscall(pathRemoveFileSpecA.Addr(), 1,
		uintptr(unsafe.Pointer(szPath)),
		0,
		0)

	return r != 0
}

func PathRemoveFileSpecW(szPath *uint16) bool {
	r, _, _ := syscall.Syscall(pathRemoveFileSpecW.Addr(), 1,
		uintptr(unsafe.Pointer(szPath)),
		0,
		0)

	return r != 0
}

func UrlUnescapeA(pszUrl *byte, pszUnescaped *byte, pcchUnescaped *uint32, dwFlags uint32) HRESULT {
	r, _, _ := syscall.Syscall6(urlUnescapeA.Addr(), 4,
		uintptr(unsafe.Pointer(pszUrl)),
		uintptr(unsafe.Pointer(pszUnescaped)),
		uintptr(unsafe.Pointer(pcchUnescaped)),
		uintptr(dwFlags),
		0,
		0)

	return HRESULT(r)
}

func UrlUnescapeW(pszUrl *uint16, pszUnescaped *uint16, pcchUnescaped *uint32, dwFlags uint32) HRESULT {
	r, _, _ := syscall.Syscall6(urlUnescapeW.Addr(), 4,
		uintptr(unsafe.Pointer(pszUrl)),
		uintptr(unsafe.Pointer(pszUnescaped)),
		uintptr(unsafe.Pointer(pcchUnescaped)),
		uintptr(dwFlags),
		0,
		0)

	return HRESULT(r)
}
