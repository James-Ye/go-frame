package win

type IFTYPE uint64

//////////////////////////////////////////////////////////////////////////////
//                                                                          //
// The following are the the operational states for WAN and LAN interfaces. //
// The order of the states seems weird, but is done for a purpose. All      //
// states >= CONNECTED can transmit data right away. States >= DISCONNECTED //
// can tx data but some set up might be needed. States < DISCONNECTED can   //
// not transmit data.                                                       //
// A card is marked UNREACHABLE if DIM calls InterfaceUnreachable for       //
// reasons other than failure to connect.                                   //
//                                                                          //
// NON_OPERATIONAL -- Valid for LAN Interfaces. Means the card is not       //
//                      working or not plugged in or has no address.        //
// UNREACHABLE     -- Valid for WAN Interfaces. Means the remote site is    //
//                      not reachable at this time.                         //
// DISCONNECTED    -- Valid for WAN Interfaces. Means the remote site is    //
//                      not connected at this time.                         //
// CONNECTING      -- Valid for WAN Interfaces. Means a connection attempt  //
//                      has been initiated to the remote site.              //
// CONNECTED       -- Valid for WAN Interfaces. Means the remote site is    //
//                      connected.                                          //
// OPERATIONAL     -- Valid for LAN Interfaces. Means the card is plugged   //
//                      in and working.                                     //
//                                                                          //
// It is the users duty to convert these values to MIB-II values if they    //
// are to be used by a subagent                                             //
//                                                                          //
//////////////////////////////////////////////////////////////////////////////

type (
	_INTERNAL_IF_OPER_STATUS uint32
	INTERNAL_IF_OPER_STATUS  _INTERNAL_IF_OPER_STATUS
)

var (
	IF_OPER_STATUS_NON_OPERATIONAL _INTERNAL_IF_OPER_STATUS = 0
	IF_OPER_STATUS_UNREACHABLE     _INTERNAL_IF_OPER_STATUS = 1
	IF_OPER_STATUS_DISCONNECTED    _INTERNAL_IF_OPER_STATUS = 2
	IF_OPER_STATUS_CONNECTING      _INTERNAL_IF_OPER_STATUS = 3
	IF_OPER_STATUS_CONNECTED       _INTERNAL_IF_OPER_STATUS = 4
	IF_OPER_STATUS_OPERATIONAL     _INTERNAL_IF_OPER_STATUS = 5
)
