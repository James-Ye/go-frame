// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const KEY_READ REGSAM = 0x20019
const KEY_WRITE REGSAM = 0x20006

const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

const (
	ERROR_NO_MORE_ITEMS = 259
)

type (
	ACCESS_MASK uint32
	HKEY        windows.Handle
	REGSAM      uint32
)

const (
	REG_NONE      = uint64(0) // No value type
	REG_SZ        = 1         // Unicode nul terminated string
	REG_EXPAND_SZ = 2         // Unicode nul terminated string
	// (with environment variable references)
	REG_BINARY                     = 3 // Free form binary
	REG_uint32                     = 4 // 32-bit number
	REG_uint32_LITTLE_ENDIAN       = 4 // 32-bit number (same as REG_uint32)
	REG_uint32_BIG_ENDIAN          = 5 // 32-bit number
	REG_LINK                       = 6 // Symbolic Link (unicode)
	REG_MULTI_SZ                   = 7 // Multiple Unicode strings
	REG_RESOURCE_LIST              = 8 // Resource list in the resource map
	REG_FULL_RESOURCE_DESCRIPTOR   = 9 // Resource list in the hardware description
	REG_RESOURCE_REQUIREMENTS_LIST = 10
	REG_QWORD                      = 11 // 64-bit number
	REG_QWORD_LITTLE_ENDIAN        = 11 // 64-bit number (same as REG_QWORD)

)

var (
	// Library
	libadvapi32 *windows.LazyDLL

	// Functions
	regEnumValue        *windows.LazyProc
	regSetValueEx       *windows.LazyProc
	regConnectRegistryW *windows.LazyProc
	regCreateKeyExW     *windows.LazyProc
	regDeleteKeyW       *windows.LazyProc
	regDeleteValueW     *windows.LazyProc
	regLoadMUIStringW   *windows.LazyProc
	regOpenKeyExW       *windows.LazyProc
	regCloseKey         *windows.LazyProc
)

func init() {
	// Library
	libadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	// Functions
	regEnumValue = libadvapi32.NewProc("RegEnumValueW")
	regSetValueEx = libadvapi32.NewProc("RegSetValueExW")
	regConnectRegistryW = libadvapi32.NewProc("RegConnectRegistryW")
	regCreateKeyExW = libadvapi32.NewProc("RegCreateKeyExW")
	regDeleteKeyW = libadvapi32.NewProc("RegDeleteKeyW")
	regDeleteValueW = libadvapi32.NewProc("RegDeleteValueW")
	regLoadMUIStringW = libadvapi32.NewProc("RegLoadMUIStringW")
	regOpenKeyExW = libadvapi32.NewProc("RegOpenKeyExW")
	regCloseKey = libadvapi32.NewProc("RegCloseKey")

}

func RegConnectRegistry(machinename *uint16, key windows.Handle, result *windows.Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall(regConnectRegistryW.Addr(), 3, uintptr(unsafe.Pointer(machinename)), uintptr(key), uintptr(unsafe.Pointer(result)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegCreateKeyEx(key windows.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *windows.Handle, disposition *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(regCreateKeyExW.Addr(), 9, uintptr(key), uintptr(unsafe.Pointer(subkey)), uintptr(reserved), uintptr(unsafe.Pointer(class)), uintptr(options), uintptr(desired), uintptr(unsafe.Pointer(sa)), uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(disposition)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegDeleteKey(key windows.Handle, subkey *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(regDeleteKeyW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(subkey)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegDeleteValue(key windows.Handle, name *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(regDeleteValueW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(name)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegEnumValue(key windows.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(regEnumValue.Addr(), 8, uintptr(key), uintptr(index), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(valtype)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(buflen)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegLoadMUIString(key windows.Handle, name *uint16, buf *uint16, buflen uint32, buflenCopied *uint32, flags uint32, dir *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall9(regLoadMUIStringW.Addr(), 7, uintptr(key), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buf)), uintptr(buflen), uintptr(unsafe.Pointer(buflenCopied)), uintptr(flags), uintptr(unsafe.Pointer(dir)), 0, 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegSetValueEx(key windows.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall6(regSetValueEx.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(valueName)), uintptr(reserved), uintptr(vtype), uintptr(unsafe.Pointer(buf)), uintptr(bufsize))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegOpenKeyEx(key windows.Handle, SubKey *uint16, Options uint32, samDesired uint32, phkResult *windows.Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall6(regOpenKeyExW.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(SubKey)), uintptr(Options), uintptr(samDesired), uintptr(unsafe.Pointer(phkResult)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegCloseKey(key windows.Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall(regCloseKey.Addr(), 1, uintptr(key), 0, 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}
