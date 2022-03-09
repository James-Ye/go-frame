package win

// const (
// 	MAX_ADAPTER_DESCRIPTION_LENGTH = 128 // arb.
// 	MAX_ADAPTER_NAME_LENGTH        = 256 // arb.
// 	MAX_ADAPTER_ADDRESS_LENGTH     = 8   // arb.
// 	DEFAULT_MINIMUM_ENTITIES       = 32  // arb.
// 	MAX_HOSTNAME_LEN               = 128 // arb.
// 	MAX_DOMAIN_NAME_LEN            = 128 // arb.
// 	MAX_SCOPE_ID_LEN               = 256 // arb.
// 	MAX_DHCPV6_DUID_LENGTH         = 130 // RFC 3315.
// )

// //
// // IP_ADDRESS_STRING - store an IP address as a dotted decimal string
// //
// type IP_ADDRESS_STRING struct {
// 	String [4 * 4]rune
// }
// type PIP_ADDRESS_STRING *IP_ADDRESS_STRING
// type IP_MASK_STRING IP_ADDRESS_STRING
// type PIP_MASK_STRING *IP_ADDRESS_STRING

// //
// // IP_ADDR_STRING - store an IP address with its corresponding subnet mask,
// // both as dotted decimal strings
// //
// type _IP_ADDR_STRING struct {
// 	Next      *_IP_ADDR_STRING
// 	IpAddress IP_ADDRESS_STRING
// 	IpMask    IP_MASK_STRING
// 	Context   uint32
// }

// type IP_ADDR_STRING _IP_ADDR_STRING
// type PIP_ADDR_STRING *IP_ADDR_STRING

// type _IP_ADAPTER_INFO struct {
// 	Next                *_IP_ADAPTER_INFO
// 	ComboIndex          uint32
// 	AdapterName         [MAX_ADAPTER_NAME_LENGTH + 4]rune
// 	Description         [MAX_ADAPTER_DESCRIPTION_LENGTH + 4]rune
// 	AddressLength       uint
// 	Address             [MAX_ADAPTER_ADDRESS_LENGTH]byte
// 	Index               uint32
// 	Type                uint
// 	DhcpEnabled         uint
// 	CurrentIpAddress    PIP_ADDR_STRING
// 	IpAddressList       IP_ADDR_STRING
// 	GatewayList         IP_ADDR_STRING
// 	DhcpServer          IP_ADDR_STRING
// 	HaveWins            BOOL
// 	PrimaryWinsServer   IP_ADDR_STRING
// 	SecondaryWinsServer IP_ADDR_STRING
// 	LeaseObtained       time_t
// 	LeaseExpires        time_t
// }

// type IP_ADAPTER_INFO _IP_ADAPTER_INFO
// type PIP_ADAPTER_INFO *IP_ADAPTER_INFO

// //
// // The following types require Winsock2.
// //
// type IP_PREFIX_ORIGIN NL_PREFIX_ORIGIN
// type IP_SUFFIX_ORIGIN NL_SUFFIX_ORIGIN
// type IP_DAD_STATE NL_DAD_STATE

// type _IP_ADAPTER_UNICAST_ADDRESS_LH struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_UNICAST_ADDRESS_LH
// 	Address SOCKET_ADDRESS

// 	PrefixOrigin IP_PREFIX_ORIGIN
// 	SuffixOrigin IP_SUFFIX_ORIGIN
// 	DadState     IP_DAD_STATE

// 	ValidLifetime      uint64
// 	PreferredLifetime  uint64
// 	LeaseLifetime      uint64
// 	OnLinkPrefixLength uint8
// }
// type IP_ADAPTER_UNICAST_ADDRESS_LH _IP_ADAPTER_UNICAST_ADDRESS_LH
// type PIP_ADAPTER_UNICAST_ADDRESS_LH *_IP_ADAPTER_UNICAST_ADDRESS_LH

// type _IP_ADAPTER_ANYCAST_ADDRESS_XP struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_ANYCAST_ADDRESS_XP
// 	Address SOCKET_ADDRESS
// }
// type IP_ADAPTER_ANYCAST_ADDRESS_XP _IP_ADAPTER_ANYCAST_ADDRESS_XP
// type PIP_ADAPTER_ANYCAST_ADDRESS_XP *_IP_ADAPTER_ANYCAST_ADDRESS_XP

// type _IP_ADAPTER_MULTICAST_ADDRESS_XP struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_MULTICAST_ADDRESS_XP
// 	Address SOCKET_ADDRESS
// }
// type IP_ADAPTER_MULTICAST_ADDRESS_XP _IP_ADAPTER_MULTICAST_ADDRESS_XP
// type PIP_ADAPTER_MULTICAST_ADDRESS_XP *_IP_ADAPTER_MULTICAST_ADDRESS_XP

// type _IP_ADAPTER_DNS_SERVER_ADDRESS_XP struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_DNS_SERVER_ADDRESS_XP
// 	Address SOCKET_ADDRESS
// }
// type IP_ADAPTER_DNS_SERVER_ADDRESS_XP _IP_ADAPTER_DNS_SERVER_ADDRESS_XP
// type PIP_ADAPTER_DNS_SERVER_ADDRESS_XP *_IP_ADAPTER_DNS_SERVER_ADDRESS_XP

// type _IP_ADAPTER_PREFIX_XP struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next         *_IP_ADAPTER_PREFIX_XP
// 	Address      SOCKET_ADDRESS
// 	PrefixLength uint64
// }
// type IP_ADAPTER_PREFIX_XP _IP_ADAPTER_PREFIX_XP
// type PIP_ADAPTER_PREFIX_XP *_IP_ADAPTER_PREFIX_XP

// type _IP_ADAPTER_WINS_SERVER_ADDRESS_LH struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_WINS_SERVER_ADDRESS_LH
// 	Address SOCKET_ADDRESS
// }
// type IP_ADAPTER_WINS_SERVER_ADDRESS_LH _IP_ADAPTER_WINS_SERVER_ADDRESS_LH
// type PIP_ADAPTER_WINS_SERVER_ADDRESS_LH *_IP_ADAPTER_WINS_SERVER_ADDRESS_LH

// type _IP_ADAPTER_GATEWAY_ADDRESS_LH struct {
// 	Alignment uint64
// 	Length    uint64
// 	Flags     uint32

// 	Next    *_IP_ADAPTER_GATEWAY_ADDRESS_LH
// 	Address SOCKET_ADDRESS
// }
// type IP_ADAPTER_GATEWAY_ADDRESS_LH _IP_ADAPTER_GATEWAY_ADDRESS_LH
// type PIP_ADAPTER_GATEWAY_ADDRESS_LH *_IP_ADAPTER_GATEWAY_ADDRESS_LH
