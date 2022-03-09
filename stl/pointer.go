package stl

import (
	"unsafe"
)

func PointerStepIn(p unsafe.Pointer, step uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + step)
}

func PointerStepOut(p unsafe.Pointer, step uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) - step)
}

func Copy_string(p *byte, str string) {
	up := unsafe.Pointer(p)
	step := unsafe.Sizeof(*p)
	s := ([]byte)(str)
	for _, v := range s {
		*((*byte)(up)) = v
		up = PointerStepIn(up, step)
	}
}

func Append_string(p *byte, limit int, str string) (*byte, bool) {
	savePointer := uintptr(unsafe.Pointer(p))
	end := unsafe.Pointer(p)
	n := 0
	for *(*byte)(end) != 0 && (limit == 0 || n < limit) {
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}

	if n >= limit {
		return (*byte)(unsafe.Pointer(savePointer)), false
	}

	Copy_string((*byte)(end), str)

	return (*byte)(unsafe.Pointer(savePointer)), true
}

func Ptr_Copy(p *byte, b *byte, length int) {
	dest := unsafe.Pointer(p)
	source := unsafe.Pointer(b)
	step := unsafe.Sizeof(*p)
	for i := 0; i < length; i++ {
		*((*byte)(dest)) = *((*byte)(source))
		dest = PointerStepIn(dest, step)
		source = PointerStepIn(source, step)
	}
}

func Ptr_Append(p *byte, limit int, buf *byte, length int) (*byte, bool) {
	savePointer := uintptr(unsafe.Pointer(p))
	end := unsafe.Pointer(p)
	n := 0
	for *(*byte)(end) != 0 && (limit == 0 || n < limit) {
		end = unsafe.Pointer(uintptr(end) + unsafe.Sizeof(*p))
		n++
	}

	if n >= limit {
		return (*byte)(unsafe.Pointer(savePointer)), false
	}

	Ptr_Copy((*byte)(end), buf, length)

	return (*byte)(unsafe.Pointer(savePointer)), true
}
