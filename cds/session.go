package cds

import (
	"strings"
	"sync"
	"syscall"
	"unsafe"

	"github.com/James-Ye/go-frame/atl"
	"github.com/James-Ye/go-frame/logger"
	"github.com/James-Ye/go-frame/process_hlp"
	"github.com/James-Ye/go-frame/utils"
	"github.com/James-Ye/go-frame/win"
	"golang.org/x/sys/windows"
)

var g_Session *Session
var once sync.Once

func GetInstance() *Session {
	once.Do(func() {
		g_Session = &Session{}
	})
	return g_Session
}

//会话、Winstation、Desktop相关功能
type Session struct {
}

//The session identifier of the session that is attached to the physical
//console. If there is no session attached to the physical console, (for
//example, if the physical console session is in the process of being attached
//or detached), this function returns 0xFFFFFFFF.
func (se Session) GetActiveConsoleSessionId() uint32 {
	retSessionId := windows.WTSGetActiveConsoleSessionId()

	vpSessionInfo, sessionCnt, ok := win.WTSEnumerateSessions(win.WTS_CURRENT_SERVER_HANDLE, 0, 1)

	if !ok {
		// logger.Trace("WTSEnumerateSessions failed")
		return uint32(retSessionId)
	}

	for i := 0; i < int(sessionCnt); i++ {
		if vpSessionInfo[i].State == win.WTSActive {
			retSessionId = vpSessionInfo[i].SessionId
			//logger.Trace(L"WTSEnumerateSessions  active sessionid:%d", dwRetSessionId);
			break
		}
	}
	//logger.Trace(L"WTSEnumerateSessions  return :%d", dwRetSessionId);
	return uint32(retSessionId)
}

func (se Session) QuerySessionTokenCDS(sessionId uint32) windows.Handle {
	var TokenHandle windows.Handle = 0

	var Entry windows.ProcessEntry32
	Entry.Size = uint32(unsafe.Sizeof(Entry))

	SnapshotHandle, err := windows.CreateToolhelp32Snapshot(win.TH32CS_SNAPPROCESS, 0)

	if err != nil {
		return windows.InvalidHandle
	}

	for err := windows.Process32First(SnapshotHandle, &Entry); err == nil; err = windows.Process32Next(SnapshotHandle, &Entry) {
		strszExeFile := syscall.UTF16ToString(Entry.ExeFile[:])
		if strings.Compare(strszExeFile, "explorer.exe") != 0 && strings.Compare(strszExeFile, "userinit.exe") != 0 {
			continue
		}

		ProcessHandle, _ := windows.OpenProcess(uint32(win.MAXIMUM_ALLOWED), false, Entry.ProcessID)

		if ProcessHandle != windows.InvalidHandle {
			continue
		}

		var ProcessToken atl.AccessToken

		if ProcessToken.GetProcessToken(uint32(win.MAXIMUM_ALLOWED), ProcessHandle) {
			var dwSessionId uint32 = 0

			if ProcessToken.GetTerminalServicesSessionId(&dwSessionId) && dwSessionId == sessionId {
				var Sid atl.GSID

				if ProcessToken.GetUser(&Sid) {
					strAccountName := Sid.AccountName()

					if strings.Compare(strAccountName, "SYSTEM") != 0 {
						TokenHandle = ProcessToken.Detach()
						//logger.Trace("in QuerySessionToken, get token from %S, accountname %S", Entry.szExeFile, (LPCTSTR)strAccountName);
					}
				}
			}
		}

		windows.CloseHandle(ProcessHandle)

		if TokenHandle != windows.InvalidHandle {
			break
		}
	}

	windows.CloseHandle(SnapshotHandle)

	return TokenHandle
}

func (se Session) IsLogon() bool {
	htoken := se.QueryUserToken(se.GetActiveConsoleSessionId())
	if htoken != 0 && htoken != windows.InvalidHandle {
		windows.CloseHandle(htoken)
		return true
	}

	htoken = se.QuerySessionTokenCDS(se.GetActiveConsoleSessionId())
	if htoken == windows.InvalidHandle {
		return false
	} else {
		windows.CloseHandle(htoken)
		return true
	}
}

func (se Session) QueryUserToken(sessionId uint32) windows.Handle {
	hToken := windows.InvalidHandle
	libwtsapi32 := windows.NewLazySystemDLL("Wtsapi32.dll")
	wTSQueryUserToken := libwtsapi32.NewProc("WTSQueryUserToken")
	//winxp or above
	if wTSQueryUserToken != nil {
		var ok bool
		var err syscall.Errno
		ok, hToken, err = func() (bool, windows.Handle, windows.Errno) {
			var token windows.Handle = windows.InvalidHandle
			r, _, err := syscall.Syscall(wTSQueryUserToken.Addr(), 2, uintptr(sessionId), uintptr(unsafe.Pointer(&token)), 0)
			return r != 0, token, err
		}()
		if !ok {
			logger.Trace("QueryUserToken failed, err = %d", err)
		}
	} else { //win2000
		//get explorer's path
		strSysPath, _ := windows.GetWindowsDirectory()
		strPath := strSysPath + "explorer.exe"

		dwPid := process_hlp.GetInstance().FindIdByName(strPath)
		h, _ := windows.OpenProcess(uint32(win.PROCESS_ALL_ACCESS), false, dwPid)

		if h == windows.InvalidHandle || windows.OpenProcessToken(h, syscall.TOKEN_ALL_ACCESS, (*windows.Token)(&hToken)) != nil {
			return windows.InvalidHandle
		}
	}

	return hToken
}

func (se Session) GetSessionState(dwSessionId uint32) win.WTS_CONNECTSTATE_CLASS {
	pBuf, _, ok := win.WTSQuerySessionInformation(win.WTS_CURRENT_SERVER_HANDLE, dwSessionId, win.WTSConnectState)
	state := utils.IFSelector(ok, uint(*pBuf), uint(0)).(uint)
	return (win.WTS_CONNECTSTATE_CLASS)(state)
}
