package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const MAX_PATH = 260

type (
	ATOM          uint16
	HGLOBAL       windows.Handle
	HINSTANCE     windows.Handle
	LCID          uint32
	LCTYPE        uint32
	LANGID        uint16
	HMODULE       uintptr
	HWINEVENTHOOK windows.Handle
	HRSRC         uintptr
)

type OSVERSIONINFOA struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	SzCSDVersion      [128]byte // Maintenance string for PSS usage
}

type OSVERSIONINFOW struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	SzCSDVersion      [128]uint16 // Maintenance string for PSS usage
}

type OSVERSIONINFO OSVERSIONINFOW

var (
	// Library
	libkernel32 *windows.LazyDLL

	createProcessA                 *windows.LazyProc
	getTickCount                   *windows.LazyProc
	getUserNameW                   *windows.LazyProc
	getVersionExA                  *windows.LazyProc
	getVersionExW                  *windows.LazyProc
	globalAlloc                    *windows.LazyProc
	globalFree                     *windows.LazyProc
	impersonateLoggedOnUser        *windows.LazyProc
	logonUser                      *windows.LazyProc
	wow64DisableWow64FsRedirection *windows.LazyProc
	wow64RevertWow64FsRedirection  *windows.LazyProc
)

func init() {
	// Library
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	createProcessA = libkernel32.NewProc("CreateProcessA")
	getTickCount = libkernel32.NewProc("GetTickCount")
	getUserNameW = libkernel32.NewProc("GetUserNameW")
	getVersionExA = libkernel32.NewProc("GetVersionExA")
	getVersionExW = libkernel32.NewProc("GetVersionExW")
	globalAlloc = libkernel32.NewProc("GlobalAlloc")
	globalFree = libkernel32.NewProc("GlobalFree")
	impersonateLoggedOnUser = libkernel32.NewProc("ImpersonateLoggedOnUser")
	logonUser = libkernel32.NewProc("LogonUserA")
	wow64DisableWow64FsRedirection = libkernel32.NewProc("Wow64DisableWow64FsRedirection")
	wow64RevertWow64FsRedirection = libkernel32.NewProc("Wow64RevertWow64FsRedirection")
}

func CreateProcessA(strApplicationName string, strCommandLine string,
	lpProcessAttributes *windows.SecurityAttributes, lpThreadAttributes *windows.SecurityAttributes,
	bInheritHandles bool, CreationFlags uint32, lpEnvironment unsafe.Pointer, strCurrentDirectory string,
	lpStartupInfo *windows.StartupInfo, lpProcessInformation *windows.ProcessInformation) (bool, windows.Errno) {

	lpApplicationName, _ := windows.UTF16PtrFromString(strApplicationName)
	lpCommandLine, _ := windows.UTF16PtrFromString(strCommandLine)
	lpCurrentDirectory, _ := windows.UTF16PtrFromString(strCurrentDirectory)

	ret, _, err := syscall.Syscall12(createProcessA.Addr(), 10,
		uintptr(unsafe.Pointer(lpApplicationName)),
		uintptr(unsafe.Pointer(lpCommandLine)),
		uintptr(unsafe.Pointer(lpProcessAttributes)),
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		uintptr(BoolToBOOL(bInheritHandles)),
		uintptr(CreationFlags),
		uintptr(lpEnvironment),
		uintptr(unsafe.Pointer(lpCurrentDirectory)),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
		0,
		0)

	return ret != 0, err
}

func GetComputerName() (string, bool) {
	var nSize uint32 = 0
	lpBuffer := []uint16{0}
	err := syscall.GetComputerName(&lpBuffer[0], &nSize)

	return syscall.UTF16ToString(lpBuffer), err == nil
}

func GetTickCount() uint32 {
	r, _, _ := syscall.Syscall(getTickCount.Addr(), 0,
		0,
		0,
		0)

	return uint32(r)
}

func GetUserName() (*uint16, int, bool, syscall.Errno) {
	var buffer []uint16
	var ReturnLength uint32 = 0
	r, _, err := syscall.Syscall(getUserNameW.Addr(), 2,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&ReturnLength)),
		0)

	return &buffer[0], int(ReturnLength), r != 0, err
}

func GetVersionExA(ov *OSVERSIONINFOA) bool {
	ret, _, _ := syscall.Syscall(getVersionExA.Addr(), 1,
		uintptr(unsafe.Pointer(ov)),
		0,
		0)
	return ret != 0
}

func GetVersionEx(ov *OSVERSIONINFOW) bool {
	ret, _, _ := syscall.Syscall(getVersionExW.Addr(), 1,
		uintptr(unsafe.Pointer(ov)),
		0,
		0)
	return ret != 0
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalAlloc.Addr(), 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalFree.Addr(), 1,
		uintptr(hMem),
		0,
		0)

	return HGLOBAL(ret)
}

func ImpersonateLoggedOnUser(hToken windows.Token) (bool, windows.Errno) {
	ret, _, err := syscall.Syscall(impersonateLoggedOnUser.Addr(), 1,
		uintptr(hToken),
		0,
		0)

	return ret != 0, err
}

func LogonUser(lpszUsername *uint16, lpszDomain *uint16, lpszPassword *uint16, LogonType uint32, LogonProvider windows.Handle) (windows.Handle, bool) {
	var hToken windows.Handle
	r, _, _ := syscall.Syscall6(logonUser.Addr(), 6,
		uintptr(unsafe.Pointer(lpszUsername)),
		uintptr(unsafe.Pointer(lpszDomain)),
		uintptr(unsafe.Pointer(lpszPassword)),
		uintptr(LogonType),
		uintptr(LogonProvider),
		uintptr(unsafe.Pointer(&hToken)))

	return hToken, r != 0
}

func Wow64DisableWow64FsRedirection() (unsafe.Pointer, bool) {
	var OldValue unsafe.Pointer
	r, _, _ := syscall.Syscall(wow64DisableWow64FsRedirection.Addr(), 1,
		uintptr(unsafe.Pointer(&OldValue)),
		0,
		0)

	return OldValue, r != 0
}

func Wow64RevertWow64FsRedirection(OldValue unsafe.Pointer) bool {
	r, _, _ := syscall.Syscall(wow64RevertWow64FsRedirection.Addr(), 1,
		uintptr(OldValue),
		0,
		0)

	return r != 0
}

func WaitForSingleObject(hHandle windows.Handle, dwMilliseconds uint32) uint32 {
	event, err := windows.WaitForSingleObject(hHandle, uint32(dwMilliseconds))
	if err != nil {
		event = syscall.WAIT_OBJECT_0
	}

	return uint32(event)
}
