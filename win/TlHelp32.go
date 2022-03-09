package win

const (
	//
	// The th32ProcessID argument is only used if TH32CS_SNAPHEAPLIST or
	// TH32CS_SNAPMODULE is specified. th32ProcessID == 0 means the current
	// process.
	//
	// NOTE that all of the snapshots are global except for the heap and module
	//      lists which are process specific. To enumerate the heap or module
	//      state for all WIN32 processes call with TH32CS_SNAPALL and the
	//      current process. Then for each process in the TH32CS_SNAPPROCESS
	//      list that isn't the current process, do a call with just
	//      TH32CS_SNAPHEAPLIST and/or TH32CS_SNAPMODULE.
	//
	// dwFlags
	//
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_SNAPALL      = (TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE)
	TH32CS_INHERIT      = 0x80000000
)

/***** Process walking *************************************************/

type tagPROCESSENTRY32W struct {
	DwSize              uint32
	cntUsage            uint32
	Th32ProcessID       uint32 // this process
	th32DefaultHeapID   uint32
	th32ModuleID        uint32 // associated exe
	cntThreads          uint32
	th32ParentProcessID uint32 // this process's parent process
	pcPriClassBase      int32  // Base priority of process's threads
	dwFlags             uint32
	SzExeFile           []uint16 // Path
}
type (
	PROCESSENTRY32W   tagPROCESSENTRY32W
	PPROCESSENTRY32W  *PROCESSENTRY32W
	LPPROCESSENTRY32W *PROCESSENTRY32W
	PROCESSENTRY32    PROCESSENTRY32W
	PPROCESSENTRY32   *PROCESSENTRY32
	LPPROCESSENTRY32  *PROCESSENTRY32
)
