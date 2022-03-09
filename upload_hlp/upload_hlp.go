package upload_hlp

import (
	"unsafe"

	"github.com/James-Ye/go-frame/logger"
)

type MemoryStruct struct {
	memory *rune
	size   uint
}

type Upload_Interface interface {
	CreateUpload(bHttps bool)
	DestoryUpload()
	Run() //线程入口
	/* 配置函数，URL、上传文件名字，回显信息写入到指定文件中 */
	Start(data *rune, url string) bool //启动上传
	start_base64(data string, url string) bool
	Get(url string) bool //启动上传
	Stop() bool          //停止上传
	isRunning() bool
	onUpload() bool
	get_return_info(info string)
	Get_return(info string)
	setOption(ispost bool) bool
	clear()

	// 新的写回信息
	WriteCallback(contents uintptr, size uint, nmemb uint, userp uintptr) uint
}

type Upload struct {
	M_timeout        int32
	M_connecttimeout int32

	m_pCURL       unsafe.Pointer //*CURL //libcurl
	m_pFromPost   unsafe.Pointer //*curl_httppost
	m_pLastElem   unsafe.Pointer //*curl_httppost
	m_pOptionList unsafe.Pointer //*curl_slist

	upChunk          MemoryStruct //存放上传信息的缓存
	chunk            MemoryStruct //存放回显信息的缓存
	m_bRunning       bool         //运行状态标志
	m_strRemoteURL   string       //URL字符串
	info_ret         string
	m_nLocalFileSize uint //本地文件大小
	m_strpost        string
	m_bHttps         bool
}

func CreateUpload(bHttps bool) *Upload {
	return new(Upload)
}

func (u *Upload) Run() bool {
	logger.Trace("run begin")
	ret := u.onUpload()
	logger.Trace("run end")
	return ret
}

func (u *Upload) start_base64(data string, url string) bool {
	return false
	//logger.Trace("start_base64 begin")
	//m_strRemoteURL = url;
	//// base64 encode.
	//data = base64_encode((const unsigned char*)data.c_str(),data.length());

	//if(m_bRunning==true)
	//	return true;

	//m_pCURL=curl_easy_init();
	//if(m_pCURL==NULL)
	//{
	//	logger.Trace("m_pCURL==NULL");
	//	return false;
	//}
	//setOption(true);
	//m_bRunning=true;
	//logger.Trace("start_base64 end")
	//return true;
}

func (u *Upload) Start(data string, url string) bool {
	logger.Trace("start begin")
	u.m_strRemoteURL = url
	u.m_strpost = data

	if u.m_bRunning {
		return true
	}

	// u.m_pCURL = curl_easy_init()
	if u.m_pCURL == nil {
		logger.Trace("m_pCURL==NULL")
		return false
	}
	u.setOption(true)
	u.m_bRunning = true
	logger.Trace("start end")
	return true
}

func (u *Upload) Get(url string) bool {
	logger.Trace("get begin")
	u.m_strRemoteURL = url

	if u.m_bRunning {
		return true
	}

	// u.m_pCURL = curl_easy_init()
	if u.m_pCURL == nil {
		logger.Trace("m_pCURL==NULL")
		return false
	}
	u.setOption(false)
	u.m_bRunning = true
	logger.Trace("get end")
	return true
}

func (u *Upload) Stop() bool {
	logger.Trace("stop begin!")
	u.clear()
	u.m_bRunning = false
	logger.Trace("stop end!")
	return true
}

func (u *Upload) isRunning() bool {
	return u.m_bRunning
}

func (u *Upload) setOption(ispost bool) bool {
	logger.Trace("setOption begin")
	// u.m_pOptionList = curl_slist_append(u.m_pOptionList, "Expect:")
	// curl_easy_setopt(u.m_pCURL, CURLOPT_HTTPHEADER, u.m_pOptionList)
	if ispost {
		// 	// curl_easy_setopt(u.m_pCURL, CURLOPT_POSTFIELDS, u.m_strpost.c_str())
	}
	//链接
	// curl_easy_setopt(u.m_pCURL, CURLOPT_URL, u.m_strRemoteURL.c_str())
	// //超时
	// curl_easy_setopt(u.m_pCURL, CURLOPT_CONNECTTIMEOUT_MS, u.m_connecttimeout)
	// curl_easy_setopt(u.m_pCURL, CURLOPT_TIMEOUT_MS, m_timeout)
	// //回调
	// curl_easy_setopt(u.m_pCURL, CURLOPT_WRITEFUNCTION, u.WriteCallback)
	// curl_easy_setopt(u.m_pCURL, CURLOPT_WRITEDATA, uintptr(&chunk))
	//证书
	if u.m_bHttps {
		// curl_easy_setopt(u.m_pCURL, CURLOPT_SSL_VERIFYPEER, false)
		// curl_easy_setopt(u.m_pCURL, CURLOPT_SSL_VERIFYHOST, false)
		//curl_easy_setopt(m_pCURL,CURLOPT_CAINFO,pathmanager.Get360Path("EntClient\\nacservice\\ca.crt"))
	}
	logger.Trace("setOption end")
	return true
}

func (u *Upload) clear() {
	logger.Trace("clear begin")
	if u.m_pCURL != nil {
		// curl_easy_cleanup(u.m_pCURL)
		u.m_pCURL = nil
		// curl_global_cleanup()
	}
	//m_strRemoteURL.clear();
	logger.Trace("clear end")
}

func (u *Upload) onUpload() bool {
	logger.Trace("onUpload begin")
	// response_code := 0
	// var return_code CURLcode = CURLE_OK
	// return_code = curl_easy_perform(u.m_pCURL)

	// if return_code != CURLE_OK {
	// 	logger.Trace("onUpload error:%d", return_code)
	// 	return 0
	// }
	// curl_easy_getinfo(u.m_pCURL, CURLINFO_RESPONSE_CODE, &response_code)
	logger.Trace("onUpload end")
	return true
}

func (u *Upload) WriteCallback(contents uintptr, size uint, nmemb uint, userp uintptr) uint {
	//uint32 dwResult;
	realsize := size * nmemb
	// wContents := (*rune)(contents)
	// mem := (*MemoryStruct)(userp)
	// mem.memory=(*rune)realloc(mem->memory,mem->size+realsize+1);
	// if mem.memory == nil {
	// 	return 0
	// }
	// memory.memcpy(&(mem.memory[mem.size]), contents, realsize)
	// mem.size += realsize
	// mem.memory[mem.size] = 0
	return realsize
}

func (u *Upload) get_return_info(info string) {
	// info.assign(chunk.memory, chunk.size)
	// info = base64_decode(info)
}

func (u *Upload) Get_return(info string) {
	// info.assign(chunk.memory, chunk.size)
}
