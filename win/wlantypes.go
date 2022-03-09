package win

type DOT11_BSS_TYPE uint32

const (
	Dot11_BSS_type_infrastructure DOT11_BSS_TYPE = 1
	Dot11_BSS_type_independent    DOT11_BSS_TYPE = 2
	Dot11_BSS_type_any            DOT11_BSS_TYPE = 3
)

const DOT11_SSID_MAX_LENGTH = 32 // 32 bytes
type DOT11_SSID struct {
	// #ifndef __midl
	//     __range(0,32)
	// #endif
	SSIDLength uint32
	SSID       [DOT11_SSID_MAX_LENGTH]byte
}

// #ifdef __midl
// // use 4-byte enum
// typedef [v1_enum] enum _DOT11_PHY_TYPE {
// #else
// typedef enum _DOT11_PHY_TYPE {
// #endif
type DOT11_PHY_TYPE uint32

const (
	Dot11_phy_type_unknown    DOT11_PHY_TYPE = 0
	Dot11_phy_type_any        DOT11_PHY_TYPE = Dot11_phy_type_unknown
	Dot11_phy_type_fhss       DOT11_PHY_TYPE = 1
	Dot11_phy_type_dsss       DOT11_PHY_TYPE = 2
	Dot11_phy_type_irbaseband DOT11_PHY_TYPE = 3
	Dot11_phy_type_ofdm       DOT11_PHY_TYPE = 4
	Dot11_phy_type_hrdsss     DOT11_PHY_TYPE = 5
	Dot11_phy_type_erp        DOT11_PHY_TYPE = 6
	Dot11_phy_type_IHV_start  DOT11_PHY_TYPE = 0x80000000
	Dot11_phy_type_IHV_end    DOT11_PHY_TYPE = 0xffffffff
)

// // DOT11_AUTH_ALGO_LIST
// #ifdef __midl
// // use the 4-byte enum
// typedef [v1_enum] enum _DOT11_AUTH_ALGORITHM {
// #else
// typedef enum _DOT11_AUTH_ALGORITHM {
// #endif
type DOT11_AUTH_ALGORITHM uint32

const (
	DOT11_AUTH_ALGO_80211_OPEN       DOT11_AUTH_ALGORITHM = 1
	DOT11_AUTH_ALGO_80211_SHARED_KEY DOT11_AUTH_ALGORITHM = 2
	DOT11_AUTH_ALGO_WPA              DOT11_AUTH_ALGORITHM = 3
	DOT11_AUTH_ALGO_WPA_PSK          DOT11_AUTH_ALGORITHM = 4
	DOT11_AUTH_ALGO_WPA_NONE         DOT11_AUTH_ALGORITHM = 5 // used in NatSTA only
	DOT11_AUTH_ALGO_RSNA             DOT11_AUTH_ALGORITHM = 6
	DOT11_AUTH_ALGO_RSNA_PSK         DOT11_AUTH_ALGORITHM = 7
	DOT11_AUTH_ALGO_IHV_START        DOT11_AUTH_ALGORITHM = 0x80000000
	DOT11_AUTH_ALGO_IHV_END          DOT11_AUTH_ALGORITHM = 0xffffffff
)

// // Cipher algorithm Ids (for little endian platform)
// #ifdef __midl
// // use the 4-byte enum
// typedef [v1_enum] enum _DOT11_CIPHER_ALGORITHM {
// #else
// typedef enum _DOT11_CIPHER_ALGORITHM {
// #endif
type DOT11_CIPHER_ALGORITHM uint32

const (
	DOT11_CIPHER_ALGO_NONE          DOT11_CIPHER_ALGORITHM = 0x00
	DOT11_CIPHER_ALGO_WEP40         DOT11_CIPHER_ALGORITHM = 0x01
	DOT11_CIPHER_ALGO_TKIP          DOT11_CIPHER_ALGORITHM = 0x02
	DOT11_CIPHER_ALGO_CCMP          DOT11_CIPHER_ALGORITHM = 0x04
	DOT11_CIPHER_ALGO_WEP104        DOT11_CIPHER_ALGORITHM = 0x05
	DOT11_CIPHER_ALGO_WPA_USE_GROUP DOT11_CIPHER_ALGORITHM = 0x100
	DOT11_CIPHER_ALGO_RSN_USE_GROUP DOT11_CIPHER_ALGORITHM = 0x100
	DOT11_CIPHER_ALGO_WEP           DOT11_CIPHER_ALGORITHM = 0x101
	DOT11_CIPHER_ALGO_IHV_START     DOT11_CIPHER_ALGORITHM = 0x80000000
	DOT11_CIPHER_ALGO_IHV_END       DOT11_CIPHER_ALGORITHM = 0xffffffff
)
