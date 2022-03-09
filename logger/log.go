package logger

// import (
// 	"github.com/lxn/win"
// )
// //日志等级
// const (
// 	LOG_DEBUG	= 0	//调试信息
// 	LOG_TRACE	= 1	//调试信息
// 	LOG_WARN	= 2	//警告
// 	LOG_ERROR	= 3	//错误
// 	LOG_NONE	= 4		//不记录日志
// )

// //日志操作类
// type Log struct {

// }

// //创建日志对象
// //hModule	为NULL时使用调用进程模块
// //szLogPath	为NULL时使用模块路径

// func (log *Log) Log(hModule unsafe.Pointer, szLogPath string) {
// 		Init(hModule, szLogPath)
// }

// func knownFolderPath(id win.CSIDL) (string, error) {
// 	var buf [win.MAX_PATH]uint16

// 	if !win.SHGetSpecialFolderPath(0, &buf[0], id, false) {
// 		return "", newError("SHGetSpecialFolderPath failed")
// 	}

// 	return syscall.UTF16ToString(buf[0:]), nil
// }

// //初始化日志设置
// //hModule	为NULL时使用调用进程模块
// //szLogPath	为NULL时使用模块路径-->>临时路径
// func Init(hModule unsafe.Pointer, szLogPath string) {
// 	//如果模块名为NULL，设置为主模块名称
// 	CString strModuleName = Path::GetModuleName(hModule);
// 	//CString strLogPath = Helper::IsNullOrEmpty(szLogPath)? Path::GetModulePath(hModule) : szLogPath;

// 		WCHAR szAPPDataPath[MAX_PATH] = {0};
// 		WCHAR szLogName[MAX_PATH + 1] = {0};

// 		var buf [win.MAX_PATH]uint16
// 		if win.SHGetSpecialFolderPath(0, &buf[0], win.CSIDL_COMMON_APPDATA, true) {
// 			return "", newError("SHGetSpecialFolderPath failed")
// 		}
// 		strAPPDataPath := knownFolderPath(win.CSIDL_COMMON_APPDATA)
// 		windows.GetModuleFileName()

// 		::PathCombine(szLogName, szAPPDataPath, _T("360Skylar6"));
// 		CString strLogPath = Helper::IsNullOrEmpty(szLogPath)? szLogName : szLogPath;

// 		Path::CreateDirectory(strLogPath);

// 		m_strFileName = Path::Combine(strLogPath, strModuleName + _T(".log"));
// 		m_eLogLevel = LOG_TRACE;
// 		m_nMaxSize = 5 * 1024 * 1024;
// 	}

// 	//write log with error level ERROR
// 	void Error(LPCSTR szFormat, ...)
// 	{
// 		va_list args;
// 		va_start(args, szFormat);
// 		PrintV(LOG_ERROR, szFormat, args);
// 		va_end(args);
// 	}

// 	//write log with error level WARN
// 	void Warn(LPCSTR szFormat, ...)
// 	{
// 		va_list args;
// 		va_start(args, szFormat);
// 		PrintV(LOG_WARN, szFormat, args);
// 		va_end(args);
// 	}

// 	//write log with error level DEBUG
// 	void Debug(LPCSTR szFormat, ...)
// 	{
// 		va_list args;
// 		va_start(args, szFormat);
// 		PrintV(LOG_DEBUG, szFormat, args);
// 		va_end(args);
// 	}

// 	void Print(LOGLEVEL eLevel, LPCSTR szFormat, ...)
// 	{
// 		va_list args;
// 		va_start(args, szFormat);
// 		PrintV(eLevel, szFormat, args);
// 		va_end(args);
// 	}

// 	void PrintV(LOGLEVEL eLevel, LPCSTR szFormat, va_list args)
// 	{
// 		if(eLevel < m_eLogLevel)
// 		{
// 			return;
// 		}

// 		ATLASSERT(!m_strFileName.IsEmpty());
// 		//未设置文件名，直接退出
// 		if(m_strFileName.IsEmpty())
// 		{
// 			return;
// 		}

// 		CStringA strMsg;
// 		strMsg.FormatV(szFormat, args);
// 		// format message
// 		COleDateTime dt(COleDateTime::GetCurrentTime());
// 		CStringA str;
// 		str.Format(
// 			"%s (pid:%d, tid:%d) %s->%s\r\n",
// 			(LPCSTR)CT2A(dt.Format(_T("%Y-%m-%d %H:%M:%S"))),
// 			syscall.GetCurrentProcessId(),
// 			GetCurrentThreadId(),
// 			ToString(eLevel),
// 			(LPCSTR)strMsg);

// 		// write message
// 		InternalPrint(str);
// 	}

// 	void SetFileName(LPCTSTR szFileName){ m_strFileName = szFileName; }
// 	CString GetFileName() const { return m_strFileName; }
// 	void SetLevel(LOGLEVEL eLogLevel){ m_eLogLevel = eLogLevel; }
// 	LOGLEVEL GetLevel() const { return m_eLogLevel; }
// 	//nMaxSize == 0 不限制日志大小，单位：字节
// 	void SetMaxSize(size_t nMaxSize = 512 * 1024){ m_nMaxSize = nMaxSize; }
// 	size_t GetMaxSize() const { return m_nMaxSize; }

// private:
// 	LOGLEVEL m_eLogLevel;	//日志记录等级
// 	CString m_strFileName;	//日志文件名
// 	size_t m_nMaxSize;		//尺寸限制（单位：字节），日志文件大于此尺寸时，自动删除，0表示不限制

// private:
// 	void InternalPrint(CStringA strMsg)
// 	{
// 		FILE* fp = _tfsopen(m_strFileName, _T("ab"), _SH_DENYWR);
// 		if(fp)
// 		{
// 			// check if the file exceed the max size, if so, truncate it
// 			if(m_nMaxSize != 0)
// 			{
// 				int fd = _fileno(fp);
// 				if(_filelengthi64(fd) > (__int64)m_nMaxSize)
// 				{
// 					fclose(fp);
// 					CString strFileNameBak = m_strFileName + _T(".bak");
// 					CopyFile((LPCTSTR)m_strFileName, (LPCTSTR)strFileNameBak, FALSE);
// 					fp = _tfsopen(m_strFileName, _T("ab"), _SH_DENYWR);
// 					if(!fp)
// 						return;
// 					fd = _fileno(fp);
// 					_chsize_s(fd, 0);
// 				}
// 			}
// 			fwrite((LPCSTR)strMsg, 1, strMsg.GetLength(), fp);

// 			fclose(fp);
// 		}
// 	}

// 	LPCSTR ToString(LOGLEVEL eLevel)
// 	{
// 		if(eLevel < LOG_DEBUG || eLevel > LOG_ERROR)
// 		{
// 			return "UNKNOWN";
// 		}

// 		static const LPCSTR g_LevelString[] =
// 		{
// 			"DEBUG",//调试信息
// 			"TRACE",
// 			"WARN",	//警告
// 			"ERROR"	//错误
// 		};

// 		return g_LevelString[eLevel];
// 	}
// };
// }
