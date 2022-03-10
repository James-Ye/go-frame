package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

/*===================================================================
==   Defines
=====================================================================*/

/*
 *  Specifies the current server
 */
const (
	WTS_CURRENT_SERVER        = windows.Handle(0)
	WTS_CURRENT_SERVER_HANDLE = windows.Handle(0)
	WTS_CURRENT_SERVER_NAME   = 0
)

type WTS_CONNECTSTATE_CLASS uint

/*
 *  Specifies the current session (SessionId)
 */
const WTS_CURRENT_SESSION uint32 = 0xFFFFFFFF

const (
	WTSActive       WTS_CONNECTSTATE_CLASS = 0 // User logged on to WinStation
	WTSConnected    WTS_CONNECTSTATE_CLASS = 1 // WinStation connected to client
	WTSConnectQuery WTS_CONNECTSTATE_CLASS = 2 // In the process of connecting to client
	WTSShadow       WTS_CONNECTSTATE_CLASS = 3 // Shadowing another WinStation
	WTSDisconnected WTS_CONNECTSTATE_CLASS = 4 // WinStation logged on without client
	WTSIdle         WTS_CONNECTSTATE_CLASS = 5 // Waiting for client to connect
	WTSListen       WTS_CONNECTSTATE_CLASS = 6 // WinStation is listening for connection
	WTSReset        WTS_CONNECTSTATE_CLASS = 7 // WinStation is being reset
	WTSDown         WTS_CONNECTSTATE_CLASS = 8 // WinStation is down due to error
	WTSInit         WTS_CONNECTSTATE_CLASS = 9 // WinStation in initialization
)

type WTS_INFO_CLASS uint32

const (
	WTSInitialProgram     WTS_INFO_CLASS = 0
	WTSApplicationName    WTS_INFO_CLASS = 1
	WTSWorkingDirectory   WTS_INFO_CLASS = 2
	WTSOEMId              WTS_INFO_CLASS = 3
	WTSSessionId          WTS_INFO_CLASS = 4
	WTSUserName           WTS_INFO_CLASS = 5
	WTSWinStationName     WTS_INFO_CLASS = 6
	WTSDomainName         WTS_INFO_CLASS = 7
	WTSConnectState       WTS_INFO_CLASS = 8
	WTSClientBuildNumber  WTS_INFO_CLASS = 9
	WTSClientName         WTS_INFO_CLASS = 10
	WTSClientDirectory    WTS_INFO_CLASS = 11
	WTSClientProductId    WTS_INFO_CLASS = 12
	WTSClientHardwareId   WTS_INFO_CLASS = 13
	WTSClientAddress      WTS_INFO_CLASS = 14
	WTSClientDisplay      WTS_INFO_CLASS = 15
	WTSClientProtocolType WTS_INFO_CLASS = 16
	WTSIdleTime           WTS_INFO_CLASS = 17
	WTSLogonTime          WTS_INFO_CLASS = 18
	WTSIncomingBytes      WTS_INFO_CLASS = 19
	WTSOutgoingBytes      WTS_INFO_CLASS = 20
	WTSIncomingFrames     WTS_INFO_CLASS = 21
	WTSOutgoingFrames     WTS_INFO_CLASS = 22
)

type WTS_SESSION_INFOW struct {
	SessionId       uint32  // session id
	PWinStationName *uint16 // name of WinStation this session is
	// connected to
	State WTS_CONNECTSTATE_CLASS // connection state (see enum)
}

var (
	// Library
	libWtsapi32 *windows.LazyDLL

	wTSEnumerateSessionsW       *windows.LazyProc
	wTSQueryUserToken           *windows.LazyProc
	wTSQuerySessionInformationW *windows.LazyProc
	wTSFreeMemory               *windows.LazyProc
)

func init() {
	// Library
	libWtsapi32 = windows.NewLazySystemDLL("Wtsapi32.dll")

	wTSEnumerateSessionsW = libWtsapi32.NewProc("WTSEnumerateSessionsW")
	wTSQueryUserToken = libWtsapi32.NewProc("WTSQueryUserToken")
	wTSQuerySessionInformationW = libWtsapi32.NewProc("WTSQuerySessionInformationW")
	wTSFreeMemory = libWtsapi32.NewProc("WTSFreeMemory")

}

func WTSQueryUserToken(SessionId uint32, hToken windows.Handle) (bool, syscall.Errno) {
	r, _, err := syscall.Syscall(wTSQueryUserToken.Addr(), 2,
		uintptr(SessionId),
		uintptr(hToken),
		0)

	return r != 0, err
}

func WTSEnumerateSessions(hServer windows.Handle, Reserved uint32, Version uint32) ([](*WTS_SESSION_INFOW), uint32, bool) {
	pSessionInfo := [](*WTS_SESSION_INFOW){nil}
	var count uint32 = 0
	r, _, _ := syscall.Syscall6(wTSEnumerateSessionsW.Addr(), 5,
		uintptr(hServer),
		uintptr(Reserved),
		uintptr(Version),
		uintptr(unsafe.Pointer(&pSessionInfo[0])),
		uintptr(count),
		0)

	return pSessionInfo, count, r != 0
}

func WTSQuerySessionInformation(hServer windows.Handle, SessionId uint32, WTSInfoClass WTS_INFO_CLASS) (*byte, int, bool) {
	var pBuffer *byte
	var BytesReturned uint32 = 0
	r, _, _ := syscall.Syscall6(wTSQuerySessionInformationW.Addr(), 5,
		uintptr(hServer),
		uintptr(SessionId),
		uintptr(WTSInfoClass),
		uintptr(unsafe.Pointer(&pBuffer)),
		uintptr(unsafe.Pointer(&BytesReturned)),
		0)

	return pBuffer, int(BytesReturned), r != 0
}

func WTSFreeMemory(pMemory uintptr) {
	syscall.Syscall(wTSFreeMemory.Addr(), 1,
		pMemory,
		0,
		0)
}
