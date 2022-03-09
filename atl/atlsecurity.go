package atl

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/James-Ye/go-frame/win"
	"golang.org/x/sys/windows"
)

type AccessToken struct {
	m_hToken   windows.Handle
	m_hProfile windows.Handle

	m_pRevert *Revert
}

type Revert struct {
}

func (at *AccessToken) GetProcessToken(desiredAccess uint32, hProcess windows.Handle /* = windows.InvalidHandle */) bool {
	if hProcess == windows.InvalidHandle {
		hProcess, _ = windows.GetCurrentProcess()
	}

	var hToken windows.Handle = windows.InvalidHandle
	if windows.OpenProcessToken(hProcess, desiredAccess, (*windows.Token)(&hToken)) != nil {
		return false
	}

	at.Clear()
	at.m_hToken = hToken
	return true
}

func (at *AccessToken) Clear() {
	if at.m_hProfile != windows.InvalidHandle {
		// ATLASSUME(cat.m_hToken)
		if at.m_hToken != windows.InvalidHandle {
			win.UnloadUserProfile(at.m_hToken, at.m_hProfile)
		}
		at.m_hProfile = windows.InvalidHandle
	}

	if at.m_hToken != windows.InvalidHandle {
		windows.CloseHandle(at.m_hToken)
		at.m_hToken = windows.InvalidHandle
	}

	// delete m_pRevert;
	at.m_pRevert = nil
}

func (at *AccessToken) GetTerminalServicesSessionId(pdwSessionId *uint32) bool {
	return at.GetInfo(pdwSessionId, win.TokenSessionId)
}

func (at *AccessToken) GetInfo(pRet interface{}, TokenClass win.TOKEN_INFORMATION_CLASS) bool {
	if pRet == nil {
		return false
	}

	var ReturnLength uint32
	if err := windows.GetTokenInformation(windows.Token(at.m_hToken), uint32(TokenClass), (*byte)(pRet.(unsafe.Pointer)), uint32(unsafe.Sizeof(*((*byte)(pRet.(unsafe.Pointer))))), &ReturnLength); err != nil {
		return false
	}
	return true
}

func (at *AccessToken) GetUser(pSid *GSID) bool {
	return at.GetInfoConvert(pSid, win.TokenUser, nil, 0)
}

func (at *AccessToken) GetInfoConvert(pRet interface{}, TokenClass win.TOKEN_INFORMATION_CLASS, pWork interface{} /*= nil*/, lengthWork uint32) bool {
	// ATLASSERT(pRet)
	if pRet == nil {
		return false
	}

	var ReturnLength uint32
	err := windows.GetTokenInformation(windows.Token(at.m_hToken), uint32(TokenClass), nil, 0, &ReturnLength)
	if err != windows.ERROR_INSUFFICIENT_BUFFER {
		return false
	}

	// USES_ATL_SAFE_ALLOCA;
	pWorkNew := []byte{}
	// pWork = static_cast<INFO_T *>(_ATL_SAFE_ALLOCA(dwLen, _ATL_SAFE_ALLOCA_DEF_THRESHOLD));
	if pWorkNew == nil {
		return false
	}
	if err := windows.GetTokenInformation(windows.Token(at.m_hToken), uint32(TokenClass), &pWorkNew[0], ReturnLength, &ReturnLength); err != nil {
		return false
	}

	at.InfoTypeToRetType(pRet, unsafe.Pointer(&pWorkNew[0]))
	pWork = &pWorkNew[0]
	return true
}

func (at *AccessToken) InfoTypeToRetType(pRet interface{}, pWork unsafe.Pointer) {
	typeStr := reflect.TypeOf(pRet).String()
	typeNames := strings.Split(typeStr, ".")
	typeName := typeNames[len(typeNames)-1]
	switch typeName {
	case "GSID":
		{
			pRet = *(*GSID)((*win.TOKEN_USER)(pWork).User.Sid)
		}
		break
	default:
		break
	}
}

func (at *AccessToken) Attach(hToken windows.Handle) {
	at.m_hToken = hToken
}

func (at *AccessToken) Detach() windows.Handle {
	hToken := at.m_hToken
	at.m_hToken = windows.InvalidHandle
	at.Clear()
	return hToken
}

type GSID struct {
	m_buffer [win.SECURITY_MAX_SID_SIZE]byte
	m_bValid bool // true if the CSid has been given a value

	m_sidnameuse     win.SID_NAME_USE
	m_strAccountName string
	m_strDomain      string
	m_strSid         string

	m_strSystem string
}

func (s *GSID) AccountName() string {
	if len(s.m_strAccountName) == 0 {
		s.GetAccountNameAndDomain()
	}
	return s.m_strAccountName
}

func (s *GSID) GetAccountNameAndDomain() {
	var cchName uint32 = 32
	var cbDomain uint32 = 32
	var Name []uint16
	var ReferencedDomainName []uint16

	pSystem, _ := windows.UTF16PtrFromString(s.m_strSystem)
	/* Prefast false warning: we do not use cbName or cbDomain as char buffers when call LookupAccountSid.*/
	if err := windows.LookupAccountSid(pSystem, (*windows.SID)(unsafe.Pointer(s._GetPSID())), &Name[0], &cchName, &ReferencedDomainName[0], &cbDomain, (*uint32)(&s.m_sidnameuse)); err == nil {
		s.m_strAccountName = windows.UTF16ToString(Name)
		s.m_strDomain = windows.UTF16ToString(ReferencedDomainName)
	} else {
		switch err {
		case windows.ERROR_INSUFFICIENT_BUFFER:
			{
				pszName, _ := windows.UTF16PtrFromString(s.m_strAccountName[0:cchName])
				pszDomain, _ := windows.UTF16PtrFromString(s.m_strDomain[0:cbDomain])

				if err := windows.LookupAccountSid(pSystem, (*windows.SID)(unsafe.Pointer(s._GetPSID())), pszName, &cchName, pszDomain, &cbDomain, (*uint32)(&s.m_sidnameuse)); err != nil {
					panic(err)
				}

				s.m_strAccountName = ""
				s.m_strDomain = ""

			}
			break
		case windows.ERROR_NONE_MAPPED:
			{
				s.m_strAccountName = ""
				s.m_strDomain = ""
				s.m_sidnameuse = win.SidTypeUnknown
			}
			break
		default:
			break
		}
	}
}

func (s *GSID) _GetPSID() *win.SID {
	var Sid win.SID
	Sid.Revision = s.m_buffer[0]
	Sid.SubAuthorityCount = s.m_buffer[1]
	copy(Sid.IdentifierAuthority.Value[:], s.m_buffer[2:8])
	copy((*[4]byte)(unsafe.Pointer(&Sid.SubAuthority[0]))[0:], s.m_buffer[8:12])
	return &Sid
}
