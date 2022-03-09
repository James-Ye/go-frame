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

const (
	L2_PROFILE_MAX_NAME_LENGTH = 256
	WLAN_MAX_NAME_LENGTH       = L2_PROFILE_MAX_NAME_LENGTH
)

type WLAN_INTERFACE_STATE int32

const (
	Wlan_interface_state_not_ready             WLAN_INTERFACE_STATE = 0
	Wlan_interface_state_connected             WLAN_INTERFACE_STATE = 1
	Wlan_interface_state_ad_hoc_network_formed WLAN_INTERFACE_STATE = 2
	Wlan_interface_state_disconnecting         WLAN_INTERFACE_STATE = 3
	Wlan_interface_state_disconnected          WLAN_INTERFACE_STATE = 4
	Wlan_interface_state_associating           WLAN_INTERFACE_STATE = 5
	Wlan_interface_state_discovering           WLAN_INTERFACE_STATE = 6
	Wlan_interface_state_authenticating        WLAN_INTERFACE_STATE = 7
)

// struct WLAN_INTERFACE_INFO defines the basic information for an interface
type WLAN_INTERFACE_INFO struct {
	InterfaceGuid        windows.GUID
	InterfaceDescription [WLAN_MAX_NAME_LENGTH]uint16
	State                WLAN_INTERFACE_STATE
}

type WLAN_INTERFACE_INFO_LIST struct {
	NumberOfItems uint32
	Index         uint32

	// #ifdef __midl
	//     [unique, size_is(dwNumberOfItems)] WLAN_INTERFACE_INFO InterfaceInfo[*];
	// #else
	InterfaceInfo [1]WLAN_INTERFACE_INFO
	// #endif

}

// the callback function for notifications
type WLAN_NOTIFICATION_CALLBACK func(*L2_NOTIFICATION_DATA, uintptr)

// This structure is the notification structure which needs to be filled in by each component
// The interface guid is filled in only by the AC
// The NotificationSource signifies the type above
// and NotificationCode is one of the enum values
type L2_NOTIFICATION_DATA struct {
	NotificationSource uint32
	NotificationCode   uint32
	InterfaceGuid      windows.GUID
	DataSize           uint32
	PData              uintptr
}

type (
	WLAN_REASON_CODE    uint32
	WLAN_SIGNAL_QUALITY uint32
)

const WLAN_MAX_PHY_TYPE_NUMBER = 8

//
// struct WLAN_AVAILABLE_NETWORK defines information needed for an available network
type WLAN_AVAILABLE_NETWORK struct {
	ProfileName              [WLAN_MAX_NAME_LENGTH]uint16
	Dot11Ssid                DOT11_SSID
	Dot11BssType             DOT11_BSS_TYPE
	NumberOfBssids           uint32
	NetworkConnectable       bool
	WlanNotConnectableReason WLAN_REASON_CODE
	NumberOfPhyTypes         uint32
	Dot11PhyTypes            [WLAN_MAX_PHY_TYPE_NUMBER]DOT11_PHY_TYPE
	// bMorePhyTypes is set to TRUE if the PHY types for the network
	// exceeds WLAN_MAX_PHY_TYPE_NUMBER.
	// In this case, uNumerOfPhyTypes is WLAN_MAX_PHY_TYPE_NUMBER and the
	// first WLAN_MAX_PHY_TYPE_NUMBER PHY types are returned.
	MorePhyTypes                bool
	WlanSignalQuality           WLAN_SIGNAL_QUALITY
	SecurityEnabled             bool
	Dot11DefaultAuthAlgorithm   DOT11_AUTH_ALGORITHM
	Dot11DefaultCipherAlgorithm DOT11_CIPHER_ALGORITHM
	Flags                       uint32
	Reserved                    uint32
}

type WLAN_AVAILABLE_NETWORK_LIST struct {
	NumberOfItems uint32
	Index         uint32

	// #ifdef __midl
	//     [unique, size_is(dwNumberOfItems)] WLAN_AVAILABLE_NETWORK Network[*];
	// #else
	Network [1]WLAN_AVAILABLE_NETWORK
	// #endif
}

var (
	// Library
	libWlanapi *windows.LazyDLL

	wlanCloseHandle             *windows.LazyProc
	wlanFreeMemory              *windows.LazyProc
	wlanGetAvailableNetworkList *windows.LazyProc
	wlanOpenHandle              *windows.LazyProc
)

func init() {
	// Library
	libWlanapi = windows.NewLazySystemDLL("Wlanapi.dll")

	wlanCloseHandle = libWlanapi.NewProc("WlanCloseHandle")
	wlanFreeMemory = libWlanapi.NewProc("WlanFreeMemory")
	wlanGetAvailableNetworkList = libWlanapi.NewProc("WlanGetAvailableNetworkList")
	wlanOpenHandle = libWlanapi.NewProc("WlanOpenHandle")

}

func WlanOpenHandle(dwClientVersion uint32, pReserved unsafe.Pointer, pdwNegotiatedVersion *uint32, phClientHandle *windows.Handle) syscall.Errno {
	_, _, err := syscall.Syscall6(wlanOpenHandle.Addr(), 4,
		uintptr(dwClientVersion),
		uintptr(pReserved),
		uintptr(unsafe.Pointer(pdwNegotiatedVersion)),
		uintptr(unsafe.Pointer(phClientHandle)),
		0,
		0)

	return err
}

func WlanGetAvailableNetworkList(hClientHandle windows.Handle, pInterfaceGuid *windows.GUID, dwFlags uint32, pReserved unsafe.Pointer, ppAvailableNetworkList **WLAN_AVAILABLE_NETWORK_LIST) syscall.Errno {
	_, _, err := syscall.Syscall6(wlanGetAvailableNetworkList.Addr(), 5,
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGuid)),
		uintptr(dwFlags),
		uintptr(pReserved),
		uintptr(unsafe.Pointer(ppAvailableNetworkList)),
		0)

	return err
}

func WlanFreeMemory(pMemory unsafe.Pointer) {
	syscall.Syscall(wlanFreeMemory.Addr(), 1,
		uintptr(pMemory),
		0,
		0)
}

func WlanCloseHandle(hClientHandle windows.Handle, pReserved unsafe.Pointer) uint32 {
	r, _, _ := syscall.Syscall(wlanCloseHandle.Addr(), 2,
		uintptr(hClientHandle),
		uintptr(pReserved),
		0)

	return uint32(r)
}
