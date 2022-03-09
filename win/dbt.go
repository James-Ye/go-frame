package win

import "golang.org/x/sys/windows"

const (
	DBT_DEVTYP_DEVICEINTERFACE = 0x00000005 // device interface class

	/*
	 * The following messages are for WM_DEVICECHANGE. The immediate list
	 * is for the wParam. ALL THESE MESSAGES PASS A POINTER TO A STRUCT
	 * STARTING WITH A DWORD SIZE AND HAVING NO POINTER IN THE STRUCT.
	 *
	 */
	DBT_DEVICEARRIVAL           = 0x8000 // system detected a new device
	DBT_DEVICEQUERYREMOVE       = 0x8001 // wants to remove, may fail
	DBT_DEVICEQUERYREMOVEFAILED = 0x8002 // removal aborted
	DBT_DEVICEREMOVEPENDING     = 0x8003 // about to remove, still avail.
	DBT_DEVICEREMOVECOMPLETE    = 0x8004 // device is gone
	DBT_DEVICETYPESPECIFIC      = 0x8005 // type specific event
)

type DEV_BROADCAST_DEVICEINTERFACE_W struct {
	Dbcc_size       uint32
	Dbcc_devicetype uint32
	Dbcc_reserved   uint32
	Dbcc_classguid  windows.GUID
	Dbcc_name       []uint16
}

type DEV_BROADCAST_DEVICEINTERFACE = DEV_BROADCAST_DEVICEINTERFACE_W
