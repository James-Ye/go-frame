package win

//
// IPv4 Internet address
// This is an 'on-wire' format structure.
//
type S_UN_B struct {
	S_b1, S_b2, S_b3, S_b4 byte
}

type S_UN_W struct {
	S_w1, S_w2 uint16
}

type S_UN struct {
	S_un_b S_UN_B
	S_un_w S_UN_W
	S_addr uint32
}

type In_addr struct {
	S_un S_UN

	// const (
	// 	S_addr  = S_un.S_addr      /* can be used for most tcp & ip code */
	// 	S_host  = S_un.S_un_b.s_b2 // host on imp
	// 	S_net   = S_un.S_un_b.s_b1 // network
	// 	S_imp   = S_un.S_un_w.s_w2 // imp
	// 	S_impno = S_un.S_un_b.s_b4 // imp #
	// 	S_lh    = S_un.S_un_b.s_b3 // logical host
	// )

}

// func (ia *In_addr) In_addr(p *UChar) {
// 	ia.S_un.S_un_b.S_b1 = *p
// 	stl.PointerStepIn(unsafe.Pointer(p), unsafe.Sizeof(*p))
// 	ia.S_un.S_un_b.S_b2 = *p
// 	stl.PointerStepIn(unsafe.Pointer(p), unsafe.Sizeof(*p))
// 	ia.S_un.S_un_b.S_b3 = *p
// 	stl.PointerStepIn(unsafe.Pointer(p), unsafe.Sizeof(*p))
// 	ia.S_un.S_un_b.S_b4 = *p
// 	stl.PointerStepIn(unsafe.Pointer(p), unsafe.Sizeof(*p))
// 	var temp [8]UChar
// 	for i := 0; i < 8; i++ {
// 		temp[i] = *p
// 		stl.PointerStepIn(unsafe.Pointer(p), unsafe.Sizeof(*p))
// 	}
// 	wc := atl.CA2W(&temp[0], 4)
// 	ia.S_un.S_un_w.S_w1 = wc[0]
// 	ia.S_un.S_un_w.S_w2 = wc[1]
// 	ulc := atl.CA2ULong(&temp[4], 4)
// 	ia.S_un.S_addr = ulc[0]
// }

//
// IPv4 Socket address, Internet style
//

type Sockaddr_in struct {

	// #if(_WIN32_WINNT < 0x0600)
	// 	short   sin_family;
	// #else //(_WIN32_WINNT < 0x0600)
	sin_family uint16
	// #endif //(_WIN32_WINNT < 0x0600)

	sin_port uint16
	sin_addr In_addr
	sin_zero [8]byte
}
