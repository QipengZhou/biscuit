package common

import "unsafe"

func Rounddown(v int, b int) int {
	return v - (v % b)
}

func Roundup(v int, b int) int {
	return Rounddown(v+b-1, b)
}

func Mkdev(_maj, _min int) uint {
	maj := uint(_maj)
	min := uint(_min)
	if min > 0xff {
		panic("bad minor")
	}
	m := maj<<8 | min
	return uint(m << 32)
}

func Readn(a []uint8, n int, off int) int {
	p := unsafe.Pointer(&a[off])
	var ret int
	switch n {
	case 8:
		ret = *(*int)(p)
	case 4:
		ret = int(*(*uint32)(p))
	case 2:
		ret = int(*(*uint16)(p))
	case 1:
		ret = int(*(*uint8)(p))
	default:
		panic("no")
	}
	return ret
}

func Writen(a []uint8, sz int, off int, val int) {
	p := unsafe.Pointer(&a[off])
	switch sz {
	case 8:
		*(*int)(p) = val
	case 4:
		*(*uint32)(p) = uint32(val)
	case 2:
		*(*uint16)(p) = uint16(val)
	case 1:
		*(*uint8)(p) = uint8(val)
	default:
		panic("no")
	}
}
