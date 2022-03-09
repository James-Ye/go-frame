package stl

import (
	"fmt"
	"unsafe"
)

func CA2W(p *uint8, size int) []uint16 {
	var wchars []uint16
	if p == nil {
		return wchars
	}

	end := unsafe.Pointer(p)
	step := unsafe.Sizeof(*p)
	step2 := unsafe.Sizeof(wchars[0])

	n := 0
	offset := 8 * (step2 - 1)

	for *((*uint8)(end)) != 0 && n < size/2 {
		wchars[n] |= uint16(*((*uint8)(end))) << offset
		end = PointerStepIn(end, step)
		if offset == 0 {
			offset = 8 * step2
			n++
		} else {
			offset -= 8 * step
		}
	}

	return wchars
}

func CA2ULong(p *uint8, size int) []uint32 {
	var ulchars []uint32
	if p == nil {
		return ulchars
	}

	end := unsafe.Pointer(p)
	step := unsafe.Sizeof(*p)
	step2 := unsafe.Sizeof(ulchars[0])

	n := 0
	offset := 8 * (step2 - 1)

	for *((*uint8)(end)) != 0 && n < size/2 {
		ulchars[n] |= uint32(*((*uint8)(end))) << offset
		end = PointerStepIn(end, step)
		if offset == 0 {
			offset = 8 * step2
			n++
		} else {
			offset -= 8 * step
		}
	}
	return ulchars
}

func W2CA(p *uint16, size int) []uint8 {
	var chars []uint8
	if p == nil {
		return chars
	}

	end := unsafe.Pointer(p)
	step := unsafe.Sizeof(chars[0])
	step2 := unsafe.Sizeof(*p)

	n := 0
	offset := 8 * step2

	for *((*uint8)(end)) != 0 && n < size*4 {
		chars[n] = uint8((*((*uint8)(end))) >> offset)
		if offset == 0 {
			offset = 8 * step2
			end = PointerStepIn(end, step2)
		} else {
			offset -= 8 * step
			n++
		}
	}

	return chars
}

func PtrToArray_uint16(p *uint16, size int) []uint16 {
	var s []uint16
	if p == nil {
		return s
	}

	// Find NUL terminator.
	end := unsafe.Pointer(p)
	n := 0
	for *(*uint16)(end) != 0 && (size == 0 || n < size) {
		s = append(s, *((*uint16)(end)))
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}

	return s
}

func PtrToArray_byte(p *byte, size int) []byte {
	var s []byte
	if p == nil {
		return s
	}

	// Find NUL terminator.
	end := unsafe.Pointer(p)
	n := 0
	for *(*byte)(end) != 0 && (size == 0 || n < size) {
		s = append(s, *((*byte)(end)))
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}

	return s
}

func PtrToString_uint16(p *uint16, size int) string {
	return ToString_uint16(PtrToArray_uint16(p, size))
}

func PtrToString_byte(p *byte, size int) string {
	return ToString_byte(PtrToArray_byte(p, size))
}

func ToString_uint16(content []uint16) string {
	return fmt.Sprintf("%s", content)
}

func ToString_byte(content []byte) string {
	return fmt.Sprintf("%s", content)
}

func Compare_String_byte(s string, p *byte, size int) bool {
	c := []byte(s)

	end := unsafe.Pointer(p)
	n := 0
	for *(*byte)(end) != 0 && (size == 0 || n < size) {
		if c[n] != *((*byte)(end)) {
			return false
		}
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}

	if n != len(s) {
		return false
	}

	return true
}

func Compare_byte_by_byte(p1 *byte, size1 int, p2 *byte, size2 int) bool {
	if size1 != size2 {
		return false
	}

	end1 := unsafe.Pointer(p1)
	end2 := unsafe.Pointer(p2)

	n := 0
	for n < size1 {
		if *((*byte)(end1)) != *((*byte)(end2)) {
			return false
		}

		if *((*byte)(end1)) == 0 {
			break
		}

		end1 = unsafe.Pointer(uintptr(end1) + unsafe.Sizeof(*p1))
		end2 = unsafe.Pointer(uintptr(end1) + unsafe.Sizeof(*p2))
		n++
	}

	return true
}

func Transurlhex(p *byte) string {
	r := ""
	end := unsafe.Pointer(p)
	for *(*byte)(end) != 0 {
		r += "%"
		r += fmt.Sprintf("%02x", *(*byte)(end))
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
	}
	return r
}
