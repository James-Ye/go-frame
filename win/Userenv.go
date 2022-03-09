package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libUserenv *windows.LazyDLL

	createEnvironmentBlock  *windows.LazyProc
	destroyEnvironmentBlock *windows.LazyProc
	unloadUserProfile       *windows.LazyProc
)

func init() {
	// Library
	libUserenv = windows.NewLazySystemDLL("Userenv.dll")

	createEnvironmentBlock = libUserenv.NewProc("CreateEnvironmentBlock")
	destroyEnvironmentBlock = libUserenv.NewProc("DestroyEnvironmentBlock")
	unloadUserProfile = libUserenv.NewProc("UnloadUserProfile")
}

func CreateEnvironmentBlock(hToken windows.Handle, bInherit bool) (unsafe.Pointer, bool) {
	lpEnvironment := new(unsafe.Pointer)
	ret, _, _ := syscall.Syscall(createEnvironmentBlock.Addr(), 3,
		uintptr(unsafe.Pointer(lpEnvironment)),
		uintptr(hToken),
		uintptr(BoolToBOOL(bInherit)))

	return *lpEnvironment, ret != 0
}

//=============================================================================
//
// DestroyEnvironmentBlock
//
// Frees environment variables created by CreateEnvironmentBlock
//
// lpEnvironment  -  A pointer to the environment block
//
// Returns:  true if successful
//           false if not.  syscall.Errno for more details
//
//=============================================================================
func DestroyEnvironmentBlock(lpEnvironment unsafe.Pointer) bool {
	ret, _, _ := syscall.Syscall(destroyEnvironmentBlock.Addr(), 1,
		uintptr(lpEnvironment),
		0,
		0)

	return ret != 0
}

//=============================================================================
//
// UnloadUserProfile
//
// Unloads a user's profile that was loaded by LoadUserProfile()
//
// hToken        -  Token for the user, returned from LogonUser()
// hProfile      -  hProfile member of the PROFILEINFO structure
//
// Returns:  TRUE if successful
//           FALSE if not.  Call GetLastError() for more details
//
// Note:     The caller of this function must have admin privileges on the machine.
//
//=============================================================================
func UnloadUserProfile(hToken windows.Handle, hProfile windows.Handle) bool {
	ret, _, _ := syscall.Syscall(unloadUserProfile.Addr(), 2,
		uintptr(hToken),
		uintptr(hProfile),
		0)

	return ret != 0
}
