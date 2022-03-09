package process_hlp

import (
	"strings"
	"sync"
	"syscall"
	"unsafe"

	"github.com/James-Ye/go-frame/atl"
	"github.com/James-Ye/go-frame/logger"
	"github.com/James-Ye/go-frame/win"
	"golang.org/x/sys/windows"
)

var g_ProcessHelper *ProcessHelper
var once sync.Once

func GetInstance() *ProcessHelper {
	once.Do(func() {
		g_ProcessHelper = &ProcessHelper{}
	})
	return g_ProcessHelper
}

type ProcessHelper struct {
}

func (ph *ProcessHelper) CreateAsUser(TokenHandle windows.Token, filename string, strParameters string, nShow int) windows.Handle {
	strCommandLine := filename

	if strCommandLine[0:1] != "\"" {
		win.PathQuoteSpaces(&strCommandLine)
	}

	if len(strParameters) > 0 && strParameters[:1] != "" {
		strCommandLine += " "
		strCommandLine += strParameters
	}

	var creationFlags uint32 = win.NORMAL_PRIORITY_CLASS | win.CREATE_NEW_CONSOLE
	lpEnvironmentBlock, ok := win.CreateEnvironmentBlock(windows.Handle(TokenHandle), false)
	if ok {
		creationFlags |= win.CREATE_UNICODE_ENVIRONMENT
	}

	var si windows.StartupInfo
	var pi windows.ProcessInformation

	si.Cb = uint32(unsafe.Sizeof(si))
	si.Desktop, _ = syscall.UTF16PtrFromString("winsta0\\default")
	si.Flags = win.STARTF_USESHOWWINDOW
	si.ShowWindow = uint16(nShow)

	pCommandLine, _ := syscall.UTF16PtrFromString(strCommandLine)

	if err := windows.CreateProcessAsUser(TokenHandle, nil, pCommandLine, nil, nil, false, creationFlags, (*uint16)(lpEnvironmentBlock), nil, &si, &pi); err != nil {
		logger.Error("in CreateAsUser, CreateProcess Failed = %d", err)
	}

	if lpEnvironmentBlock != nil {
		win.DestroyEnvironmentBlock(lpEnvironmentBlock)
	}

	if pi.Thread == 0 {
		windows.CloseHandle(pi.Thread)
	}

	return pi.Thread
}

//******************************************************************************
func (ph *ProcessHelper) CreateAsSession(SessionId uint32, strFile string, strParameters string /* = ""*/, nShow int /* = 5*/) windows.Handle {
	logger.Trace("CreateAsSession = %d", SessionId)

	var TokenHandle windows.Handle

	if ok, err := win.WTSQueryUserToken(SessionId, TokenHandle); !ok {
		logger.Error("QueryUserToken Failed = %d", err)

		TokenHandle = ph.QuerySessionToken(SessionId)

		if TokenHandle != windows.InvalidHandle {
			logger.Error("QuerySessionToken Failed ")
			return windows.InvalidHandle
		}
	} else {
		logger.Trace("WTSQueryUserToken success")
		var Token atl.AccessToken

		Token.Attach(TokenHandle)

		var Sid atl.GSID

		if Token.GetUser(&Sid) {
			strAccountName := Sid.AccountName()

			if strings.Compare(strAccountName, "SYSTEM") == 0 {
				logger.Trace("in CreateAsSession, WTSQueryUserToken AccountName is SYSTEM")
				return windows.InvalidHandle
			}
		}

		Token.Detach()
	}

	ProcessHandle := ph.CreateAsUser(windows.Token(TokenHandle), strFile, strParameters, nShow)

	windows.CloseHandle(TokenHandle)

	return ProcessHandle
}

//******************************************************************************
func (ph *ProcessHelper) QuerySessionToken(SessionId uint32) windows.Handle {
	var TokenHandle windows.Handle = 0

	var Entry syscall.ProcessEntry32
	Entry.Size = uint32(unsafe.Sizeof(Entry))

	SnapshotHandle, err := syscall.CreateToolhelp32Snapshot(win.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		logger.Trace("QuerySessionToken::, CreateToolhelp32Snapshot failed, err = %d", err)
		return windows.InvalidHandle
	}

	for err := syscall.Process32First(SnapshotHandle, &Entry); err == nil; err = syscall.Process32Next(SnapshotHandle, &Entry) {
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

			if ProcessToken.GetTerminalServicesSessionId(&dwSessionId) && dwSessionId == SessionId {
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

	syscall.CloseHandle(SnapshotHandle)

	return TokenHandle
}

//******************************************************************************
func (ph *ProcessHelper) FindIdByName(szProcName string) uint32 {
	hSP, _ := windows.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if hSP != windows.InvalidHandle {
		bFound := false
		var pe windows.ProcessEntry32
		pe.Size = uint32(unsafe.Sizeof(pe))
		for err := windows.Process32First(hSP, &pe); err == nil; err = windows.Process32Next(hSP, &pe) {
			if strings.Compare(szProcName, syscall.UTF16ToString(pe.ExeFile[:])) == 0 {
				bFound = true
				break
			}
		}

		windows.CloseHandle(hSP)

		if bFound {
			return pe.ProcessID
		}
	}

	return 0
}
