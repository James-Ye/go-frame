package win

// //
// // Address families.
// //

// type ADDRESS_FAMILY uint16

// //
// // Although AF_UNSPEC is defined for backwards compatibility, using
// // AF_UNSPEC for the "af" parameter when creating a socket is STRONGLY
// // DISCOURAGED.  The interpretation of the "protocol" parameter
// // depends on the actual address family chosen.  As environments grow
// // to include more and more address families that use overlapping
// // protocol values there is more and more chance of choosing an
// // undesired address family when AF_UNSPEC is used.
// //
// const (
// 	AF_UNSPEC    ADDRESS_FAMILY = 0      // unspecified
// 	AF_UNIX      ADDRESS_FAMILY = 1      // local to host (pipes, portals)
// 	AF_INET      ADDRESS_FAMILY = 2      // internetwork: UDP, TCP, etc.
// 	AF_IMPLINK   ADDRESS_FAMILY = 3      // arpanet imp addresses
// 	AF_PUP       ADDRESS_FAMILY = 4      // pup protocols: e.g. BSP
// 	AF_CHAOS     ADDRESS_FAMILY = 5      // mit CHAOS protocols
// 	AF_NS        ADDRESS_FAMILY = 6      // XEROX NS protocols
// 	AF_IPX       ADDRESS_FAMILY = AF_NS  // IPX protocols: IPX, SPX, etc.
// 	AF_ISO       ADDRESS_FAMILY = 7      // ISO protocols
// 	AF_OSI       ADDRESS_FAMILY = AF_ISO // OSI is ISO
// 	AF_ECMA      ADDRESS_FAMILY = 8      // european computer manufacturers
// 	AF_DATAKIT   ADDRESS_FAMILY = 9      // datakit protocols
// 	AF_CCITT     ADDRESS_FAMILY = 10     // CCITT protocols, X.25 etc
// 	AF_SNA       ADDRESS_FAMILY = 11     // IBM SNA
// 	AF_DECnet    ADDRESS_FAMILY = 12     // DECnet
// 	AF_DLI       ADDRESS_FAMILY = 13     // Direct data link interface
// 	AF_LAT       ADDRESS_FAMILY = 14     // LAT
// 	AF_HYLINK    ADDRESS_FAMILY = 15     // NSC Hyperchannel
// 	AF_APPLETALK ADDRESS_FAMILY = 16     // AppleTalk
// 	AF_NETBIOS   ADDRESS_FAMILY = 17     // NetBios-style addresses
// 	AF_VOICEVIEW ADDRESS_FAMILY = 18     // VoiceView
// 	AF_FIREFOX   ADDRESS_FAMILY = 19     // Protocols from Firefox
// 	AF_UNKNOWN1  ADDRESS_FAMILY = 20     // Somebody is using this!
// 	AF_BAN       ADDRESS_FAMILY = 21     // Banyan
// 	AF_ATM       ADDRESS_FAMILY = 22     // Native ATM Services
// 	AF_INET6     ADDRESS_FAMILY = 23     // Internetwork Version 6
// 	AF_CLUSTER   ADDRESS_FAMILY = 24     // Microsoft Wolfpack
// 	AF_12844     ADDRESS_FAMILY = 25     // IEEE 1284.4 WG AF
// 	AF_IRDA      ADDRESS_FAMILY = 26     // IrDA
// 	AF_NETDES    ADDRESS_FAMILY = 28     // Network Designers OSI & gateway
// )

// //
// // Structure used to store most addresses.
// //
// type sockaddr struct {

// 	// #if (_WIN32_WINNT < 0x0600)
// 	// 	u_short sa_family;
// 	// #else
// 	Sa_family ADDRESS_FAMILY // Address family.
// 	// #endif //(_WIN32_WINNT < 0x0600)

// 	Sa_data [14]byte // Up to 14 bytes of direct address.
// }
// type SOCKADDR sockaddr
// type PSOCKADDR *sockaddr
// type LPSOCKADDR *sockaddr

// /*
//  * SockAddr Information
//  */
// type _SOCKET_ADDRESS struct {
// 	LpSockaddr      LPSOCKADDR
// 	ISockaddrLength int
// }
// type SOCKET_ADDRESS _SOCKET_ADDRESS
// type PSOCKET_ADDRESS *_SOCKET_ADDRESS
// type LPSOCKET_ADDRESS *_SOCKET_ADDRESS

const (
	INADDR_NONE uint32 = 0xffffffff
)
