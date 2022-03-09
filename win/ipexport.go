package win

const MAX_ADAPTER_NAME = 128

type IP_ADAPTER_INDEX_MAP struct {
	Index uint32
	Name  [MAX_ADAPTER_NAME]uint16
}

type IP_INTERFACE_INFO struct {
	NumAdapters int32
	Adapter     [1]IP_ADAPTER_INDEX_MAP
}
