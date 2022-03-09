package win

import "golang.org/x/sys/windows"

//
// Types of tunnels (sub-type of IF_TYPE when IF_TYPE is IF_TYPE_TUNNEL).
// See http://www.iana.org/assignments/ianaiftype-mib.
//
type _TUNNEL_TYPE uint

const (
	TUNNEL_TYPE_NONE   _TUNNEL_TYPE = 0
	TUNNEL_TYPE_OTHER  _TUNNEL_TYPE = 1
	TUNNEL_TYPE_DIRECT _TUNNEL_TYPE = 2
	TUNNEL_TYPE_6TO4   _TUNNEL_TYPE = 11
	TUNNEL_TYPE_ISATAP _TUNNEL_TYPE = 13
	TUNNEL_TYPE_TEREDO _TUNNEL_TYPE = 14
)

type TUNNEL_TYPE _TUNNEL_TYPE
type PTUNNEL_TYPE *_TUNNEL_TYPE

type _INFO struct {
	Reserved     uint64 //:24;
	NetLuidIndex uint64 //:24;
	IfType       uint64 //:16;                  // equal to IANA IF type
}

type _NET_LUID_LH struct {
	Value uint64
	Info  _INFO
}

type NET_LUID_LH _NET_LUID_LH
type PNET_LUID_LH *_NET_LUID_LH

//
// Need to make this visible on all platforms (for the purpose of IF_LUID).
//
type NET_LUID NET_LUID_LH
type PNET_LUID *NET_LUID

//
// IF_LUID
//
// Define the locally unique datalink interface identifier type.
// This type is persistable.
//
type IF_LUID NET_LUID
type PIF_LUID *NET_LUID

// Interface Index (ifIndex)
type NET_IFINDEX uint64
type PNET_IFINDEX *uint64

// Interface Type (IANA ifType)
type NET_IFTYPE uint16
type PNET_IFTYPE *uint16

//
// IF_INDEX
//
// Define the interface index type.
// This type is not persistable.
// This must be unsigned (not an enum) to replace previous uses of
// an index that used a uint32 type.
//

type IF_INDEX NET_IFINDEX
type PIF_INDEX *NET_IFINDEX

//
// OperStatus values from RFC 2863
//
type if_oper_status uint

const (
	IfOperStatusUp if_oper_status = 1
	IfOperStatusDown
	IfOperStatusTesting
	IfOperStatusUnknown
	IfOperStatusDormant
	IfOperStatusNotPresent
	IfOperStatusLowerLayerDown
)

type IF_OPER_STATUS if_oper_status

type NET_IF_COMPARTMENT_ID uint32
type PNET_IF_COMPARTMENT_ID *uint32

//
// Define compartment ID type:
//
const (
	NET_IF_COMPARTMENT_ID_UNSPECIFIED NET_IF_COMPARTMENT_ID = 0
	NET_IF_COMPARTMENT_ID_PRIMARY     NET_IF_COMPARTMENT_ID = 1
)

//
// Define NetworkGUID type:
//
type NET_IF_NETWORK_GUID windows.GUID
type PNET_IF_NETWORK_GUID *windows.GUID

type _NET_IF_CONNECTION_TYPE uint

const (
	NET_IF_CONNECTION_DEDICATED _NET_IF_CONNECTION_TYPE = 1
	NET_IF_CONNECTION_PASSIVE
	NET_IF_CONNECTION_DEMAND
	NET_IF_CONNECTION_MAXIMUM
)

type NET_IF_CONNECTION_TYPE _NET_IF_CONNECTION_TYPE
type PNET_IF_CONNECTION_TYPE *_NET_IF_CONNECTION_TYPE
