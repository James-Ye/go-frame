package win

type _NL_PREFIX_ORIGIN uint

const (
	//
	// These values are from iptypes.h.
	// They need to fit in a 4 bit field.
	//
	IpPrefixOriginOther _NL_PREFIX_ORIGIN = 0
	IpPrefixOriginManual
	IpPrefixOriginWellKnown
	IpPrefixOriginDhcp
	IpPrefixOriginRouterAdvertisement
	IpPrefixOriginUnchanged _NL_PREFIX_ORIGIN = 1 << 4
)

type NL_PREFIX_ORIGIN _NL_PREFIX_ORIGIN

type _NL_SUFFIX_ORIGIN uint

const (
	//
	// TODO: Remove the Nlso* definitions.
	//
	NlsoOther _NL_SUFFIX_ORIGIN = 0
	NlsoManual
	NlsoWellKnown
	NlsoDhcp
	NlsoLinkLayerAddress
	NlsoRandom

	//
	// These values are from in iptypes.h.
	// They need to fit in a 4 bit field.
	//
	IpSuffixOriginOther _NL_SUFFIX_ORIGIN = 0
	IpSuffixOriginManual
	IpSuffixOriginWellKnown
	IpSuffixOriginDhcp
	IpSuffixOriginLinkLayerAddress
	IpSuffixOriginRandom
	IpSuffixOriginUnchanged _NL_SUFFIX_ORIGIN = 1 << 4
)

type NL_SUFFIX_ORIGIN _NL_SUFFIX_ORIGIN

type _NL_DAD_STATE uint

const (
	//
	// TODO: Remove the Nlds* definitions.
	//
	NldsInvalid _NL_DAD_STATE = 0
	NldsTentative
	NldsDuplicate
	NldsDeprecated
	NldsPreferred

	//
	// These values are from in iptypes.h.
	//
	IpDadStateInvalid _NL_DAD_STATE = 0
	IpDadStateTentative
	IpDadStateDuplicate
	IpDadStateDeprecated
	IpDadStatePreferred
)

type NL_DAD_STATE _NL_DAD_STATE
