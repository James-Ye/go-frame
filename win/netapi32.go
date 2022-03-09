package win

import (
	"syscall"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libbetapi32 *windows.LazyDLL

	netGetJoinInformation *windows.LazyProc
)

func init() {
	// Library
	libbetapi32 = windows.NewLazySystemDLL("Netapi32.dll")

	netGetJoinInformation = libbetapi32.NewProc("NetGetJoinInformation")
}

func NetGetJoinInformation(lpServer *uint16) (error, string, uint32) {
	nameBuffer := []uint16{0}
	pNameBuffer := &nameBuffer[0]
	var status uint32
	err := windows.NetGetJoinInformation(lpServer, &pNameBuffer, &status)

	return err, syscall.UTF16ToString(nameBuffer), status
}
